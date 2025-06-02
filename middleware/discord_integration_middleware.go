package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"net/http"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// WebhookChannelSource represents the source of the Discord channel
type WebhookChannelSource string

const (
	WebhookChannelSourceERRORDATABASETRX       WebhookChannelSource = "error database transaction"
	WebhookChannelSourceERRORCONNECTIONTIMEOUT WebhookChannelSource = "connection timeout"
	WebhookChannelSourceNEWS                   WebhookChannelSource = "news"
)

// DiscordEmbed represents the embedded message structure for Discord
type DiscordEmbed struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Color       int          `json:"color"`
	Fields      []EmbedField `json:"fields"`
	Timestamp   string       `json:"timestamp"`
}

// EmbedField represents a field inside an embed
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

// DiscordWebhookPayload represents the structure for sending messages to Discord
type DiscordWebhookPayload struct {
	Embeds []DiscordEmbed `json:"embeds"`
}

// SendDiscordNotificationError sends an error message with an embed to a Discord webhook
func SendDiscordNotificationError(source WebhookChannelSource, err error) error {
	if err == nil {
		return nil // No error to send
	}

	var webhookURL string
	switch source {
	case WebhookChannelSourceERRORCONNECTIONTIMEOUT:
		webhookURL = viper.GetString("discord.alert.connection.timeout.transaction.webhook")
	case WebhookChannelSourceERRORDATABASETRX:
		webhookURL = viper.GetString("discord.alert.database.transaction.webhook")
	}

	if webhookURL == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL is not set")
	}

	// Get function name where the error happened
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()

	// Construct the embed message
	embed := DiscordEmbed{
		Title:       "🚨 Error Alert 🚨",
		Description: "An error occurred in mceasy.",
		Color:       0xFF0000, // Red color for errors
		Fields: []EmbedField{
			{Name: "🔹 **Source**", Value: fmt.Sprintf("`%s`", source), Inline: false},
			{Name: "ℹ️ **Function**", Value: fmt.Sprintf("`%s`", funcName), Inline: false},
			{Name: "❌ **Error**", Value: fmt.Sprintf("```%v```", err), Inline: false},
			{Name: "🕒 **Timestamp**", Value: time.Now().Format(time.RFC3339), Inline: false},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	payload := DiscordWebhookPayload{Embeds: []DiscordEmbed{embed}}

	payloadBytes, jsonErr := json.Marshal(payload)
	if jsonErr != nil {
		return fmt.Errorf("failed to marshal Discord embed message: %w", jsonErr)
	}

	// Create a request with proper headers
	req, reqErr := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if reqErr != nil {
		return fmt.Errorf("failed to create request: %w", reqErr)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		return fmt.Errorf("failed to send request to Discord: %w", reqErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Discord responded with status: %d", resp.StatusCode)
	}

	return nil
}

func SendDiscordActiveTransactionNotification(newSignals []*dto.TradeSignal, dailyTimestamp string) error {
	if len(newSignals) == 0 {
		return nil
	}
	webhookURL := viper.GetString("webhook.active.transaction.url")

	var fields []EmbedField

	// Add new signals
	if len(newSignals) > 0 {
		fields = append(fields, EmbedField{
			Name:   "📈 New Signals",
			Value:  "-------------------",
			Inline: false,
		})

		for _, signal := range newSignals {
			reqID := uuid.New().String()
			fields = append(fields, EmbedField{
				Name:   fmt.Sprintf("🔹 %s", signal.Symbol),
				Value:  fmt.Sprintf("**Type:** %s\n**Signal Timestamp:** %s\n**Notification Identifier:** %s", signal.TradeType, dailyTimestamp, reqID),
				Inline: false,
			})
		}
	}

	embed := DiscordEmbed{
		Title:       "📢 Trade Signals Update",
		Description: "Here are the latest trade signals:",
		Color:       0x00FF00,
		Fields:      fields,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	payload := DiscordWebhookPayload{Embeds: []DiscordEmbed{embed}}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling Discord webhook payload:", err)
		return errors.New("error formatting trade signals.")
	}

	// Send POST request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("discord webhook returned status code: %d", resp.StatusCode)
	}
	return nil
}

func SendDiscordActiveTransactionNotificationV2(newSignals []*v2tickdto.TradeSignalV2) error {
	if len(newSignals) == 0 {
		return nil
	}
	webhookURL := viper.GetString("webhook.active.transaction.url")

	var fields []EmbedField

	// Add new signals
	if len(newSignals) > 0 {
		fields = append(fields, EmbedField{
			Name:   "📈 New Signals",
			Value:  "-------------------",
			Inline: false,
		})

		for _, signal := range newSignals {
			reqID := uuid.New().String()
			fields = append(fields, EmbedField{
				Name:   fmt.Sprintf("🔹 %s", signal.Symbol),
				Value:  fmt.Sprintf("**Type:** %s\n**Notification Identifier:** %s", signal.TradeType, reqID),
				Inline: false,
			})
		}
	}

	embed := DiscordEmbed{
		Title:       "📢 Trade Signals Update",
		Description: "Here are the latest trade signals:",
		Color:       0x00FF00,
		Fields:      fields,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	payload := DiscordWebhookPayload{Embeds: []DiscordEmbed{embed}}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling Discord webhook payload:", err)
		return errors.New("error formatting trade signals.")
	}

	// Send POST request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("discord webhook returned status code: %d", resp.StatusCode)
	}
	return nil
}

func SendDiscordNewsNotification(items []newsDTO.FinancialNewsItem, whitelistWords []string) error {
	webhookURL := viper.GetString("news.webhook.url")
	var fields []EmbedField

	fields = append(fields, EmbedField{
		Name:   "🔍 **Filtered By**",
		Value:  fmt.Sprintf("Impact containing: `%s` AND Today's date: `%s`", strings.Join(whitelistWords, "`, `"), time.Now().Format("2006-01-02")),
		Inline: false,
	})

	fields = append(fields, EmbedField{
		Name:   "📰 **News Items**",
		Value:  "-------------------",
		Inline: false,
	})

	for _, item := range items {
		dateTimeStr := item.Date
		if item.Time != "" {
			dateTimeStr += " " + item.Time
		}

		fields = append(fields, EmbedField{
			Name: fmt.Sprintf("🔹 %s", item.Title),
			Value: fmt.Sprintf("**Country:** %s\n**Date:** %s\n**Impact:** %s\n**Forecast:** %s\n**Previous:** %s",
				item.Country, dateTimeStr, item.Impact, item.Forecast, item.Previous),
			Inline: false,
		})
	}

	embed := DiscordEmbed{
		Title:       "📊 Financial News Alert",
		Description: fmt.Sprintf("Found %d financial news items for today. All trading activities are prohibited.", len(items)),
		Color:       0x00BFFF, // Deep Sky Blue color for news
		Fields:      fields,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	payload := DiscordWebhookPayload{
		Embeds: []DiscordEmbed{embed},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal Discord payload: %w", err)
	}

	// Send POST request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("discord webhook returned status code: %d", resp.StatusCode)
	}
	return nil
}

// sendSSEClientConnectedNotification sends a Discord webhook notification when a new SSE client connects
func SendSSEClientConnectedNotification(status string) error {
	webhookURL := viper.GetString("sse.clients.webhook.url")

	// Create the message payload
	payload := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       fmt.Sprintf("SSE client %s", status),
				"description": fmt.Sprintf("SSE client has %s to the system", status),
				"color":       3447003, // Blue color
				"timestamp":   time.Now().Format(time.RFC3339),
				"footer": map[string]interface{}{
					"text": "Mceasy System",
				},
			},
		},
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create the HTTP request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("discord webhook returned status code: %d", resp.StatusCode)
	}
	return nil
}
