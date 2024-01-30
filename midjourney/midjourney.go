package midjourney

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bernardoforcillo/midjourney-go/discord"
)

const (
	ApplicationID string = "936929561302675456"
	SessionID     string = "2fb980f65e5c9a77c96ca01f2c242cf6"
)

type MidjourneyClient struct {
	client    discord.DiscordClient
	channelID string
}

func New(authToken string, channelID string) *MidjourneyClient {
	client := discord.NewClient(authToken)
	return &MidjourneyClient{
		client:    *client,
		channelID: channelID,
	}
}

func (c *MidjourneyClient) Imagine(prompt string, waitUntilGenerated bool) (*GeneratedImage, error) {
	channel, err := c.client.Channel(c.channelID)
	if err != nil {
		return nil, err
	}
	interactionsReq := &discord.InteractionRequest{
		Type:          2,
		ApplicationID: ApplicationID,
		GuildID:       channel.GuildID,
		ChannelID:     c.channelID,
		SessionID:     SessionID,
		Data: map[string]any{
			"version": "1166847114203123795",
			"id":      "938956540159881230",
			"name":    "imagine",
			"type":    "1",
			"options": []map[string]any{
				{
					"type":  3,
					"name":  "prompt",
					"value": prompt,
				},
			},
			"application_command": map[string]any{
				"id":                         "938956540159881230",
				"application_id":             ApplicationID,
				"version":                    "1166847114203123795",
				"default_permission":         true,
				"default_member_permissions": nil,
				"type":                       1,
				"nsfw":                       false,
				"name":                       "imagine",
				"description":                "Create images with Midjourney",
				"dm_permission":              true,
				"options": []map[string]any{
					{
						"type":        3,
						"name":        "prompt",
						"description": "The prompt to imagine",
						"required":    true,
					},
				},
				"attachments": []any{},
			},
		},
	}
	err = c.client.SendInteraction(interactionsReq)
	if err != nil {
		return nil, err
	}
	if waitUntilGenerated {
		var txtmessage string
		time.Sleep(2 * time.Second)
		for txtmessage != "fast" {
			msx, err := c.SearchMesssageByPrompt(prompt)
			if err != nil {
				return nil, err
			}
			regexPattern := `\*\*(.*?)\*\*.*?\((.*?)\)`
			regexpPattern, err := regexp.Compile(regexPattern)
			if err != nil {
				return nil, err
			}
			matches := regexpPattern.FindStringSubmatch(msx.Content)
			if len(matches) >= 3 {
				txtmessage = matches[2]
			}
		}
	}
	result, err := c.SearchMesssageByPrompt(prompt)
	if err != nil {
		return nil, err
	}
	generated := GeneratedImage{
		mj: c,
		message: result,
	}
	return &generated, nil
}

func (c MidjourneyClient) SearchMesssageByPrompt(prompt string) (discord.Message, error) {
	messages, err := c.client.Messages(c.channelID)
	if err != nil {
		return discord.Message{}, err
	}
	var split []string
	var result discord.Message
	for _, message := range *messages {
		split = strings.Split(message.Content, "**")
		if len(split) > 0 && split[1] == prompt {
			result = message
			break
		}
	}
	return result, nil
}

func (c MidjourneyClient) SearchMesssageWithContent(content string) (discord.Message, error) {
	messages, err := c.client.Messages(c.channelID)
	if err != nil {
		return discord.Message{}, err
	}
	var result discord.Message
	for _, message := range *messages {
		if strings.Contains(message.Content, content) {
			result = message
			break
		}
	}
	return result, nil
}

type GeneratedImage struct {
	mj  *MidjourneyClient
	message discord.Message
}

func (g GeneratedImage) URL() string {
	return g.message.Attachments[0].Url
}

func (g *GeneratedImage) Upscale(index int, waitUntilGenerated bool) (*UpscaledImage, error) {
	if index < 0 || index > 3 {
		return nil, fmt.Errorf("upscale index must be between 0 and 3")
	}
	channel, err := g.mj.client.Channel(g.mj.channelID)
	if err != nil {
		return nil, err
	}
	flags := 0
	interactionsReq := &discord.InteractionRequest{
		Type:          3,
		ApplicationID: ApplicationID,
		GuildID:       channel.GuildID,
		ChannelID:     channel.ID,
		MessageFlags:  &flags,
		MessageID:     &g.message.ID,
		SessionID:     SessionID,
		Data: map[string]any{
			"component_type": 2,
			"custom_id":      g.message.Components[0].Components[index].CustomID,
		},
	}
	err = g.mj.client.SendInteraction(interactionsReq)
	if err != nil {
		return nil, err
	}
	if waitUntilGenerated {
		time.Sleep(3 * time.Second)
	}
	prompt := strings.Split(g.message.Content, "**")[1]
	result, err := g.mj.SearchMesssageWithContent(fmt.Sprintf("**%s** - Image #%d", prompt, index+1))
	if err != nil {
		return nil, err
	}
	upscaled := UpscaledImage{
		mj:     g.mj,
		message: result,
	} 
	return &upscaled, nil
}

type UpscaledImage struct {
	mj *MidjourneyClient
	message discord.Message
}

func (u UpscaledImage) URL() string {
	return u.message.Attachments[0].Url
}
