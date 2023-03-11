package main

import (
	"testing"

	"github.com/qyqx233/chat-go-api/util"
)

func TestParse(t *testing.T) {
	t.Log(yamlConfig.Param.Openai)
}

func Test(t *testing.T) {
	// config := parseYaml("app.yaml")
	openai := util.ApiConfig{Ak: "23"}
	c := &util.Config{
		App: &util.AppConfig{
			Addr: ":3000",
		},
		ApiKey: map[string][]util.ApiConfig{},
	}
	c.ApiKey["openai"] = []util.ApiConfig{openai}
	// c.Param = make(map[string]map[string]any)
	// c.Param["openai"] = map[string]any{
	// 	"token": 4096,
	// }
	util.DumpYaml(c)
}
