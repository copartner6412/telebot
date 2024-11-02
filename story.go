package telegram

// Story represents a story.
//
// See "Story" https://core.telegram.org/bots/api#story
type Story struct {
	// (Required) Chat that posted the story.
	Chat Chat `json:"chat"`

	// (Required) Unique identifier for the story in the chat.
	ID int `json:"id"`
}
