package telegram

// InputPollOption contains information about one answer option in a poll to be sent.
//
// See "InputPollOption" https://core.telegram.org/bots/api#inputpolloption
type InputPollOption struct {
	// (Required) Option text, 1-100 characters.
	Text string `json:"text"`

	// (Optional) Mode for parsing entities in the text. See formatting options for more details. Currently, only custom emoji entities are allowed.
	TextParseMode *string `json:"text_parse_mode,omitempty"`

	// (Optional) A JSON-serialized list of special entities that appear in the poll option text. It can be specified instead of text_parse_mode.
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
}
