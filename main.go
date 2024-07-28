package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var maxCount = 0
var HeaderAuth = ""

func sendRequest() {
	url := "https://api.hamsterkombatgame.io/clicker/tap"
	authHeader := HeaderAuth

	timestamp := time.Now().Unix()

	body := map[string]interface{}{
		"count":         maxCount,
		"availableTaps": 0,
		"timestamp":     timestamp,
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Ошибка при сериализации JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Статус ответа:", resp.Status)
}

func main() {
	fmt.Println("Program start")
	fmt.Println("Enter max count of your click")
	fmt.Scan(&maxCount)
	fmt.Println("Enter you header auth like this Bearer 17221144434937L8RsSgFbNtwxSRMLtsKJZffEsXn5vk4FcD9XJ8FZNsa6zpZCrTzxr0YkPHHydEt1097316536")
	fmt.Scan(&HeaderAuth)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	sendRequest()
	for {
		select {
		case <-ticker.C:
			fmt.Println("Отправка кликов на сервер")
			sendRequest()
		}
	}
}
