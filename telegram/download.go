package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DownloadFromTelegram(fileID string) string {

	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(
		os.Getenv("TELEGRAM_BOT_TOKEN"),
		os.Getenv("TELEGRAM_BOT_API"),
	)

	if err != nil {
		log.Panic(err)
	}

	if res, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID}); err != nil {

		log.Println(err)

	} else {

		return res.FilePath
	}

	return ""
}
