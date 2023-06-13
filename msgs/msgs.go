package msgs

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var START_KEYBOARD = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact("send confirm sms"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("تبادل"),
	), 
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ُSupport"),
	),
)
