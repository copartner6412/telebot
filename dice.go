package telegram

// Dice represents an animated emoji that displays a random value.
//
// See "Dice" https://core.telegram.org/bots/api#dice
type Dice struct {
	// (Required) Emoji on which the dice throw animation is based.
	Emoji string `json:"emoji"`

	// (Required) Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji.
	Value int `json:"value"`
}
