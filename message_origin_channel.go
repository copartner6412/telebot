package telegram

// MessageOriginChannel represents the origin of a message sent to a channel chat.
//
// See "MessageOriginChannel" https://core.telegram.org/bots/api#messageoriginchannel
type MessageOriginChannel struct {
	// (Required) Type of the message origin, always “channel”.
	Type string `json:"type"`

	// (Required) Date the message was sent originally in Unix time.
	Date int `json:"date"`

	// (Required) Channel chat to which the message was originally sent.
	Chat Chat `json:"chat"`

	// (Required) Unique message identifier inside the chat.
	MessageID int `json:"message_id"`

	// (Optional) Signature of the original post author.
	AuthorSignature *string `json:"author_signature,omitempty"`
}
