package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strconv"

	TGBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

var ConfigFile *string
var BotConfig Config
var bot *TGBotAPI.BotAPI
var Authorized map[string]string
var PhotosToSave []TGFile
var VideoToSave []TGFile

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

		if update.Message.Photo != nil || update.Message.Video != nil {

			// if photo
			if update.Message.Photo != nil {
				photos := (*update.Message.Photo)
				biggestPhoto := &photos[len(photos)-1]
				var photo TGFile = TGPhoto{
					ID:     biggestPhoto.FileID,
					Height: biggestPhoto.Height,
					Width:  biggestPhoto.Width,
					Size:   biggestPhoto.FileSize}

				PhotosToSave = append(PhotosToSave, photo)
				continue
			}

			if update.Message.Video != nil {
				updateVideo := (*update.Message.Video)
				var video TGFile = TGVideo{
					ID:       updateVideo.FileID,
					Height:   updateVideo.Height,
					Width:    updateVideo.Width,
					MimeType: updateVideo.MimeType,
					Duration: updateVideo.Duration,
					Thumb: TGPhoto{
						ID:     updateVideo.Thumbnail.FileID,
						Height: updateVideo.Thumbnail.Height,
						Width:  updateVideo.Thumbnail.Width,
						Size:   updateVideo.Thumbnail.FileSize},
					Size: updateVideo.FileSize}

				VideoToSave = append(VideoToSave, video)
				continue
			}

		} else {
			// the idea:  execute if /upload command. but temporary use this, just for development
			if PhotosToSave != nil {
				go saveJSON(&PhotosToSave, "photos.json")
			}

			if VideoToSave != nil {
				go saveJSON(&VideoToSave, "videos.json")
			}

			bot.Send(TGBotAPI.NewMessage(update.Message.Chat.ID, "No Photo? No Video? .. Hmm"))
			continue
		}

	}
}

// pass by address, to clear the slice
func saveJSON(slice *[]TGFile, name string) {

	x, err := json.Marshal(slice)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(name, x, 0644)
	if err != nil {
		panic(err)
	}
	// clear the slice, after save it..
	*slice = nil
}
