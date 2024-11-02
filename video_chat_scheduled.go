package telegram

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
//
// See "VideoChatScheduled" https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	// (Required) Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
	StartDate int64 `json:"start_date"`
}
