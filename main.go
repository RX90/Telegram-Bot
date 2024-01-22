package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	connectWithTelegram()

	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updConfig)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			log.Printf(`Message: "%s" User: "%s" Chat: "%s"`, update.Message.Text, update.Message.From.UserName, update.Message.Chat.Title)
			switch update.Message.Command() {
				case "start":
					text := fmt.Sprintf("Привет, *%s*🤗😲\nМожешь посмотреть мои команды:\n/start - Запуск бота👾\n/relax - Кнопка спокойствия✨\n/gopher - Случайный гофер🐿️", update.Message.From.FirstName)
					sendMessage(bot, update.Message.Chat.ID, text)
				case "relax":
					sendInlineKeyboard(bot, update.Message.Chat.ID, "Нажми на кнопку спокойствия✨")
				case "gopher":
					sendGopher(bot, update.Message.Chat.ID)
			}
		} else if update.Message != nil && update.Message.Chat.Type != "private" {
			switch {
				case update.Message.NewChatTitle != "":
					log.Printf(`User: "%s" New Title: "%s"`, update.Message.From.UserName, update.Message.NewChatTitle)
					text := "Хорошее название"
					sendMessage(bot, update.Message.Chat.ID, text)
				case update.Message.NewChatMembers != nil:
					log.Printf(`New Member: "%s"`, update.Message.NewChatMembers[len(update.Message.NewChatMembers)-1].FirstName)
					text := fmt.Sprintf("Привет, *%s*!", update.Message.NewChatMembers[len(update.Message.NewChatMembers)-1].FirstName)
					sendMessage(bot, update.Message.Chat.ID, text)
				case update.Message.NewChatPhoto != nil:
					log.Printf(`User: "%s" Chat: "%s" New Chat Photo`, update.Message.From.UserName, update.Message.Chat.Title)
					sendMessage(bot, update.Message.Chat.ID, "Отличное фото")
			}
		} else if update.CallbackQuery != nil {
			log.Printf(`User: "%s" Chat: "%s" Data: "%s"`, update.CallbackQuery.From.UserName, update.CallbackQuery.Message.Chat.Title, update.CallbackQuery.Data)
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Может нажмёте ещё раз?\nЭта кнопка успокаивает✨")
			msg.ReplyMarkup = inlineRelaxButton
			bot.Send(msg)
		}
	}
}