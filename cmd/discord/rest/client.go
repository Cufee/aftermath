package rest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/json"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	token string
	http  http.Client

	rateLimitMx      sync.Mutex
	rateLimitBuckets map[string]time.Time

	applicationID string
}

func NewClient(token string) (*Client, error) {
	client := &Client{
		token:            token,
		rateLimitMx:      sync.Mutex{},
		rateLimitBuckets: make(map[string]time.Time),
		http:             http.Client{Timeout: time.Millisecond * 5000}, // discord is very slow sometimes
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := client.lookupApplicationID(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) request(method string, url string, payload any) (func() (*http.Request, error), error) {
	var body []byte
	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to encode json payload: %s", err)
		}
		body = encoded
	}

	return func() (*http.Request, error) {
		var payload io.Reader
		if body != nil {
			payload = bytes.NewBuffer(body)
		}

		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			return nil, fmt.Errorf("failed to make a new request: %s", err)
		}
		req.Header.Set("Authorization", "Bot "+c.token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		return req, nil
	}, nil
}

func (c *Client) requestWithFiles(method string, url string, payload any, files []File) (func() (*http.Request, error), error) {
	if len(files) > 0 {
		var df []*discordgo.File
		for _, f := range files {
			df = append(df, &discordgo.File{Name: f.Name, Reader: bytes.NewReader(f.Data)})
		}
		return c.requestWithFormData(method, url, payload, df)
	}
	return c.request(method, url, payload)
}

func (c *Client) requestWithFormData(method string, url string, payload any, files []*discordgo.File) (func() (*http.Request, error), error) {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	writer.FormDataContentType()

	encoded, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode json payload: %s", err)
	}

	part, err := writer.CreatePart(partHeader(`form-data; name="payload_json"`, "application/json"))
	if err != nil {
		return nil, err
	}

	if _, err = part.Write(encoded); err != nil {
		return nil, err
	}

	for i, file := range files {
		part, err = writer.CreatePart(partHeader(fmt.Sprintf(`form-data; name="files[%d]"; filename="%s"`, i, file.Name), "application/octet-stream"))
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(part, file.Reader); err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}
	data := buffer.Bytes()

	return func() (*http.Request, error) {
		req, err := http.NewRequest(method, url, bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("failed to make a new request: %s", err)
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Authorization", "Bot "+c.token)
		req.Header.Set("Accept", "application/json")
		return req, nil
	}, nil
}

func partHeader(contentDisposition string, contentType string) textproto.MIMEHeader {
	return textproto.MIMEHeader{
		"Content-Disposition": []string{contentDisposition},
		"Content-Type":        []string{contentType},
	}
}

func (c *Client) do(ctx context.Context, makeReq func() (*http.Request, error), target any, bucket ...string) error {
	if len(bucket) > 0 {
		c.rateLimitMx.Lock()
		resetTime := c.rateLimitBuckets[bucket[0]]
		c.rateLimitMx.Unlock()

		time.Sleep(time.Until(resetTime))
	}

	req, err := makeReq()
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 429 {
		b := res.Header.Get("X-RateLimit-Bucket")
		retryAfterSec := res.Header.Get("Retry-After")
		retryAfter, _ := strconv.Atoi(retryAfterSec)
		resetTime := time.Now().Add(time.Second * time.Duration(retryAfter+1))

		log.Info().Str("bucket", b).Time("resetTime", resetTime).Msg("discord api: you are being rate limited")

		c.rateLimitMx.Lock()
		c.rateLimitBuckets[b] = resetTime
		c.rateLimitMx.Unlock()

		return c.do(ctx, makeReq, target, b)
	}

	if res.StatusCode > 299 {
		var body discordgo.APIErrorMessage
		raw, _ := io.ReadAll(res.Body)
		_ = json.NewDecoder(bytes.NewBuffer(raw)).Decode(&body)
		message := body.Message
		if message == "" {
			log.Warn().Str("body", string(raw)).Msg("discord api returned invalid response")
			message = res.Status + ", response was not valid json"
		}
		if err := knownError(body.Code); err != nil {
			return err
		}
		return errors.New("discord api: " + message)
	}

	if target != nil {
		if res.Header.Get("Content-Type") != "application/json" {
			log.Warn().Msg("discord api returned invalid headers when json response body was expected")
			return nil
		}

		err = json.NewDecoder(res.Body).Decode(target)
		if err != nil {
			return fmt.Errorf("failed to decode response body :%w", err)
		}
	}
	return nil
}

func (c *Client) lookupApplicationID(ctx context.Context) (string, error) {
	newReq, err := c.request("GET", discordgo.EndpointApplication("@me"), nil)
	if err != nil {
		return "", err
	}

	var data discordgo.Application
	err = c.do(ctx, newReq, &data)
	if err != nil {
		return "", err
	}
	if data.ID == "" {
		return "", errors.New("blank application id returned")
	}

	c.applicationID = data.ID
	return data.ID, nil
}
