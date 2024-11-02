package telegram

// BackgroundTypeChatTheme represents the background taken directly from a built-in chat theme.
//
// See "BackgroundTypeChatTheme" https://core.telegram.org/bots/api#backgroundtypechattheme
type BackgroundTypeChatTheme struct {
	// (Required) Type of the background, always “chat_theme”.
	Type string `json:"type"`

	// (Required) Name of the chat theme, which is usually an emoji.
	ThemeName string `json:"theme_name"`
}
