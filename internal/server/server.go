package server

import (
	"tgbotimgstor/internal/controller"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Server struct {
	bot  *tgbotapi.BotAPI
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller, token string) (s Server, err error) {
	s.bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}
	s.ctrl = ctrl
	return
}

func (s *Server) Run() {

	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = 30

	updates := s.bot.GetUpdatesChan(updCfg)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		resp := (*s.ctrl).HandleMessage(update.Message)
		s.bot.Send(resp)
	}
}
