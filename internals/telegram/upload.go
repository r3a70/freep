package telegram

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UploadToTelegram(file string) string {

	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(
		os.Getenv("TELEGRAM_BOT_TOKEN"),
		os.Getenv("TELEGRAM_BOT_API"),
	)
	if err != nil {
		log.Panic(err)
	}

	newFile := tgbotapi.FilePath(file)
	newDoc := tgbotapi.NewDocument(1302633753, newFile)
	fmt.Println(newDoc)
	if res, err := bot.Send(newDoc); err != nil {
		log.Println(err)
	} else {

		var fileID string
		if res.Animation != nil {
			fileID = res.Animation.FileID
		} else if res.Audio != nil {
			fileID = res.Audio.FileID
		} else if res.Document != nil {
			fileID = res.Document.FileID
		} else if res.Photo != nil {
			fileID = res.Photo[0].FileID
		} else if res.Video != nil {
			fileID = res.Video.FileID
		}

		return os.Getenv("BASE_URL") + "/download/" + fileID
	}

	return ""
}
