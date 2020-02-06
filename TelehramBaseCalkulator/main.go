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
	updChan, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic("update error:", err)
	}

	for {
		update := <-updChan
		if update.Message.IsCommand() == true && update.Message.Command() != "Calculate" {
			cmd := update.Message.Command()
			if cmd == "help" {
				msgAnswer := tgbotapi.NewMessage(update.Message.Chat.ID, "there is nothing yet")
				bot.Send(msgAnswer)
			}
			if cmd == "start" {
				msgAnswer := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello to you")
				bot.Send(msgAnswer)
			}
		}
		//Todo: there i am trying to get separeted user request's
		if update.Message.Command() == "Calculate" {
			msgAnswer := tgbotapi.NewMessage(update.Message.Chat.ID, "Okay, lets start calculating, please write first number")
			bot.Send(msgAnswer)
			for {
				if update.Message.Text == "" {
					msgAnswer := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					bot.Send(msgAnswer)

				}
				if update.Message.Text == "Hey" {
					msgAnswer := tgbotapi.NewMessage(update.Message.Chat.ID, "Hey to you")
					bot.Send(msgAnswer)
					break
				}
			}

		} else {
			if update.Message.Text != "" {
				msgAnswer := tgbotapi.NewMessage(update.Message.Chat.ID, "hmmm")
				bot.Send(msgAnswer)
			}
		}
	}

}
