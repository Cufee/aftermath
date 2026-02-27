package rest

import (
	"bytes"
	"context"
	"errors"
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

	"github.com/bwmarrin/discordgo"
)

type RequestObserver interface {
	Record(source, operation string, failed bool)
}

type ClientOption func(*clientOptions)

type clientOptions struct {
	observer RequestObserver
}

func WithObserver(observer RequestObserver) ClientOption {
	return func(options *clientOptions) {
		options.observer = observer
	}
}

type Client struct {
	token string
	http  http.Client

	rateLimitMx      sync.Mutex
	rateLimitBuckets map[string]time.Time

	applicationID string
	observer      RequestObserver
}

func NewClient(token string, opts ...ClientOption) (*Client, error) {
	options := clientOptions{}
	for _, apply := range opts {
		apply(&options)
	}

	client := &Client{
		token:            token,
		rateLimitMx:      sync.Mutex{},
		rateLimitBuckets: make(map[string]time.Time),
		http:             http.Client{Timeout: time.Millisecond * 5000}, // discord is very slow sometimes
		observer:         options.observer,
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

func (c *Client) do(ctx context.Context, operation string, makeReq func() (*http.Request, error), target any, bucket ...string) error {
	var bucketID string
	if len(bucket) > 0 {
		bucketID = bucket[0]
	}

	for {
		if bucketID != "" {
			c.rateLimitMx.Lock()
			resetTime := c.rateLimitBuckets[bucketID]
			c.rateLimitMx.Unlock()

			time.Sleep(time.Until(resetTime))
		}

		req, err := makeReq()
		if err != nil {
			return c.observe(operation, err)
		}
		req = req.WithContext(ctx)

		res, err := c.http.Do(req)
		if err != nil {
			return c.observe(operation, err)
		}

		if res.StatusCode == 429 {
			b := res.Header.Get("X-RateLimit-Bucket")
			retryAfterSec := res.Header.Get("Retry-After")
			retryAfter, _ := strconv.Atoi(retryAfterSec)
			resetTime := time.Now().Add(time.Second * time.Duration(retryAfter+1))

			log.Info().Str("bucket", b).Time("resetTime", resetTime).Msg("discord api: you are being rate limited")

			c.rateLimitMx.Lock()
			c.rateLimitBuckets[b] = resetTime
			c.rateLimitMx.Unlock()
			_ = res.Body.Close()

			bucketID = b
			continue
		}

		if res.StatusCode > 299 {
			var body discordgo.APIErrorMessage
			raw, _ := io.ReadAll(res.Body)
			_ = res.Body.Close()
			_ = json.NewDecoder(bytes.NewBuffer(raw)).Decode(&body)

			message := body.Message
			if message == "" {
				log.Warn().Str("body", string(raw)).Msg("discord api returned invalid response")
				message = res.Status + ", response was not valid json"
			}
			if err := knownError(body.Code); err != nil {
				return c.observe(operation, err)
			}
			return c.observe(operation, errors.New("discord api: "+message))
		}

		if target != nil {
			if res.Header.Get("Content-Type") != "application/json" {
				log.Warn().Msg("discord api returned invalid headers when json response body was expected")
				_ = res.Body.Close()
				return c.observe(operation, nil)
			}

			err = json.NewDecoder(res.Body).Decode(target)
			_ = res.Body.Close()
			if err != nil {
				return c.observe(operation, fmt.Errorf("failed to decode response body :%w", err))
			}

			return c.observe(operation, nil)
		}

		_ = res.Body.Close()
		return c.observe(operation, nil)
	}
}

func (c *Client) lookupApplicationID(ctx context.Context) (string, error) {
	newReq, err := c.request("GET", discordgo.EndpointApplication("@me"), nil)
	if err != nil {
		return "", err
	}

	var data discordgo.Application
	err = c.do(ctx, "lookup_application", newReq, &data)
	if err != nil {
		return "", err
	}
	if data.ID == "" {
		return "", errors.New("blank application id returned")
	}

	c.applicationID = data.ID
	return data.ID, nil
}

func (c *Client) observe(operation string, err error) error {
	if c.observer == nil {
		return err
	}

	failed := err != nil && shouldCountDiscordFailure(err)
	c.observer.Record("discord", operation, failed)
	return err
}

func shouldCountDiscordFailure(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrUnknownWebhook) {
		return false
	}
	if errors.Is(err, ErrUnknownInteraction) {
		return false
	}
	if errors.Is(err, ErrInteractionAlreadyAcked) {
		return false
	}
	if errors.Is(err, ErrMissingPermissions) {
		return false
	}
	if errors.Is(err, ErrMissingUserUnreachable) {
		return false
	}
	return true
}
