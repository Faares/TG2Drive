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
		f, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

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

func notAuthorizedReport(update TGBotAPI.Update) {
	f, err := os.OpenFile("not_authorized.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	logger := log.New(f, "[Security]", log.LstdFlags)
	logger.Println("Not Authorized Access")
	logger.Println("data:")
	logger.Println(spew.Sdump(update))
}
