package main

import (
	"log"

	"github.com/halushko/core-go/logger"
	"github.com/halushko/tg-bot-go/bot"
)

func main() {
	logFile := logger.SoftPrepareLogFile("bot.log")
	b, err := bot.Create()

	if err != nil {
		log.Printf("Error: %v", err)
	}

	b.StartHandleTextMessages()
	b.StartSendTextMessages()
	b.StartHandleMemberJoined()

	log.Println("Бота запущено")
	b.Start()
	logger.SoftLogClose(logFile)
}
