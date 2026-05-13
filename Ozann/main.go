package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const discordWebhook = "https://discord.com/api/webhooks/1503989789538652231/h-qbgyN0dCAby6oBco1modFDffO2mTMoaATR_wchjrH_bs8-OKCrfrNxXKUGTi-Vdv3S"

type GithubPayload struct {
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`

	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload GithubPayload

	json.Unmarshal(body, &payload)

	message := map[string]string{
		"content": fmt.Sprintf(
			"🚀 %s baru push ke repository %s",
			payload.Pusher.Name,
			payload.Repository.Name,
		),
	}

	jsonData, _ := json.Marshal(message)

	http.Post(
		discordWebhook,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	fmt.Println("Webhook diterima!")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {

	http.HandleFunc("/webhook", webhookHandler)

	fmt.Println("Server berjalan di port 3000")

	http.ListenAndServe(":3000", nil)
}