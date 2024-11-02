package telegram

// MessageOriginUser represents the origin of a message sent by a known user.
//
// See "MessageOriginUser" https://core.telegram.org/bots/api#messageoriginuser
type MessageOriginUser struct {
	// (Required) Type of the message origin, always “user”.
	Type string `json:"type"`

	// (Required) Date the message was sent originally in Unix time.
	Date int `json:"date"`

	// (Required) User that sent the message originally.
	SenderUser User `json:"sender_user"`
}
