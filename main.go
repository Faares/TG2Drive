package main

import (
	"flag"
	"fmt"
	"strconv"

	TGBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

var ConfigFile *string
var BotConfig Config
var bot *TGBotAPI.BotAPI
var Authorized map[string]string

func main() {

	Authorized = make(map[string]string)

	ConfigFile = flag.String("config", "./Config.json", "Config File")
	flag.Parse()
	BotConfig = getConfig(*ConfigFile)

	// Initalize bot
	var err error
	bot, err = TGBotAPI.NewBotAPI(BotConfig.Telegram.Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := TGBotAPI.NewUpdate(0)

	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		defer crashReport(update)

		if _, ok := Authorized[strconv.FormatInt(update.Message.Chat.ID, 10)]; !ok {
			auth(update)
			continue
		}

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

		bot.Send(TGBotAPI.NewMessage(update.Message.Chat.ID, "No Photo? No Video? .. Hmm"))

	}
}
