package telegram

// ChatShared contains information about a chat that was shared with the bot using a KeyboardButtonRequestChat button.
//
// See "ChatShared" https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	// (Required) Identifier of the request.
	RequestID int64 `json:"request_id"`

	// (Required) Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	ChatID int64 `json:"chat_id"`

	// (Optional) Title of the chat, if the title was requested by the bot.
	Title *string `json:"title,omitempty"`

	// (Optional) Username of the chat, if the username was requested by the bot and available.
	Username *string `json:"username,omitempty"`

	// (Optional) Available sizes of the chat photo, if the photo was requested by the bot.
	Photo *[]PhotoSize `json:"photo,omitempty"`
}
