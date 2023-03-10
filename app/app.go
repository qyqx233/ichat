package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	api "github.com/qyqx233/chat-go-api/api"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "c", "app.yml", "config path")
	flag.Parse()
	config := parseYaml(configPath)
	if config.App.MaxCpu == 0 {
		config.App.MaxCpu = 4
	}
	newDB(config.Db)
	// runtime.GOMAXPROCS(config.App.MaxCpu)
	cfgs := make([]api.ClientCfg, 8)
	for k, vs := range config.ApiKey {
		for _, v := range vs {
			cfgs = append(cfgs, api.ClientCfg{
				Name: k,
				Ak:   v.Ak,
			})
		}
	}
	api.NewAll(cfgs)
	app := fiber.New()
	app.Post("/api/chat/qa", handleChat)
	app.Post("/api/chat/session", newSession)
	app.Listen(config.App.Addr)
}
