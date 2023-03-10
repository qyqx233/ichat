package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	config := parseYaml("app.yaml")
	t.Log(config.Param.Openai)
}

func Test(t *testing.T) {
	// config := parseYaml("app.yaml")
	openai := ApiConfig{Ak: "23"}
	c := &Config{
		App: &AppConfig{
			Addr: ":3000",
		},
		ApiKey: map[string][]ApiConfig{},
	}
	c.ApiKey["openai"] = []ApiConfig{openai}
	// c.Param = make(map[string]map[string]any)
	// c.Param["openai"] = map[string]any{
	// 	"token": 4096,
	// }
	dumpYaml(c)
}
