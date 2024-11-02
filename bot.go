package telegram

import (
	"net/http"
	"strings"
)

type Bot struct {
	Token  string
	client *http.Client
}

// Each bot is given a unique authentication token when it is created.
func NewBot(token string) Bot {
	return Bot{
		Token:  strings.TrimSpace(token),
		client: &http.Client{Timeout: 10},
	}
}
