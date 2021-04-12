package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"TGBot/inline"
	"TGBot/sheet"
)

type TeleBot struct {
        botAPI *tgbotapi.BotAPI
}

var Category string
var Type string
var Tag string
var Price string

var Stage map[string]bool = map[string]bool {
	"add": false,
	"category": false,
	"type": false,
	"tag": false,
	"price": false,
}

var Bot_info struct {
	Token string
	Username string
	WebhookURL string
	WebhookPort string
}

var Admin struct {
	ID int64
}


func bot_init()(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	Bot_info.Token = os.Getenv("BotToken")
	Bot_info.Username = os.Getenv("BotUserName")
	Bot_info.WebhookURL = os.Getenv("WebhookURL")
	Bot_info.WebhookPort = ":" + os.Getenv("WebhookPort")

	Admin.ID, err = strconv.ParseInt(os.Getenv("UserID"), 10, 64)
	if err != nil {
		log.Println(err)
	}

	bot, err = tgbotapi.NewBotAPI(Bot_info.Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(Bot_info.WebhookURL, nil))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
        if info.LastErrorDate != 0 {
                log.Printf("[Telegram callback failed]%s\n\n", info.LastErrorMessage)
        }
	updates = bot.ListenForWebhook("/")
	go http.ListenAndServe(Bot_info.WebhookPort, nil)

	return bot, updates
}

func (t *TeleBot)sendMessage(ChatID int64, m string) {
        msg := tgbotapi.NewMessage(ChatID, m)
        t.botAPI.Send(msg)
}

func (t *TeleBot)createInlineKeyboardMarkup(ChatID int64) {
        msg := tgbotapi.NewMessage(ChatID, "Please choose a category")
	Stage["category"] = true
        msg.ReplyMarkup = inline.CategoryKeyboard
        t.botAPI.Send(msg)
}

func (t *TeleBot)inlineKeybaordHandler(ChatID int64, MessageID int, CallBackData string){
        log.Println("CallBackData:", CallBackData)
        if CallBackData ==  "cancel" {
		for key, _ := range Stage {
			Stage[key] = false
		}
                edited_text := tgbotapi.NewEditMessageText (
                        ChatID,
                        MessageID,
                        "cancel this operation",
                )
                t.botAPI.Send(edited_text)
	} else if CallBackData == "back"{
		Stage["category"] = true
		edited_str := "Please choose a category"
		edited_text := tgbotapi.NewEditMessageText (
			ChatID,
			MessageID,
			edited_str,
		)
		edited_text.ReplyMarkup = &inline.CategoryKeyboard
		t.botAPI.Send(edited_text)
	} else if Stage["category"] {
		Stage["category"] = false
		Stage["type"] = true
		Category = CallBackData
		edited_str :=	"Category: " + Category + "\n" +
				"Please choose the type"
                edited_text := tgbotapi.NewEditMessageText (
                        ChatID,
                        MessageID,
                        edited_str,
                )
                edited_text.ReplyMarkup = &inline.TypeKeyboard
                t.botAPI.Send(edited_text)
        } else if Stage["type" ]{
		Type = CallBackData
		Stage["type"] = false
		Stage["tag"] = true
		edited_str :=	"Category: " + Category + "\n" +
				"Type: " + Type + "\n" +
				"Please enter a tag"
		edited_text := tgbotapi.NewEditMessageText (
			ChatID,
			MessageID,
			edited_str,
		)
		t.botAPI.Send(edited_text)
	} else {
		edited_str := "The instruction is canceled.\nPlease enter again."
		edited_text := tgbotapi.NewEditMessageText (
			ChatID,
			MessageID,
			edited_str,
		)
		t.botAPI.Send(edited_text)
	}
}

func main() {
	bot, updates := bot_init()

	teleBot := TeleBot {
		botAPI: bot,
	}

	for update := range updates {
		fmt.Printf("%t\n\n\n", update)
		if update.Message == nil {
			if update.CallbackQuery != nil {
				log.Println(update.CallbackQuery)
				chatID := update.CallbackQuery.Message.Chat.ID
				messageID := update.CallbackQuery.Message.MessageID
                                data := update.CallbackQuery.Data
                                teleBot.inlineKeybaordHandler(chatID, messageID, data)
			}
			continue
		}

		log.Println(update.Message)

		chatID := int64(update.Message.Chat.ID)
		if chatID != Admin.ID {
			fmt.Println("From other user")
			msg := tgbotapi.NewMessage(chatID, "Thanks for your testing")
			bot.Send(msg)
			bot.Send(tgbotapi.NewMessage(Admin.ID, "Other user is using this bot"))
			continue
		}
		input := update.Message.Text
		if input == "/cancel" || input == "/cancel" + Bot_info.Username {
			for key, _ := range Stage {
				Stage[key] = false
			}
			teleBot.sendMessage(update.Message.Chat.ID, "All instruction is canceled")
		} else if input == "/add" || input == "/add" + Bot_info.Username {
			for key, _ := range Stage {
				Stage[key] = false
			}
			Stage["add"] = true
			teleBot.createInlineKeyboardMarkup(update.Message.Chat.ID)
		} else if Stage["tag"] {
			Stage["tag"] = false
			Stage["price"] = true
			Tag = update.Message.Text
			str := fmt.Sprintf("Category: %s\nType: %s\nTag: %s\nPlesase enter the price", Category, Type, Tag)
			teleBot.sendMessage(update.Message.Chat.ID, str)
		} else if Stage["price"] {
			Stage["price"] = false
			Price = update.Message.Text
			str := fmt.Sprintf("Category: %s\nType: %s\nTag: %s\nPrice: %s\nUpdating...", Category, Type, Tag, Price)
			teleBot.sendMessage(update.Message.Chat.ID, str)
			sheet.Fill_the_sheet(Category, Type, Tag, Price)
			teleBot.sendMessage(update.Message.Chat.ID, "Update successfully")
		} else {
			teleBot.sendMessage(update.Message.Chat.ID, "Please choose a instruction")
		}
	}
}
