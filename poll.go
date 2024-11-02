package telegram

// Poll contains information about a poll.
//
// See "Poll" https://core.telegram.org/bots/api#poll
type Poll struct {
	// (Required) Unique poll identifier.
	ID string `json:"id"`

	// (Required) Poll question, 1-300 characters.
	Question string `json:"question"`

	// (Optional) Special entities that appear in the question. Currently, only custom emoji entities are allowed in poll questions.
	QuestionEntities *[]MessageEntity `json:"question_entities,omitempty"`

	// (Required) List of poll options.
	Options []PollOption `json:"options"`

	// (Required) Total number of users that voted in the poll.
	TotalVoterCount int `json:"total_voter_count"`

	// (Required) True, if the poll is closed.
	IsClosed bool `json:"is_closed"`

	// (Required) True, if the poll is anonymous.
	IsAnonymous bool `json:"is_anonymous"`

	// (Required) Poll type, currently can be “regular” or “quiz”.
	Type string `json:"type"`

	// (Required) True, if the poll allows multiple answers.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// (Optional) 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionID *int `json:"correct_option_id,omitempty"`

	// (Optional) Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters.
	Explanation *string `json:"explanation,omitempty"`

	// (Optional) Special entities like usernames, URLs, bot commands, etc. that appear in the explanation.
	ExplanationEntities *[]MessageEntity `json:"explanation_entities,omitempty"`

	// (Optional) Amount of time in seconds the poll will be active after creation.
	OpenPeriod *int `json:"open_period,omitempty"`

	// (Optional) Point in time (Unix timestamp) when the poll will be automatically closed.
	CloseDate *int `json:"close_date,omitempty"`
}
