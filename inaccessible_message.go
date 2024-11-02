package telegram

// InaccessibleMessage represents a message that was deleted or is otherwise inaccessible to the bot.
//
// See "InaccessibleMessage" https://core.telegram.org/bots/api#inaccessiblemessage
type InaccessibleMessage struct {
	// (Required) Chat the message belonged to.
	Chat Chat `json:"chat"`

	// (Required) Unique message identifier inside the chat.
	MessageID int `json:"message_id"`

	// (Required) Always 0. The field can be used to differentiate regular and inaccessible messages.
	Date int `json:"date"`
}
