package inline

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var CategoryKeyboard = tgbotapi.NewInlineKeyboardMarkup (
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("food", "food"),
		tgbotapi.NewInlineKeyboardButtonData("traffic", "traffic"),
		tgbotapi.NewInlineKeyboardButtonData("sports", "sports"),
	),
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("clothing", "clothing"),
		tgbotapi.NewInlineKeyboardButtonData("gift", "gift"),
	),
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("household supplies", "household supplies"),
	),
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("entertainment", "entertainment"),
	),
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("snacks", "snacks"),
		tgbotapi.NewInlineKeyboardButtonData("others", "others"),
	),
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("cancel", "cancel"),
	),
)

var TypeKeyboard = tgbotapi.NewInlineKeyboardMarkup (
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("need", "need"),
		tgbotapi.NewInlineKeyboardButtonData("want", "want"),
		tgbotapi.NewInlineKeyboardButtonData("others", "others"),
	),
	tgbotapi.NewInlineKeyboardRow (
		tgbotapi.NewInlineKeyboardButtonData("back to category", "back"),
		tgbotapi.NewInlineKeyboardButtonData("cancel", "cancel"),
	),
)

