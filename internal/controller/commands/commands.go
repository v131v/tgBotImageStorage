package commands

import (
	"tgbotimgstor/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandsController struct {
	service service.Service
}

func New(service service.Service) *CommandsController {
	return &CommandsController{service}
}

func (c *CommandsController) HandleMessage(msg *tgbotapi.Message) tgbotapi.Chattable {
	groupName := msg.CommandArguments()

	switch msg.Command() {
	case "load":
		fileNames, err := c.service.Get(groupName)
		if err != nil {
			errMsg := tgbotapi.NewMessage(msg.Chat.ID, err.Error())
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
