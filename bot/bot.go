package bot

import (
	alertPkg "alertBot/alerts"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ListenForUpdates(bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil || update.CallbackQuery != nil {
			go handleUserInput(&update, *bot)
		}
	}
}

func GoCheckForAlerts(bot *tgbotapi.BotAPI) {
	alertPkg.CheckForAlerts(bot)
}
