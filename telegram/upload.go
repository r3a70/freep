package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UploadToTelegram(file string) string {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	newFile := tgbotapi.FilePath(file)
	newDoc := tgbotapi.NewDocument(1302633753, newFile)
	if res, err := bot.Send(newDoc); err != nil {
		log.Println(err)
	} else {
		// url, _ := bot.GetFileDirectURL(res.Document.FileID)
		return "http://0.0.0.0:8000/download/" + res.Document.FileID
	}

	return ""
}
