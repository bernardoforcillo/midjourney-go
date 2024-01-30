package discord

import (
	"time"
)

type Message struct {
	ID               string               `json:"id"`
	ChannelID        string               `json:"channel_id"`
	GuildID          string               `json:"guild_id"`
	Content          string               `json:"content"`
	Author           *User                `json:"author"`
	Flags            int                  `json:"flags"`
	Timestamp        time.Time            `json:"timestamp"`
	EditedTimestamp  *time.Time           `json:"edited_timestamp"`
	Tts              bool                 `json:"tts"`
	MentionEveryone  bool                 `json:"mention_everyone"`
	Mentions         []User               `json:"mentions"`
	MentionsChannels []any                `json:"mention_channels"`
	Attachments      []MesssageAttachment `json:"attachments"`
	Components       []MessageActionRow   `json:"components"`
}

type MesssageAttachment struct {
	Id          string  `json:"id"`
	Filename    string  `json:"filename"`
	Description string  `json:"description"`
	ContentType string  `json:"content_type"`
	Size        int64   `json:"size"`
	Url         string  `json:"url"`
	ProxyUrl    string  `json:"proxy_url"`
	Height      int64   `json:"height"`
	Width       int64   `json:"width"`
	Ephemeral   bool    `json:"ephemeral"`
	DurationSec float64 `json:"duration_secs"`
	Waveform    string  `json:"waveform"`
	Flags       int64   `json:"flags"`
}

type MessageActionRow struct {
	Type       int                `json:"type"`
	Components []MessageButton `json:"components"`
}

type MessageButton struct {
	Type  int    `json:"type"`
	Style int    `json:"style"`
	Label string `json:"label"`
	Emoji any `json:"emoji"`
	CustomID string `json:"custom_id"`
	Url   string `json:"url"`
	Disabled bool `json:"disabled"`
}
