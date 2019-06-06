package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"
	TGBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

func crashReport(updateError TGBotAPI.Update) {

	if err := recover(); err != nil {
		f, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		logger := log.New(f, "[ERROR]", log.LstdFlags)

		logger.Println(err)
		logger.Println("Message")
		logger.Println(spew.Sdump(updateError))
		logger.Println("Stack:")
		logger.Println(string(debug.Stack()))
		logger.Println("===================================")
	}

}
