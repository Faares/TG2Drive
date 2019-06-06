package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	TGBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

func auth(update TGBotAPI.Update) (isAuth bool, user User) {

	// check if message from authirzed users and not null
	if _, ok := BotConfig.Telegram.Authorized[update.Message.From.UserName]; !ok || update.Message == nil {
		notAuthorizedReport(update)
		fmt.Println("Not Authorized")
		return false, User{}
	}

	// We know now the user is allowed to contact the bot, then let's get it data.
	userData := BotConfig.Telegram.Authorized[update.Message.From.UserName]

	// require password
	if update.Message.Command() == "login" {

		h := sha256.New()
		h.Write([]byte(update.Message.CommandArguments()))

		if hex.EncodeToString(h.Sum(nil)) == userData.Password {
			Authorized[strconv.FormatInt(update.Message.Chat.ID, 10)] = userData.Name
			message := TGBotAPI.NewMessage(update.Message.Chat.ID, fmt.Sprintf("DONE, I KNOW U NOW! Welcome %s", userData.Name))
			bot.Send(message)

			return true, userData
		} else {
			message := TGBotAPI.NewMessage(update.Message.Chat.ID, "WRONG PASSOWRD!!")
			bot.Send(message)
		}
	}

	// if no password .. please auth, or go to hell
	if _, ok := Authorized[string(update.Message.Chat.ID)]; !ok {
		message := TGBotAPI.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Greeting %s, Please use command /login to authenticate u..", userData.Name))
		bot.Send(message)
	}

	return false, userData
}
