package telegram

// ChatBackground represents a chat background.
//
// See "ChatBackground" https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	// (Required) Type of the background.
	Type BackgroundType `json:"type"`
}
