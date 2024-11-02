package telegram

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
//
// See "ProximityAlertTriggered" https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	// (Required) User that triggered the alert.
	TravelerUser User `json:"traveler_user"`

	// (Required) User that set the alert.
	WatcherUser User `json:"watcher_user"`

	// (Required) The distance between the users.
	Distance int `json:"distance"`
}
