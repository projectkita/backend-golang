package config

import (
	"io/ioutil"
	"time"

	"github.com/go-yaml/yaml"
)

type AppConfig struct {
	Server struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Env         string `yaml:"env"`
		Debug       bool   `yaml:"debug"`
		Maintenance bool   `yaml:"maintenance"`
	}
	Log struct {
		Error     string `yaml:"error"`
		Access    string `yaml:"access"`
		SentryDsn string `yaml:"sentry_dsn"`
	}
	AppDB       DatabaseServer `yaml:"database_server"`
	CacheServer struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"cache_server"`
	CachedData struct {
		Duration int `yaml:"duration"`
	} `yaml:"cached_data"`
	ChatNotification struct {
		URL      string              `yaml:"url"`
		MainURL  string              `yaml:"main_url"`
		IconURL  string              `yaml:"icon"`
		UserName string              `yaml:"username"`
		Channels map[string]string   `yaml:"channels"`
		Bots     map[string]SlackBot `yaml:"bots"`
	} `yaml:"chat_notification"`
	TimeZoneLocation struct {
		Location map[string]*time.Location
	}
}

type DatabaseServer struct {
	Master Database `yaml:"master"`
	Slave  Database `yaml:"slave"`
}

type SlackBot struct {
	IconURL  string `yaml:"icon_url"`
	UserName string `yaml:"username"`
}

type Database struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	DbName           string `yaml:"db_name"`
	ConnectionString string `yaml:"connection_string"`
}

type CronConfig struct {
	Cron struct {
		IsCron bool
		CronIP string
	}
}

type Server struct {
	BaseUrl string
	Prefix  string
}

func (s Server) GetBaseURL() string {
	return s.BaseUrl
}

func (s Server) GetPrefix() string {
	return s.Prefix
}

func (config *AppConfig) LoadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return err
	}

	return nil
}
