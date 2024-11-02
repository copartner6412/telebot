package telegram

// VideoChatEnded represents a service message about a video chat ended in the chat.
//
// See "VideoChatEnded" https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	// (Required) Video chat duration in seconds
	Duration int `json:"duration"`
}
