package telegram

// MessageOriginChat represents the origin of a message sent on behalf of a chat.
//
// See "MessageOriginChat" https://core.telegram.org/bots/api#messageoriginchat
type MessageOriginChat struct {
	// (Required) Type of the message origin, always “chat”.
	Type string `json:"type"`

	// (Required) Date the message was sent originally in Unix time.
	Date int `json:"date"`

	// (Required) Chat that sent the message originally.
	SenderChat Chat `json:"sender_chat"`

	// (Optional) For messages originally sent by an anonymous chat administrator, original message author signature.
	AuthorSignature *string `json:"author_signature,omitempty"`
}
