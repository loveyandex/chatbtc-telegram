package flow

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" 
)

type Message struct {
	Update tgbotapi.Update
	Bot    *tgbotapi.BotAPI
}
type Flow struct {
	Name string
	Func func(update *tgbotapi.Update, bot *tgbotapi.BotAPI)
}
type StateFlow struct {
	State string
	Func  func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, chatId int64, text string)
}

type RegexFlow struct {
	Regex string

	Func func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, chatId int64, text string)
}

type IncludeFlow struct {
	Include string
	Func    func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, include string)
}

type IncludeStringFlow struct {
	Include string
	Func    func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, include string)
}

type ContactFlow struct {
	Func func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, contact *tgbotapi.Contact)
}

type StateActionFlow struct {
	TextFunc          func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, text string)
	CallbackQueryFunc func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, data string)
}

type InlineQueryFlow struct {
	Splits int
	Func   func(update *tgbotapi.Update, bot *tgbotapi.BotAPI, query string)
}

type Transsistor struct {
	StateActionFlows         []StateActionFlow
	PublicIncludeFlows       []IncludeFlow
	PublicStringIncludeFlows []IncludeStringFlow
	TextFlows                []Flow
	ContactFlow              *ContactFlow
	InlineQueryFLow          InlineQueryFlow
}

var Transsistors []Transsistor

func NewCircuit() []Transsistor {

	//add your transsistor
	Transsistors = append(Transsistors, GeneralFlow)
	 

	return Transsistors
}
