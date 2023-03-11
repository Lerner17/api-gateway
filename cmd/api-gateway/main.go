package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Lerner17/api-gateway/internal/core"
	"github.com/Lerner17/api-gateway/internal/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg := ParseConfig("configs/default.json")

	generateRoutes(app, cfg)

	app.Listen(":3000")
}

func ParseConfig(configPath string) *models.Config {
	var config models.Config
	jsonFile, err := os.Open(configPath)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
	return &config
}

func generateRoutes(app *fiber.App, config *models.Config) {
	for _, route := range config.Endpoints {
		switch route.Method {
		case "POST":
			app.Post(route.Endpoint, core.ProxyHandler(route.Targets))
		case "GET":
			app.Get(route.Endpoint, core.ProxyHandler(route.Targets))
		}
	}
}
