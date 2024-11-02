package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WebhookInfo describes the current status of a webhook.
//
// See "WebhookInfo" https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	// (Required) Webhook URL, may be empty if webhook is not set up.
	URL string `json:"url"`

	// (Required) True, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// (Required) Number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`

	// (Optional) Currently used webhook IP address.
	IPAddress string `json:"ip_address"`

	// (Optional) Unix time for the most recent error that happened when trying to deliver an update via webhook.
	LastErrorDate int `json:"last_error_date"`

	// (Optional) Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook.
	LastErrorMessage string `json:"last_error_message"`

	// (Optional) Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters.
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date"`

	// (Optional) The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery.
	MaxConnections int `json:"max_connections"`

	// (Optional) A list of update types the bot is subscribed to. Defaults to all update types except chat_member.
	AllowedUpdates []string `json:"allowed_updates"`
}

// GetWebhookInfo retrieves the current status of the webhook.
//
// If the bot is using getUpdates, will return an object with the url field empty.
//
// See "getWebhookInfo" https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo() (WebhookInfo, error) {
	httpResponse, err := b.sendRequest("GET", "application/json", "getWebhookInfo", nil)
	if err != nil {
		return WebhookInfo{}, err
	}
	defer httpResponse.Body.Close()

	var response struct {
		Ok          bool               `json:"ok"`
		Result      WebhookInfo        `json:"result"`
		Description string             `json:"description"`
		ErrorCode   int                `json:"error_code"`
		Parameters  ResponseParameters `json:"parameters"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return WebhookInfo{}, fmt.Errorf("error decoding response: %w", err)
	}

	if httpResponse.StatusCode != http.StatusOK || !response.Ok {
		return WebhookInfo{}, fmt.Errorf("HTTP status %s, Telegram code %d, Telegram API error: %s", httpResponse.Status, response.ErrorCode, response.Description)
	}

	return response.Result, nil
}
