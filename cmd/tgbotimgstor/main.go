package main

import (
	"log"
	"os"
	"tgbotimgstor/internal/controller/commands"
	"tgbotimgstor/internal/server"
	"tgbotimgstor/internal/service/loader"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	storagePath := os.Getenv("STORAGE_PATH")
	botToken := os.Getenv("BOT_TOKEN")

	ls := loader.New(storagePath)
	ctrl := commands.New(ls)
	serv, err := server.New(ctrl, botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Started")
	serv.Run()
}
