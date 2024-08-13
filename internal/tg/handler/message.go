package handler

import (
	"fmt"
	"path"
	"regexp"
	"secretary/internal/blueprint"
	"secretary/internal/operator"
	"secretary/internal/printer"
	"secretary/internal/tg/stl"
	"strings"

	tgbotapi "github.com/iisakov/telegram-bot-api"
)

func MessageHandle(m *tgbotapi.Message, t *tgbotapi.BotAPI) {
	if m.Text == "DEFAULT" {
		text := fmt.Sprintf("Файл %s.pdf готовится\n", m.Text)
		stl.SendText(t, m.From.ID, text)

		bp := blueprint.NewDefaultBluprint()
		p := printer.NewDefaultPrinter(bp.GetOptions())
		o := operator.NewDefaultOperator(p, bp)
		o.UseBluprint()

		stl.SendDocument(t, m.From.ID, "../../out/DefaultDoc/Default.pdf")
		text = fmt.Sprintf("Файл %s.pdf готов\n", m.Text)
		stl.SendText(t, m.From.ID, text)

		return
	}

	var mi, text string

	switch {
	case strings.Contains(m.Text, "fgis.gost.ru/fundmetrology/cm/results/"):
		mi = path.Base(m.Text)
	default:
		mi = m.Text
	}

	if matched, _ := regexp.MatchString(`^[0-9]-[0-9]{1,10}$`, mi); matched {
		text = fmt.Sprintf("Файл %s.pdf готовится\n", mi)
	} else {
		text = fmt.Sprintf("Простите, но %s не выглядит как номер сертификата\n", mi)
	}

	stl.SendText(t, m.From.ID, text)

}
