package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	bot *tgbotapi.BotAPI
	token string
	inlineRelaxButton = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("+1 к спокойствию😌", "relax_button",
	)))
	mediaSlice = []string{"media/gofriends.png", "media/gobuilder.png", "media/gobuilder2.png", "media/goclassic.jpg", "media/gogalaxy.png", "media/gomusic.png", "media/goroutines.jpg", "media/goru.png", "media/gotravel.png", "media/gogc.png", "media/gostone.png"}
	mediaMap = map[string]string{
		"media/gofriends.png": "Гофер со своими лучшими друзьями🤝",
		"media/gobuilder.png": "Гофер что-то мастерит🔨",
		"media/gobuilder2.png": "Гофер на стройке👷",
		"media/goclassic.jpg": "Классический гофер",
		"media/gogalaxy.png": "Гофер в космосе🌍",
		"media/gomusic.png": "Гофер-музыкант🎹",
		"media/goroutines.jpg": "Гофер, показывающий силу горутин💪",
		"media/goru.png": "Русский гофер🎻",
		"media/gotravel.png": "Гофер-путешественник🗺️",
		"media/gogc.png": "Гофер-сборщик мусора",
		"media/gostone.png": "Гофер-каменщик",
	} 
)