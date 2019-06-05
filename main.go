package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"

	TGBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

func crashReport(updateError TGBotAPI.Update) {
	f, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	logger := log.New(f, "[ERROR]", log.LstdFlags)

	if err := recover(); err != nil {
		logger.Println(err)
		logger.Println("Message")
		logger.Println(spew.Sdump(updateError))
		logger.Println("Stack:")
		logger.Println(string(debug.Stack()))
		logger.Println("===================================")
	}

}

func main() {

	Authrized := make(map[string]string)

	config := getConfig()

	// Initalize bot
	bot, err := TGBotAPI.NewBotAPI(config.Telegram.Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := TGBotAPI.NewUpdate(0)

	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		log.Println("Start")

		defer crashReport(update)

		// check if message from authirzed users and not null
		if _, ok := config.Telegram.Authrized[update.Message.From.UserName]; !ok || update.Message == nil {
			log.Println("Not Authrized")
			continue
		}

		// We know now the user is allowed to contact the bot, then let's get it data.
		userData := config.Telegram.Authrized[update.Message.From.UserName]

		// require password
		if update.Message.Command() == "login" {

			h := sha256.New()
			h.Write([]byte(update.Message.CommandArguments()))

			if hex.EncodeToString(h.Sum(nil)) == userData.Password {
				Authrized[string(update.Message.Chat.ID)] = userData.Name
				message := TGBotAPI.NewMessage(update.Message.Chat.ID, fmt.Sprintf("DONE, I KNOW U NOW! Welcome %s", userData.Name))
				bot.Send(message)
			} else {
				message := TGBotAPI.NewMessage(update.Message.Chat.ID, "WRONG PASSOWRD!!")
				bot.Send(message)
			}

			continue
		}

		// if no password .. please auth, or go to hell
		if _, ok := Authrized[string(update.Message.Chat.ID)]; !ok {
			message := TGBotAPI.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Greeting %s, Please use command /login to authicate u..", userData.Name))
			bot.Send(message)
			continue
		}

		/**
		@TODO Check Message Content: does it image?

		*/

		if update.Message.Photo != nil {
			photos := (*update.Message.Photo)
			biggestPhoto := &photos[len(photos)-1]
			var photo TGFile = TGPhoto{
				ID:     biggestPhoto.FileID,
				Height: biggestPhoto.Height,
				Width:  biggestPhoto.Width,
				Size:   biggestPhoto.FileSize}

			fmt.Println(photo)
		}

	}
}