package main

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log"
	db "rateme/db/sqlc"
	"rateme/handlers"
	"rateme/utils"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Ошибка загрузки конфигурации:", err)
		return
	}

	conn, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	db := db.New(conn)

	handlers.SetupHandlers(bot, db)
}
