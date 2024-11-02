package telegram

// MessageId represents a unique message identifier in a Telegram message.
//
// See "MessageId" https://core.telegram.org/bots/api#messageid
type MessageId struct {
	// (Required) Unique message identifier. In specific instances (e.g., message containing a video sent to a big chat),
	// the server might automatically schedule a message instead of sending it immediately. In such cases, this field
	// will be 0 and the relevant message will be unusable until it is actually sent.
	MessageID int `json:"message_id"`
}
