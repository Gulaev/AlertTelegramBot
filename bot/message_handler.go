package bot

import (
	"alertBot/common"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"sync"
)

var rwMu = sync.RWMutex{}

func handleUserInput(update *tgbotapi.Update, botApi tgbotapi.BotAPI) {

	if update.Message != nil && update.Message.Text == "/start" {
		user := common.User{
			ChatID:                  update.Message.Chat.ID,
			Username:                update.Message.From.UserName,
			IfAlertNotificationSend: false,
			State:                   common.StateChoosingCity,
		}
		if !common.IsChatIdContainsInUser(user.ChatID) {
			rwMu.Lock()
			common.Users = append(common.Users, user)
			rwMu.Unlock()

		} else {
			go common.UpdateExistUser(user)
		}

		message := "Привіт @" + user.Username + "\nВибери місто щоб отримувати сповіщення про тривогу"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyMarkup = buildChoiceCityInlineKeyboard()

		messageWhatSend, err := botApi.Send(msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
		fmt.Println(messageWhatSend.MessageID)
		user.PreviousMessageID = messageWhatSend.MessageID
		go common.UpdateExistUser(user)

	} else if update.CallbackQuery != nil &&
		common.FindUserByChatID(update.CallbackQuery.Message.Chat.ID).State == common.StateChoosingCity {
		selectedOblastNumber := update.CallbackQuery.Data
		fmt.Println(selectedOblastNumber)
		currentUser := common.FindUserByChatID(update.CallbackQuery.Message.Chat.ID)

		oblastNumber, err := strconv.Atoi(selectedOblastNumber)
		if err != nil {
			fmt.Println("Error converting callback data to int:", err)
			return
		}
		oblastName := common.GetOblastNameByNumber(oblastNumber)
		currentUser.City = int32(oblastNumber)
		currentUser.State = common.StateConfirmed

		confirmationMessage := fmt.Sprintf("Ви вибрали %s для отримання сповіщень про тривогу.", oblastName)

		editMsg := tgbotapi.NewEditMessageText(currentUser.ChatID, currentUser.PreviousMessageID,
			confirmationMessage)

		message, err := botApi.Send(editMsg)
		if err != nil {
			fmt.Println("Error editing message:", err)
		}
		currentUser.PreviousMessageID = message.MessageID
		go common.UpdateExistUser(*currentUser)

		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		if _, err := botApi.Request(callback); err != nil {
			fmt.Println("Error acknowledging callback:", err)
		}
		fmt.Println(common.Users)
	}
}

func buildChoiceCityInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for num, name := range common.OblastNames {
		button := tgbotapi.NewInlineKeyboardButtonData(name, fmt.Sprintf("%d", num))
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(button))
	}
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	return inlineKeyboard
}
