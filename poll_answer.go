package telegram

// PollAnswer represents an answer of a user in a non-anonymous poll.
//
// See "PollAnswer" https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	// (Required) Unique poll identifier.
	PollID string `json:"poll_id"`

	// (Optional) The chat that changed the answer to the poll, if the voter is anonymous.
	VoterChat *Chat `json:"voter_chat,omitempty"`

	// (Optional) The user that changed the answer to the poll, if the voter isn't anonymous.
	User *User `json:"user,omitempty"`

	// (Required) 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIDs []int `json:"option_ids"`
}
