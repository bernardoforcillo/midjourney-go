package discord

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	DiscordApi = "https://discord.com/api/v10"
)

type DiscordClient struct {
	authToken string
	client *http.Client
}



func NewClient(authToken string) *DiscordClient {
	client := http.DefaultClient
	return &DiscordClient{
		client: client,
		authToken: authToken,
	}
}

func (c DiscordClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.authToken)
	return c.client.Do(req)
}

func (c DiscordClient) GetUser() (*User, error) {
	req, err := http.NewRequest("GET", DiscordApi + "/users/@me", nil)
	if err != nil {
		return nil, err
	}	
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return unMarshal[User](body)
}

func (c DiscordClient) Channel(channelID string) (*Channel, error) {
	req, err := http.NewRequest("GET", DiscordApi + "/channels/" + channelID, nil)
	if err != nil {
		return nil, err
	}	
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return unMarshal[Channel](body)
}

func (c DiscordClient) Messages(channelID string) (*[]Message, error) {
	req, err := http.NewRequest("GET", DiscordApi + "/channels/" + channelID + "/messages", nil)
	if err != nil {
		return nil, err
	}	
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return unMarshal[[]Message](body)
}

func unMarshal[T any](body []byte) (*T, error) {
	if body == nil {
		return nil, nil
	}
	var t T
	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

