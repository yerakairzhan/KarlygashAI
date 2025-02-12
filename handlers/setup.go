package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"rateme/api"
	db "rateme/db/sqlc"
	"rateme/utils"
)

func SetupHandlers(bot *tgbotapi.BotAPI, queries *db.Queries) {
	userConversationHistory := make(map[int64][]api.Message)
	config, _ := utils.LoadConfig()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		var chatID int64
		var userID int64
		var command string

		if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
			userID = update.CallbackQuery.From.ID
			command = update.CallbackQuery.Data
		} else if update.Message != nil {
			chatID = update.Message.Chat.ID
			userID = update.Message.From.ID

			if update.Message.IsCommand() {
				command = update.Message.Command()
				handleCommand(bot, update, queries, chatID, userID, command)
			} else {
				conversationHistory, exists := userConversationHistory[userID]
				if !exists {
					conversationHistory = []api.Message{
						{
							Role: "user",
							Content: []api.ContentItem{
								{Type: "text", Text: config.Prompt},
							},
						},
					}
				}

				userMessage := api.Message{
					Role: "user",
					Content: []api.ContentItem{
						{Type: "text", Text: update.Message.Text},
					},
				}
				conversationHistory = append(conversationHistory, userMessage)

				response, err := api.CallOpenRouterAPI(update.Message.Text, &conversationHistory)
				if err != nil {
					log.Fatalf("Error here: %v", err)
				}

				aiMessage := api.Message{
					Role: "assistant",
					Content: []api.ContentItem{
						{Type: "text", Text: response},
					},
				}
				conversationHistory = append(conversationHistory, aiMessage)

				userConversationHistory[userID] = conversationHistory

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			}
		}
	}
}
