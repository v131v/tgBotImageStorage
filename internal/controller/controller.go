package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller interface {
	HandleMessage(*tgbotapi.Message) tgbotapi.Chattable
}
