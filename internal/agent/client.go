package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	ServerAddr string
}

func NewClient() *Client {
	return &Client{
		ServerAddr: "http://localhost:8080/api/v1/calculate",
	}
}

func (c *Client) SendExpression(expression string) (int, error) {
	payload, err := json.Marshal(map[string]string{"expression": expression})
	if err != nil {
		return 0, fmt.Errorf("ошибка при кодировании JSON: %v", err)
	}

	resp, err := http.Post(c.ServerAddr, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return 0, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var errResp map[string]string
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		if msg, exists := errResp["error"]; exists {
			return 0, fmt.Errorf("ошибка сервера: %s", msg)
		}
		return 0, fmt.Errorf("неожиданный код ответа: %d", resp.StatusCode)
	}

	var result struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("ошибка при декодировании ответа: %v", err)
	}

	return result.ID, nil
}
