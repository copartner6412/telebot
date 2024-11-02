package telegram

// PollOption contains information about one answer option in a poll.
//
// See "PollOption" https://core.telegram.org/bots/api#polloption
type PollOption struct {
	// (Required) Option text, 1-100 characters.
	Text string `json:"text"`

	// (Optional) Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts.
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`

	// (Required) Number of users that voted for this option.
	VoterCount int `json:"voter_count"`
}
