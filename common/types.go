package common

import (
	"sync"
)

// Alert represents the structure of an alert
type Alert struct {
	ID             int    `json:"id"`
	LocationTitle  string `json:"location_title"`
	LocationType   string `json:"location_type"`
	StartedAt      string `json:"started_at"`
	FinishedAt     string `json:"finished_at"`
	UpdatedAt      string `json:"updated_at"`
	AlertType      string `json:"alert_type"`
	LocationUID    string `json:"location_uid"`
	LocationOblast string `json:"location_oblast"`
}

type AlertResponse struct {
	Alerts []Alert `json:"alerts"`
}

var (
	mu    = sync.Mutex{}
	Users []User
)

type User struct {
	Username                string
	ChatID                  int64
	IfAlertNotificationSend bool
	City                    int32
	Step                    string
	State                   UserState
	PreviousMessageID       int
}

func FindUserByChatID(chatID int64) *User {
	for _, user := range Users {
		if user.ChatID == chatID {
			return &user
		}
	}
	return nil
}

func IsChatIdContainsInUser(chatID int64) bool {
	for _, user := range Users {
		if user.ChatID == chatID {
			return true
		}
	}
	return false
}

func UpdateUserIsAlertNotificationSend(chatID int64, isSend bool) {
	for i, user := range Users {
		if user.ChatID == chatID {
			mu.Lock()
			Users[i].IfAlertNotificationSend = isSend
			mu.Unlock()
			break
		}
	}
}

func UpdateExistUser(user User) {
	for i, userRange := range Users {
		if userRange.ChatID == user.ChatID {
			mu.Lock()
			Users[i] = user
			mu.Unlock()
		}
	}
}
