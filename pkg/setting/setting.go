package setting

import (
	"log"
	"os"
)

type App struct {
	RequestUrl string
}

var AppSetting = &App{}

func Setup() {
	loadAppEnv()
}

func loadAppEnv() {
	AppSetting.RequestUrl = os.Getenv("REQUEST_URL")

	requiredSettings := map[string]string{
		"REQUEST_URL": AppSetting.RequestUrl,
	}

	for settingKey, settingValue := range requiredSettings {
		if settingValue == "" {
			log.Fatalf("setting.Setup, required environment variable %s is missing!", settingKey)
		}
	}
}
