package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gops/agent"
	"github.com/projectkita/project-harapan-backend-golang/src/common"
	"github.com/projectkita/project-harapan-backend-golang/src/config"
	"github.com/projectkita/project-harapan-backend-golang/src/route"
	"github.com/tokopedia/panics"

	logging "gopkg.in/tokopedia/logging.v1"
)

var appConfig config.AppConfig
var server config.Server

func init() {
	configPath := common.GenerateConfigPath(common.GetConfigDir(), "main", "yaml")
	err := appConfig.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	server = config.Server{BaseUrl: fmt.Sprintf("http://%s:%s", appConfig.Server.Host, appConfig.Server.Port), Prefix: "api"}
}

func openLogFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func setupLog() error {
	// Init Logger.
	logging.LogInit()

	logging.SetDebug(appConfig.Server.Debug)

	return nil
}

func setupPanicHandler() {
	hostname, _ := os.Hostname()
	environ := common.GetEnv()

	panics.SetOptions(&panics.Options{
		Env:             environ,
		SlackWebhookURL: appConfig.ChatNotification.MainURL,
		Filepath:        "/var/log/projectkita", // it'll generate panics.log
		Tags:            panics.Tags{"host": hostname},
	})
}

func main() {
	// gops helps us get stack trace if something wrong/slow in production
	if err := agent.Listen(nil); err != nil {
		log.Fatal(err)
	}

	setupPanicHandler()

	err := setupLog()
	if err != nil {
		log.Fatal(err)
	}

	route.InitRoutes(appConfig.Server.Port)
}
