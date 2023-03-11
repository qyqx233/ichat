package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	api "github.com/qyqx233/chat-go-api/api"
	"github.com/qyqx233/chat-go-api/util"
)

func newApp() *fiber.App {
	app := fiber.New()
	app.Post("/api/chat/qa", handleChat)
	app.Post("/api/chat/session", newSession)
	return app
}

func init() {
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func initApi(config *util.Config) {
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
}

var yamlConfig *util.Config

func main() {
	var configPath string
	flag.StringVar(&configPath, "c", "app.yaml", "config path")
	flag.Parse()
	util.ParseYaml(configPath)
	if yamlConfig.App.MaxCpu == 0 {
		yamlConfig.App.MaxCpu = 4
	}
	newDB(yamlConfig.Db)
	initApi(yamlConfig)
	// runtime.GOMAXPROCS(config.App.MaxCpu)

	app := newApp()
	app.Listen(yamlConfig.App.Addr)
}
