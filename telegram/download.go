package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DownloadFromTelegram(fileID string) string {

	bot, err := tgbotapi.NewBotAPI(
		os.Getenv("TELEGRAM_BOT_TOKEN"),
	)
	if err != nil {
		log.Panic(err)
	}

	if res, err := bot.GetFileDirectURL(fileID); err != nil {

		log.Println(err)

	} else {

		return res
	}

	return ""
}
