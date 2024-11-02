package telegram

// Chat represents a chat in Telegram.
//
// This object holds information about various types of chats in the platform.
//
// See "Chat" https://core.telegram.org/bots/api#chat
type Chat struct {
	// (Required) Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// identifier.
	ID int64 `json:"id"`

	// (Required) Type of the chat, can be either “private”, “group”, “supergroup” or “channel”.
	Type string `json:"type"`

	// (Optional) Title, for supergroups, channels, and group chats.
	Title *string `json:"title,omitempty"`

	// (Optional) Username, for private chats, supergroups, and channels if available.
	Username *string `json:"username,omitempty"`

	// (Optional) First name of the other party in a private chat.
	FirstName *string `json:"first_name,omitempty"`

	// (Optional) Last name of the other party in a private chat.
	LastName *string `json:"last_name,omitempty"`

	// (Optional) True, if the supergroup chat is a forum (has topics enabled).
	IsForum *bool `json:"is_forum,omitempty"`
}
