package stl

import (
	tgbotapi "github.com/iisakov/telegram-bot-api"
)

func SendText(t *tgbotapi.BotAPI, chatId int64, msgText ...string) error {
	text := ""
	for _, t := range msgText {
		text = text + "\n" + t
	}
	msg := tgbotapi.NewMessage(chatId, text)

	_, err := t.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func SendDocument(t *tgbotapi.BotAPI, chatId int64, docPath string) error {
	msg := tgbotapi.NewDocument(chatId, tgbotapi.FilePath(docPath))

	_, err := t.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
