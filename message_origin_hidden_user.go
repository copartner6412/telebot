package telegram

// MessageOriginHiddenUser represents the origin of a message sent by an unknown user.
//
// See "MessageOriginHiddenUser" https://core.telegram.org/bots/api#messageoriginhiddenuser
type MessageOriginHiddenUser struct {
	// (Required) Type of the message origin, always “hidden_user”.
	Type string `json:"type"`

	// (Required) Date the message was sent originally in Unix time.
	Date int `json:"date"`

	// (Required) Name of the user that sent the message originally.
	SenderUserName string `json:"sender_user_name"`
}
