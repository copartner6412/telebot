package telegram

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
//
// See "MessageAutoDeleteTimerChanged" https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	// (Required) New auto-delete time for messages in the chat; in seconds.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}
