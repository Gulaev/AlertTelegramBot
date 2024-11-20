package alerts

import (
	"alertBot/common"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AlertResponse struct {
	Alerts []common.Alert `json:"alerts"`
}

var apiToken = os.Getenv("ALERT_TOKEN")

// GetActiveAlerts fetches active alerts from the API
func GetActiveAlerts() (*AlertResponse, error) {
	url := "https://api.alerts.in.ua/v1/alerts/active.json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	var alerts AlertResponse
	if err := json.NewDecoder(resp.Body).Decode(&alerts); err != nil {
		return nil, err
	}

	return &alerts, nil
}
