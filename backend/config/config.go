package config

import (
	"fmt"
	"os"
	"sync"

	yaml "github.com/goccy/go-yaml"
)

type Config struct {
	Auth0 Auth0
}

type Auth0 struct {
	Domain   string
	Audience string
}

var config Config
var once = new(sync.Once)

func init() {
	once.Do(func() {
		path := fmt.Sprintf("data/config/%s.yaml", GetAppEnv())
		content, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		expandContent := []byte(os.ExpandEnv(string(content)))
		if err := yaml.Unmarshal(expandContent, &config); err != nil {
			panic(err)
		}
	})
}

func GetConfig() Config {
	return config
}

func GetAppEnv() string {
	appenv := os.Getenv("APP_ENV")
	if appenv == "" {
		return "development"
	}

	return appenv
}
