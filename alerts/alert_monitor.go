package alerts

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"alertBot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

var (
	rwMu              = sync.RWMutex{}
	GetActiveAlertsFn = GetActiveAlerts
)

func CheckForAlerts(botAPI BotAPI) {
	for {

		alerts, err := GetActiveAlertsFn()
		if err != nil {
			log.Println("Error fetching alerts:", err)
			time.Sleep(30 * time.Second)
			continue
		}
		checkByEndingAlerts(alerts, &botAPI)
		for _, alert := range alerts.Alerts {
			sendMessageAboutAlert(&alert, &botAPI)
		}
		time.Sleep(10 * time.Second)
	}
}

func parseStringToInt(str string) int32 {
	num32, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("Error parsing string:", err)
		return 0
	}
	return int32(num32)
}

func sendMessageAboutAlert(alert *common.Alert, api *BotAPI) {
	rwMu.RLock()
	defer rwMu.RUnlock()

	for i, user := range common.Users {
		if user.City == parseStringToInt(alert.LocationUID) {
			if !user.IfAlertNotificationSend {
				t, err := time.Parse(time.RFC3339, alert.StartedAt)
				if err != nil {
					log.Println("Error parsing alert start time:", err)
					continue
				}
				t = t.Add(3 * time.Hour)
				alertStartTime := t.Format("15:04")
				message := fmt.Sprintf("⚠️Повітряна тривога у %s\nПочалася у %s\nПерейдіть в укриття",
					alert.LocationTitle, alertStartTime)

				go sendMessage(user.ChatID, message, *api)

				rwMu.RUnlock()
				rwMu.Lock()
				common.Users[i].IfAlertNotificationSend = true
				rwMu.Unlock()
				rwMu.RLock()
			}
		}
	}
}

func checkByEndingAlerts(alerts *AlertResponse, botAPI *BotAPI) {
	rwMu.RLock()
	defer rwMu.RUnlock()

	for i, user := range common.Users {
		isAlertActive := false
		for _, alert := range alerts.Alerts {
			city := user.City
			if parseStringToInt(alert.LocationUID) == city {
				isAlertActive = true
				break
			}
		}
		if !isAlertActive && user.IfAlertNotificationSend {
			message := fmt.Sprintf("Відбій повітряної тривоги у %s", common.GetOblastNameByNumber(int(user.City)))
			go sendMessage(user.ChatID, message, *botAPI)

			rwMu.RUnlock()
			rwMu.Lock()
			common.Users[i].IfAlertNotificationSend = false
			rwMu.Unlock()
			rwMu.RLock()
		}
	}
}

func sendMessage(chatID int64, message string, botAPI BotAPI) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := botAPI.Send(msg)
	if err != nil {
		log.Printf("Error sending message to user %d: %v", chatID, err)
		return
	}
}
