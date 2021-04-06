package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"sheet"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("<your bot token>")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on accont %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("<your webhook url>", nil))

	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	fmt.Printf("\n%v\n", info)
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":<port that you open>", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := int64(update.Message.Chat.ID)
		if update.Message.From.ID != <your UID(int)> {
			fmt.Println("From other user")
			msg := tgbotapi.NewMessage(chatID, "Thanks for your testing")
			bot.Send(msg)
			bot.Send(tgbotapi.NewMessage(int64(<your UID>), "Other user is using this bot"))
			continue
		}
		input := update.Message.Text
		input_list := strings.Split(input, " ")
		if input_list[0] == "/add" || input_list[0] == "/add@<bot username>" {
			if len(input_list) == 5 {
				fmt.Println(input_list)
				Category := input_list[0]
				Type := input_list[1]
				Tag := input_list[2]
				Price := input_list[3]
				sheet.Fill_the_sheet(Category, Type, Tag, Price)
				fmt.Println("Update successfully")
				msg := tgbotapi.NewMessage(int64(userID), "Update successfully")
				bot.Send(msg)
			} else {
				fmt.Println("Invalid input")
				msg := tgbotapi.NewMessage(int64(userID), "Invalid input")
				bot.Send(msg)
			}
		}
	}

}
