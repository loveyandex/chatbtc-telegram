package msgs

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendColumnsInline(texts []string) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < len(texts); i++ {

		var ikb2 []tgbotapi.InlineKeyboardButton
		for j := 0; j < 1; j++ {

			ikb := tgbotapi.NewInlineKeyboardButtonData(texts[i], texts[i])
			ikb2 = append(ikb2, ikb)

		}

		keyboard = append(keyboard, ikb2)

	}

	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

}

func SendColumnsInlineWithKey(texts []string, keys []string) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < len(texts); i++ {

		var ikb2 []tgbotapi.InlineKeyboardButton
		for j := 0; j < 1; j++ {

			ikb := tgbotapi.NewInlineKeyboardButtonData(texts[i], keys[i])
			ikb2 = append(ikb2, ikb)

		}

		keyboard = append(keyboard, ikb2)

	}

	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

}

func SendRowsInlineWithKey(texts []string, keys []string) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < 1; i++ {

		var ikb2 []tgbotapi.InlineKeyboardButton
		for j := 0; j < len(texts); j++ {

			ikb := tgbotapi.NewInlineKeyboardButtonData(texts[j], keys[j])
			ikb2 = append(ikb2, ikb)

		}

		keyboard = append(keyboard, ikb2)

	}

	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

}

func NewSotoon(sooon *Sotoon) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for _, r := range sooon.Rows {

		var ikb2 []tgbotapi.InlineKeyboardButton
		for j := 0; j < len(r.Keys); j++ {

			ikb := tgbotapi.NewInlineKeyboardButtonData(r.Vals[j], r.Keys[j])
			ikb2 = append(ikb2, ikb)

		}

		keyboard = append(keyboard, ikb2)

	}

	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

}

type Sotoon struct {
	Rows []Row
}
type Row struct {
	Vals []string
	Keys []string
}

func NewRect(row Row, m,n int) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	var k = 0

	for i := 0; i < m; i++ {
		var ikb2 []tgbotapi.InlineKeyboardButton
		for j := 0; j < n; j++ {
			ikb := tgbotapi.NewInlineKeyboardButtonData(row.Vals[k], row.Keys[k])
			ikb2 = append(ikb2, ikb)
			k++
			if k == len(row.Keys) {
				break
			}

		}
		keyboard = append(keyboard, ikb2)
		if k == len(row.Keys) {
			break
		}
	}

	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

}

var OrderTypes = SendColumnsInlineWithKey([]string{"خرید آنی با بهترین قیمت", "خرید با قیمت"}, []string{fmt.Sprint(MARKETTRADETYPE), fmt.Sprint(OrderLimitTYPE)})
var VerifyOrderPriceAndamount = SendColumnsInlineWithKey([]string{"تایید", "بازگشت"}, []string{fmt.Sprint(VERIFYPRICEANDAMOUNT), fmt.Sprint(BACKTOSRCAMOUNT)})
var LetsPayIRTLInline = SendColumnsInlineWithKey([]string{"بزن بریم", "< برگشت"}, []string{fmt.Sprint(LetsPayIRTL), fmt.Sprint(BACKTOSRCAMOUNT)})
var CancelOrderInline = SendColumnsInlineWithKey([]string{"< برگشت"}, []string{fmt.Sprint(CancelOrder)})

var BackTOMAINMENUInline = SendColumnsInlineWithKey([]string{"< برگشت"}, []string{fmt.Sprint(BackTOMAINMENU)})

var ConfirmUserForwardedForTransfernline = SendColumnsInlineWithKey([]string{"تایید کاربر مقصد", "< برگشت"}, []string{fmt.Sprint(ConfirmUserForwardedForTransfer), fmt.Sprint(BackTOMAINMENU)})
var ConfirmForTransferingInline = SendColumnsInlineWithKey([]string{"تایید نهایی انتقال", "لفو"}, []string{fmt.Sprint(ConfirmForTransfering), fmt.Sprint(CancelTransfering)})

var WalletsInline = SendRowsInlineWithKey([]string{"USDT", "IRT"}, []string{fmt.Sprint(USDT), fmt.Sprint(IRT)})
var MyWallet = NewSotoon(&Sotoon{
	Rows: []Row{
		Row{
			Vals: []string{"واریز کردن", "برداشت"},
			Keys: []string{fmt.Sprint(USDT), fmt.Sprint(IRT)},
		},
		Row{
			Vals: []string{"تبادل"},
			Keys: []string{fmt.Sprint(Tabadol)},
		},
		Row{
			Vals: []string{"خرید تتر با بهترین قیمت"},
			Keys: []string{fmt.Sprint(BuyTetherByBestPrice)},
		},
		Row{
			Vals: []string{"< برگشت"},
			Keys: []string{fmt.Sprint(CancelOrder)},
		},
	},
})

func NewMyWallet() *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	//row 1
	var ikb2 []tgbotapi.InlineKeyboardButton
	ikb := tgbotapi.NewInlineKeyboardButtonData("واریز", "variz")
	ikb2 = append(ikb2, ikb)
	ikb1 := tgbotapi.NewInlineKeyboardButtonData("برداشت", "bardasht")
	ikb2 = append(ikb2, ikb1)

	var ikb22 []tgbotapi.InlineKeyboardButton
	ikb_ := tgbotapi.NewInlineKeyboardButtonSwitch("switch", "233 usdt give")
	ikb22 = append(ikb22, ikb_)

	keyboard = append(keyboard, ikb2)
	keyboard = append(keyboard, ikb22)

	rr := &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
	return rr
}
