package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatid int64, userid int64, command string) {
	switch command {
	case "start":
		text := "👋 <b>Здравствуйте!</b>\nЯ ваш виртуальный помощник 📱 для отзывов и предложений о нашем заведении.\n\n<b>Что я могу для вас сделать?</b>\n✨ Оставить отзыв\n⚡ Пожаловаться\n🌟 Поделиться впечатлением\n\n<b>Ваше мнение важно для нас!</b> 🙌 Все сообщения <b>100% анонимны</b>.\n\n<b>Начнем?</b> Напишите своё сообщение здесь! ✍️"
		text = "Привет давай познакомимся!"
		msg := tgbotapi.NewMessage(chatid, text)
		msg.ParseMode = "HTML"
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
