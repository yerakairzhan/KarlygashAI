package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log"
	"rateme/handlers"
	"rateme/utils"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Ошибка загрузки конфигурации:", err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	handlers.SetupHandlers(bot)
}
