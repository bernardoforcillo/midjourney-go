package discord

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type InteractionRequest struct {
	Type          int            `json:"type"`
	ApplicationID string         `json:"application_id"`
	MessageFlags  *int           `json:"message_flags,omitempty"`
	MessageID     *string        `json:"message_id,omitempty"`
	GuildID       string         `json:"guild_id"`
	ChannelID     string         `json:"channel_id"`
	SessionID     string         `json:"session_id"`
	Data          map[string]any `json:"data"`
}

func (c *DiscordClient) SendInteraction(request *InteractionRequest) (error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", DiscordApi + "/interactions", strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return  err
	}
	log.Default().Printf("resp: %s\n, reqbody: %s \n, respbody: %s", resp.Status, string(reqBody), string(respBody))
	return nil
}
