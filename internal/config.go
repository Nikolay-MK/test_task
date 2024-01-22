package internal

import (
	"encoding/json"
	"io/ioutil"
)

// AppConfig структура для конфигураций
type AppConfig struct {
	URL   string `json:"url"`
	DBURL string `json:"db_url"`
}

// LoadConfig загружает конфигурацию из файла
func LoadConfig(configPath string) (AppConfig, error) {
	//embed.fs почитать!
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return AppConfig{}, err
	}

	var appConfig AppConfig
	err = json.Unmarshal(configData, &appConfig)
	if err != nil {
		return AppConfig{}, err
	}

	return appConfig, nil
}
