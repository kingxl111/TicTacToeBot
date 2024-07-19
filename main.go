package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// func updateCallbackQueryHandler(u update.CallbackQuery) {

// }

func main() {
	bot, err := tgbotapi.NewBotAPI("7199838256:AAHWeohLnNmjFPT2B73Q8fDXCuwFBxKrSdw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	buttons := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for update := range updates {

		txt := update.Message.Text
		for i := 0; i < len(buttons); i++ {
			if buttons[i] == txt {
				buttons[i] = " "
			}
		}

		var gameField = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(buttons[0]),
				tgbotapi.NewKeyboardButton(buttons[1]),
				tgbotapi.NewKeyboardButton(buttons[2]),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(buttons[3]),
				tgbotapi.NewKeyboardButton(buttons[4]),
				tgbotapi.NewKeyboardButton(buttons[5]),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(buttons[6]),
				tgbotapi.NewKeyboardButton(buttons[7]),
				tgbotapi.NewKeyboardButton(buttons[8]),
			),
		)

		// switch update.Message.Text {
		// case "play":
		// 	msg.ReplyMarkup = gameField
		// case "stop":
		// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		// }

		if update.Message == nil { // ignore non-Message updates
			continue
		}

		// echo message generating
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyMarkup = gameField

		// echo message sending
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		// if update.CallbackQuery != nil {
		// 	updateCallbackQueryHandler(update.CallbackQuery)
		// 	continue
		// }
		// if update.Message == nil {
		// 	continue
		// }
	}
}
