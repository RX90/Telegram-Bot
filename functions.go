package main

import (
	"log"
	"math/rand"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func connectWithTelegram() {
	token = os.Getenv("TELEGRAM_API_TOKEN")
	var err error
	if bot, err = tgbotapi.NewBotAPI(token); err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "markdown"
	bot.Send(msg)
}

func sendGopher(bot *tgbotapi.BotAPI, chatID int64) {
	filePath := mediaSlice[rand.Intn(len(mediaSlice))]
	text := mediaMap[filePath]
	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(filePath))
	msg.Caption = text
	bot.Send(msg)
}

func sendInlineKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = inlineRelaxButton
	bot.Send(msg)
}