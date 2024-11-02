package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// SetWebhookRequest represents a request to set a webhook for receiving updates via HTTPS POST requests.
//
// You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.
//
// To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.
//
// Ports currently supported for webhooks: 443, 80, 88, 8443.
//
// See "setWebhook" https://core.telegram.org/bots/api#setwebhook
//
// See "amazing guide to webhooks" https://core.telegram.org/bots/webhooks
type SetWebhookRequest struct {
	// (Required) HTTPS URL to send updates to. Use an empty string to remove webhook integration.
	URL string `json:"url"`

	// (Optional) Upload your public key certificate so that the root certificate in use can be checked.
	// See our self-signed guide for details https://core.telegram.org/bots/self-signed
	CertificatePath *string `json:"certificate,omitempty"`

	// (Optional) The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS.
	IPAddress *string `json:"ip_address,omitempty"`

	//
	// (Optional) The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	// Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections *int `json:"max_connections,omitempty"`

	// (Optional) A JSON-serialized list of the update types you want your bot to receive. For example, specify
	// ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types.
	// Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`

	// (Optional) Pass True to drop all pending updates.
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`

	// (Optional) A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request,
	// 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
	SecretToken *string `json:"secret_token,omitempty"`
}

// SetWebhook sets a webhook to receive updates via HTTPS POST requests.
//
// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
//
// See "setWebhook" https://core.telegram.org/bots/api#setwebhook
//
// See "amazing guide to webhooks" https://core.telegram.org/bots/webhooks
func (b *Bot) SetWebhook(request SetWebhookRequest) error {
	requestPayload := &bytes.Buffer{}
	writer := multipart.NewWriter(requestPayload)

	if err := writer.WriteField("url", request.URL); err != nil {
		return fmt.Errorf("error adding URL field: %w", err)
	}

	if request.CertificatePath != nil {
		certFile, err := os.Open(*request.CertificatePath)
		if err != nil {
			return fmt.Errorf("error opening certificate file: %w", err)
		}
		defer certFile.Close()
		fileWriter, err := writer.CreateFormFile("certificate", certFile.Name())
		if err != nil {
			return fmt.Errorf("error creating form file for certificate: %w", err)
		}
		if _, err := io.Copy(fileWriter, certFile); err != nil {
			return fmt.Errorf("error copying certificate file: %w", err)
		}
	}

	if request.IPAddress != nil {
		if err := writer.WriteField("ip_address", *request.IPAddress); err != nil {
			return fmt.Errorf("error adding IP address field: %w", err)
		}
	}
	if request.MaxConnections != nil {
		if err := writer.WriteField("max_connections", fmt.Sprintf("%d", *request.MaxConnections)); err != nil {
			return fmt.Errorf("error adding max connections field: %w", err)
		}
	}
	if request.AllowedUpdates != nil {
		updatesJSON, err := json.Marshal(request.AllowedUpdates)
		if err != nil {
			return fmt.Errorf("error marshaling allowed updates: %w", err)
		}
		if err := writer.WriteField("allowed_updates", string(updatesJSON)); err != nil {
			return fmt.Errorf("error adding allowed updates field: %w", err)
		}
	}
	if request.DropPendingUpdates != nil {
		if err := writer.WriteField("drop_pending_updates", fmt.Sprintf("%t", *request.DropPendingUpdates)); err != nil {
			return fmt.Errorf("error adding drop pending updates field: %w", err)
		}
	}

	url := fmt.Sprintf(telegramEndpoint, b.Token, "setWebhook")
	httpRequest, err := http.NewRequest("POST", url, requestPayload)
	if err != nil {
		return fmt.Errorf("error creating new POST request to setWebhook: %w", err)
	}

	if request.SecretToken != nil {
		if err := writer.WriteField("secret_token", *request.SecretToken); err != nil {
			return fmt.Errorf("error adding secret token field: %w", err)
		}
		httpRequest.Header.Set("X-Telegram-Bot-Api-Secret-Token", *request.SecretToken)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("error closing multipart writer: %w", err)
	}

	httpRequest.Header.Set("Content-Type", writer.FormDataContentType())
	httpResponse, err := b.client.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("error sending POST request to setWebhook: %w", err)
	}
	defer httpResponse.Body.Close()

	var response Response
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if httpResponse.StatusCode != http.StatusOK || !response.Ok {
		return fmt.Errorf("HTTP status %s, Telegram code %d, Telegram API error: %s", httpResponse.Status, response.ErrorCode, response.Description)
	}

	return nil
}
