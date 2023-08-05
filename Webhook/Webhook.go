package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"beemer/logger"
)

// Insert your webhook url in webhookURl.
//
// Username is the username that is displayed in the discord chat.
//
// Avatar_url is the url of the avatar you want to discplay in discord.
//
// Content is the main textbox from the webhook .
func DiscordWebhook(WebhookURL string, username string, avatar_url string, content string) (succes bool) {


	// Data to send as a message to the Discord webhook
	data := map[string]string{
		"username":   username,
		"avatar_url": avatar_url,
		"content":    content,
	}

	// Convert data to JSON
	payload, err := json.Marshal(data)
	if err != nil {
		logger.LogErrorToFile("Error encoding JSON:", err)
		return false
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", WebhookURL, bytes.NewBuffer(payload))
	if err != nil {
		logger.LogErrorToFile("Error creating request:", err)
		return false
	}

	// Set the content type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Logger.LogErrorToFile("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusNoContent {
		Logger.LogErrorToFile("Message sent to Discord webhook successfully!")
		return true
	} else {
		Logger.LogErrorToFile("Unexpected response status:", resp.Status)
		return false
	}
}
