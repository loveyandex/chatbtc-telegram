package main

import (
	 
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/loveyandex/chatbtc-telegram.git/flow"
)
 

func main() {

	bot, err := tgbotapi.NewBotAPI("6279508336:AAGxO1241tiI9rQ31_-77JS8aXoJjkpIzXI")
	if err != nil {
		log.Panic(err)
	}  

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

xxxxxx:
	for update := range updates {

		for _, t := range flow.NewCircuit() {
			if update.Message != nil {

				if update.Message.Text != "" {
					for _, f := range t.TextFlows {
						if f.Name == update.Message.Text {
							f.Func(&update, bot)
							continue xxxxxx
						}
					}

				}
				if update.Message.Contact != nil {
					if t.ContactFlow != nil {
						t.ContactFlow.Func(&update, bot, update.Message.Contact)
						continue xxxxxx
					}
				}
			}

			if update.InlineQuery != nil {
				if t.InlineQueryFLow.Splits == len(strings.Split(update.InlineQuery.Query, " ")) {
					t.InlineQueryFLow.Func(&update, bot, update.InlineQuery.Query)
					continue xxxxxx
				}

			}
			if update.CallbackQuery != nil {

				for _, isf := range t.PublicStringIncludeFlows {
					if strings.Contains(update.CallbackQuery.Data, (isf.Include)) {
						isf.Func(&update, bot, update.CallbackQuery.Data)
						continue xxxxxx
					}
				}
 

			}
		}

		if update.CallbackQuery != nil {
			 

		}
		 

		//no neeed to user state-transitions
		if update.CallbackQuery != nil {
 

		}

		//user need automata
		if update.SentFrom() == nil {
			continue xxxxxx

		} 

	}

}
