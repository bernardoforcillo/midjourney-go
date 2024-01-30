package discord

type Channel struct {
	ID            string `json:"id"`
	Type          int    `json:"type"`
	GuildID       string `json:"guild_id"`
	Name          string `json:"name"`
	Position      int    `json:"position"`
	Topic         string `json:"topic"`
	LastMessageID string `json:"last_message_id"`
	Bitrate       int    `json:"bitrate"`
	UserLimit     int    `json:"user_limit"`
	ParentID      string `json:"parent_id"`
	Flags         int    `json:"flags"`
}
