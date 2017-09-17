package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func GetEnv() string {
	environ := os.Getenv("ENV")
	if environ == "" {
		environ = "development"
	}

	return environ
}

func GenerateConfigPath(path string, module string, extension string) string {
	environ := GetEnv()

	fname := path + "/" + module + "." + environ + "." + extension

	return fname
}

func GetConfigDir() string {
	prefix := ""

	env := os.Getenv("APPENV")
	if env == "" || env == "development" {
		prefix = "./files"
	}

	dir := prefix + "/etc/projectkita"
	return fmt.Sprintf("%s/config", dir)
}

// Log to Slack Channel
func LogToSlack(payload interface{}) {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	// create request
	req, _ := http.NewRequest("POST", app.Config.ChatNotification.URL, bytes.NewBuffer(jsonPayload))

	// send http post to slack
	client := &http.Client{Timeout: time.Second * 5}

	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
}
