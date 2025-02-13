package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"rateme/utils"
	"time"
)

type ContentItem struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageURL struct {
		URL string `json:"url"`
	} `json:"image_url,omitempty"`
}

type Message struct {
	Role    string        `json:"role"`
	Content []ContentItem `json:"content"`
}

type AIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type AIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Function to call OpenRouter API and remember the conversation
func CallOpenRouterAPI(prompt string, conversationHistory *[]Message) (string, error) {
	log.Println("Conversation History:", conversationHistory)
	client := resty.New()
	config, _ := utils.LoadConfig()

	// Create the request body with the dynamic prompt and conversation history
	request := AIRequest{
		Model: config.Model,
		Messages: append(*conversationHistory, Message{
			Role: "user",
			Content: []ContentItem{
				{
					Type: "text",
					Text: prompt,
				},
			},
		}),
	}
	// Convert the request struct to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Retry logic
	var aiResponse AIResponse
	var resp *resty.Response
	retries := 3

	for i := 0; i < retries; i++ {
		// Make the POST request with a timeout
		resp, err = client.R().
			SetHeader("Authorization", "Bearer "+config.DeepSeekAPIKey).
			SetHeader("Content-Type", "application/json").
			SetBody(jsonData).
			Post(config.AI_URL)

		if err == nil && resp.StatusCode() == 200 {
			break
		}

		// Log error and retry after a short delay
		log.Printf("Error calling OpenRouter API (attempt %d/%d): %v\n", i+1, retries, err)
		time.Sleep(1 * time.Second)
	}

	if err != nil || resp.StatusCode() != 200 {
		log.Println(err)
		return "Sorry, I couldn't get a response right now. Please try again later.", nil
	}

	// Parse the response body
	if err := json.Unmarshal(resp.Body(), &aiResponse); err != nil {
		return "", fmt.Errorf("failed to parse API response: %w", err)
	}
	log.Println(string(jsonData))
	log.Println(resp.String())
	if len(aiResponse.Choices) == 0 {
		log.Println(err)
		return "Sorry, I couldn't get a response from the assistant. Please try again later.", nil
	}

	// Add AI response to the conversation history
	*conversationHistory = append(*conversationHistory, Message{
		Role: "assistant",
		Content: []ContentItem{
			{
				Type: "text",
				Text: aiResponse.Choices[0].Message.Content,
			},
		},
	})

	// Return the AI's response
	return aiResponse.Choices[0].Message.Content, nil
}
