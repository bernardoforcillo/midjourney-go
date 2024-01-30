package discord

type User struct {
	Id                   string        `json:"id"`
	Username             string        `json:"username"`
	Avatar               interface{}   `json:"avatar"`
	Bot                  bool          `json:"bot"`
	System               bool          `json:"system"`
	Discriminator        string        `json:"discriminator"`
	PublicFlags          int           `json:"public_flags"`
	PremiumType          int           `json:"premium_type"`
	Flags                int           `json:"flags"`
	Banner               interface{}   `json:"banner"`
	AccentColor          interface{}   `json:"accent_color"`
	GlobalName           string        `json:"global_name"`
	AvatarDecorationData interface{}   `json:"avatar_decoration_data"`
	BannerColor          interface{}   `json:"banner_color"`
	MfaEnabled           bool          `json:"mfa_enabled"`
	Locale               string        `json:"locale"`
	Email                string        `json:"email"`
	Verified             bool          `json:"verified"`
	Phone                string        `json:"phone"`
	NsfwAllowed          bool          `json:"nsfw_allowed"`
	LinkedUsers          []interface{} `json:"linked_users"`
	Bio                  string        `json:"bio"`
	AuthenticatorTypes   []interface{} `json:"authenticator_types"`
}