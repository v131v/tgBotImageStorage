package controller

import (
	"tgbotimgstor/internal/loader"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	loader *loader.Loader
}

func New(loader *loader.Loader) Controller {
	return Controller{loader}
}

func (c Controller) HandleMessage(msg *tgbotapi.Message) tgbotapi.Chattable {
	groupName := msg.CommandArguments()

	switch msg.Command() {
	case "load":
		fileNames, err := loader.Get(groupName)
		if err != nil {
			errMsg := tgbotapi.NewMessage(msg.Chat.ID, string(err))
			return errMsg
		}

		photos := []any{}

		for _, filePath := range fileNames {
			photos = append(photos, tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(filePath)))
		}

		group := tgbotapi.NewMediaGroup(msg.Chat.ID, photos)
		return group

	default:
		errMsg := tgbotapi.NewMessage(msg.Chat.ID, "Undefined command")
		return errMsg
	}
}
