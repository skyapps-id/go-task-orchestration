package tasks

import (
	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Payload struct {
	Url  string `json:"url"`
	Body string `json:"body"`
}

func DecodeToTask(msg string, task interface{}) (err error) {
	decodedstg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return
	}
	msgBytes := []byte(decodedstg)
	err = json.Unmarshal(msgBytes, task)
	if err != nil {
		return
	}
	return
}

func SendWebhook(b64payload string) (bool, error) {
	payload := Payload{}
	DecodeToTask(b64payload, &payload)

	log.Info("=========> Task Run", payload)

	return false, nil
}
