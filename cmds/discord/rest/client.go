package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/internal/retry"
)

type Client struct {
	token string
	http  http.Client

	applicationID string
}

func NewClient(token string) (*Client, error) {
	client := &Client{
		token: token,
		http:  http.Client{Timeout: time.Millisecond * 1000},
	}

	_, err := client.lookupApplicationID()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) request(method string, url string, payload any) (*http.Request, error) {
	var body io.Reader
	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to encode json payload: %s", err)
		}
		body = bytes.NewBuffer(encoded)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make a new request: %s", err)
	}
	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) requestWithFiles(method string, url string, payload any, files []*discordgo.File) (*http.Request, error) {
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

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to make a new request: %s", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func partHeader(contentDisposition string, contentType string) textproto.MIMEHeader {
	return textproto.MIMEHeader{
		"Content-Disposition": []string{contentDisposition},
		"Content-Type":        []string{contentType},
	}
}

func (c *Client) do(req *http.Request, target any) error {
	result := retry.Retry(func() (any, error) {
		res, err := c.http.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			var body discordgo.APIErrorMessage
			_ = json.NewDecoder(res.Body).Decode(&body)
			message := body.Message
			if message == "" {
				message = res.Status
			}

			return nil, errors.New("discord: " + strings.ToLower(message))
		}

		if target != nil {
			err = json.NewDecoder(res.Body).Decode(target)
			if err != nil {
				return nil, fmt.Errorf("failed to decode response body :%w", err)
			}
		}

		return nil, nil
	}, 3, time.Millisecond*150)

	return result.Err
}

func (c *Client) lookupApplicationID() (string, error) {
	req, err := c.request("GET", discordgo.EndpointApplication("@me"), nil)
	if err != nil {
		return "", err
	}

	var data discordgo.Application
	err = c.do(req, &data)
	if err != nil {
		return "", err
	}
	if data.ID == "" {
		return "", errors.New("blank application id returned")
	}

	c.applicationID = data.ID
	return data.ID, nil
}
