package flow

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" 
)

var GeneralFlow Transsistor = Transsistor{
	StateActionFlows:         []StateActionFlow{},
	PublicIncludeFlows:       []IncludeFlow{},
	PublicStringIncludeFlows: []IncludeStringFlow{},
	TextFlows: []Flow{

		Flow{
			Name: "/start",
			Func: func(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				msg.Text = "Your wallet options:" 
				bot.Send(msg)

			},
		},
		Flow{
			Name: "/inl",
			Func: func(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				msg.Text = "select one option"

				var keyboard [][]tgbotapi.InlineKeyboardButton

				var ikb2 []tgbotapi.InlineKeyboardButton

				ikb := tgbotapi.NewInlineKeyboardButtonURL("utl", "https://google.com")
				ikb2 = append(ikb2, ikb)

				ikb_ := tgbotapi.NewInlineKeyboardButtonSwitch("switch", "233 usdt give")
				ikb2 = append(ikb2, ikb_)

				jj := ("1 btc give")
				ikb__ := tgbotapi.InlineKeyboardButton{
					Text:                         "swcurrent",
					SwitchInlineQueryCurrentChat: &jj,
				}
				ikb2 = append(ikb2, ikb__)

				keyboard = append(keyboard, ikb2)

				msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{
					InlineKeyboard: keyboard,
				}

				bot.Send(msg)

			},
		},
	},
	ContactFlow:     &ContactFlow{},
	InlineQueryFLow: InlineQueryFlow{},
}
