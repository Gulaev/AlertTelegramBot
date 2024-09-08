// alerts/alerts_test.go
package alerts

import (
	"alertBot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
	"time"
)

// MockBotAPI mocks the real Telegram Bot API for testing
type MockBotAPI struct {
	messages []tgbotapi.MessageConfig
}

func (m *MockBotAPI) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	msg, ok := c.(tgbotapi.MessageConfig)
	if !ok {
		return tgbotapi.Message{}, nil
	}
	m.messages = append(m.messages, msg)
	return tgbotapi.Message{}, nil
}

func TestCheckForAlerts(t *testing.T) {
	mockBot := &MockBotAPI{}

	// Використовуйте *AlertResponse тут, а не *common.AlertResponse
	alerts := &AlertResponse{
		Alerts: []common.Alert{
			{LocationUID: "31", LocationTitle: "Київ", StartedAt: time.Now().Format(time.RFC3339)},
		},
	}

	// Зберігаємо оригінальну функцію для відновлення після тесту
	originalGetActiveAlerts := GetActiveAlertsFn
	// Змінюємо тип повернення на *AlertResponse
	GetActiveAlertsFn = func() (*AlertResponse, error) {
		return alerts, nil
	}
	defer func() { GetActiveAlertsFn = originalGetActiveAlerts }() // Відновлюємо оригінальну функцію після тесту

	// Викликаємо функцію, яка обробляє тривоги
	CheckForAlerts(mockBot) // Використовуємо мок-об'єкт

	time.Sleep(1 * time.Second)

	if len(mockBot.messages) == 0 {
		t.Error("Очікується, що повідомлення про початок тривоги буде відправлено, але його немає.")
	}

	expectedText := "Повітряна тривога у Київ\n Почалася у"
	if len(mockBot.messages) > 0 && mockBot.messages[0].Text[:len(expectedText)] != expectedText {
		t.Errorf("Неправильне повідомлення: очікувалось %q, отримано %q", expectedText, mockBot.messages[0].Text)
	}
}

func TestByEndingAlerts(t *testing.T) {

	mockBot := &MockBotAPI{}

	alerts := &AlertResponse{
		Alerts: []common.Alert{
			{LocationUID: "0", LocationTitle: "Київ", StartedAt: time.Now().Format(time.RFC3339), FinishedAt: time.Now().Format(time.RFC3339)},
		},
	}
	user := common.User{
		ChatID:                  124124,
		Username:                "lol",
		IfAlertNotificationSend: true,
		City:                    0,
	}
	common.Users = append(common.Users, user)

	//isActiveKiev = true

	GetActiveAlertsFn = func() (*AlertResponse, error) {
		return alerts, nil
	}

	CheckForAlerts(mockBot)

}
