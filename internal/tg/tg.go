package tg

import (
	"log"
	"os"
	"secretary/internal/tg/handler"

	tgbotapi "github.com/iisakov/telegram-bot-api"
	"github.com/joho/godotenv"
)

func NewTg() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Не найден файл .env")
	}

	myBot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	myBot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := myBot.GetUpdatesChan(u)
	for update := range updates {
		switch {
		case update.Message.IsCommand():
			handler.ComandHandle(update.Message, myBot)
		default:
			handler.MessageHandle(update.Message, myBot)
		}

	}
}
