package main

import (
	"log"
	"math/rand"
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

	const fieldSize int = 3

	updates := bot.GetUpdatesChan(u)
	buttons := [fieldSize][fieldSize]string{{"11", "12", "13"}, {"21", "22", "23"}, {"31", "32", "33"}}

	const cross = "❌"	
	const null = "⭕️"

	for update := range updates {

		txt := update.Message.Text
		for i := 0; i < len(buttons); i++ {
			for j := 0; j < len(buttons[i]); j++ {
				if buttons[i][j] == txt {
					buttons[i][j] = cross
					for i := 0; i < fieldSize * fieldSize; i++ {
						x := rand.Intn(fieldSize)
						y := rand.Intn(fieldSize)
						if(buttons[x][y] != cross && buttons[x][y] != null) {
							buttons[x][y] = null
							break
						}
					}
				}
			}
		}

		var gameField = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(buttons[0][0]),
				tgbotapi.NewKeyboardButton(buttons[0][1]),
				tgbotapi.NewKeyboardButton(buttons[0][2]),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(buttons[1][0]),
				tgbotapi.NewKeyboardButton(buttons[1][1]),
				tgbotapi.NewKeyboardButton(buttons[1][2]),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(buttons[2][0]),
				tgbotapi.NewKeyboardButton(buttons[2][1]),
				tgbotapi.NewKeyboardButton(buttons[2][2]),
			),
		)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyMarkup = gameField
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		switch update.Message.Text {
		case "заново":
			msg.ReplyMarkup = gameField
		case "хватит":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if update.Message == nil { // ignore non-Message updates
			continue
		}

		playerWin := false

		// rows checking
		for i := 0; i < len(buttons); i++ {
			goodRow := true
			for j := 0; j < len(buttons[i]) - 1; j++ {
				if ((buttons[i][j] != buttons[i][j+1]) || (buttons[i][j] != cross)) {
					goodRow = false				
				}
			}
			if(goodRow) {
				playerWin = true
				break
			}
		}

		// column checking
		for i := 0; i < len(buttons) - 1; i++ {
			goodColumn := true
			for j := 0; j < len(buttons[i]); j++ {
				if ((buttons[i][j] != buttons[i+1][j]) || (buttons[i][j] != cross)) {
					goodColumn = false				
				}
			}
			if(goodColumn) {
				playerWin = true
				break
			}
		}

		// main diag checking
		for i := 0; i < len(buttons) - 1; i++ {
			good := true
			for j := 0; j < len(buttons[i]) - 1; j++ {
				if ((buttons[i][j] != buttons[i+1][j+1]) || (buttons[i][j] != cross)) {
					good = false				
				}
			}
			if(good) {
				playerWin = true
				break
			}
		}

		// side diag checking
		for i := 0; i < len(buttons) - 1; i++ {
			good := true
			for j := len(buttons[i]) - 1; j >= 1; j-- {
				if ((buttons[i][j] != buttons[i+1][j-1]) || (buttons[i][j] != cross)) {
					good = false				
				}
			}
			if(good) {
				playerWin = true
				break
			}
		}	

		if(playerWin) {
			
			winMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Тебе повезло! Ты выиграл(а)!")
			// echo message sending
			if _, err := bot.Send(winMsg); err != nil {
				log.Panic(err)
			}
			buttons = [fieldSize][fieldSize]string{{"11", "12", "13"}, {"21", "22", "23"}, {"31", "32", "33"}}
			
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
