package controller

import (
	"tgbotimgstor/internal/loader"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	loader loader.LoadService
}

func New() Controller {
	return Controller{}
}

func (c Controller) HandleMessage(msg tgbotapi.Message) (*tgbotapi.MediaGroupConfig, *tgbotapi.MessageConfig) {
	groupName := msg.CommandArguments()

	switch msg.Command() {
	case "load":
		fileNames, err := loader.Load(groupName)
		if err != nil {
			errMsg := tgbotapi.NewMessage(msg.Chat.ID, string(err))
			return nil, &errMsg
		}

		photos := []any{}

		for _, filePath := range fileNames {
			photos = append(photos, tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(filePath)))
		}

		group := tgbotapi.NewMediaGroup(msg.Chat.ID, photos)
		return &group, nil

	default:
		errMsg := tgbotapi.NewMessage(msg.Chat.ID, "Undefined command")
		return nil, &errMsg
	}
}
