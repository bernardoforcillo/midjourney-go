package midjourney

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	*http.Client

	Config *struct {
		UserToken string
	}
}

func NewClient(cfg *struct{ UserToken string }) *Client {
	return &Client{http.DefaultClient, cfg}
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("call io.ReadAll failed, err: %w", err)
		}

		return fmt.Errorf("resp.StatusCode: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
