package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	var err error
	bot, err := tgbotapi.NewBotAPI("1019630704:AAHmL-rjSKKw-xwjKKqRVfOjBGLoOzWmjQU")
	if err != nil {
		log.Panic("bot init error:", err)
		return
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.Limit = 10
	u.Offset = 0
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic("update error:", err)
	}
	var Menu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Base Convertion"),
			tgbotapi.NewKeyboardButton("Base Calkulator"),
			tgbotapi.NewKeyboardButton("Usuall Calkulator"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Help"),
			tgbotapi.NewKeyboardButton("Web Version"),
			tgbotapi.NewKeyboardButton("Button"),
		),
	)
	for update := range updates {

		cmd := update.Message.Command()

		if cmd == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There menu:")
			msg.ReplyMarkup = Menu
			bot.Send(msg)
		}

		if update.Message.Text == Menu.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Okay, lets start calculating, please write first number")
			bot.Send(msg)
		}

		if update.Message.Text == Menu.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "there going to be the chanse for you to convers youre base")
			bot.Send(msg)
		}
		if update.Message.Text == Menu.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There we going to use our algorithm")
			bot.Send(msg)
		}
		if update.Message.Text == Menu.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "if you closed you panel write /start")
			bot.Send(msg)
		}
		if update.Message.Text == Menu.Keyboard[1][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Web version in proces")
			bot.Send(msg)
		}
		if update.Message.Text == Menu.Keyboard[1][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just button")
			bot.Send(msg)
		}

	}

}
