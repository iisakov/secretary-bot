package handler

import (
	"fmt"
	"secretary/internal/tg/model"
	"secretary/internal/tg/stl"

	tgbotapi "github.com/iisakov/telegram-bot-api"
)

func ComandHandle(c *tgbotapi.Message, t *tgbotapi.BotAPI) {

	switch c.Command() {
	case "start":
		uName := c.From.UserName
		if c.From.FirstName != "" {
			uName = c.From.FirstName
		}
		text := fmt.Sprintf(
			"Привет, %s.\nЯ Секретарь от компании [by_Artisan].\n\nВот что я умею:\n%s",
			uName,
			model.Skills)
		stl.SendText(t, c.From.ID, text)
	}
}
