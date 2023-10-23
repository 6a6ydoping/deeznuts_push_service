package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/6a6ydoping/deeznuts_push_service/internal/entity"
	"net/http"
)

func (m *Manager) SendMessage(ctx context.Context, payload *entity.PushPayload) (message string, err error) {
	requestBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	req, err := http.NewRequest("POST", m.Config.Pusher.Url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.Config.Pusher.ApiToken)

	resp, err := m.Client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Notification sent successfully!")
		message = "Notification sent successfully!"
	} else {
		fmt.Println("Notification sending failed. Status code:", resp.Status)
		message = "Notification sending failed. Status code:" + resp.Status
	}
	return message, err
}
