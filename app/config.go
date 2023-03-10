package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppStaticConfig struct {
	Dir    string `json:"dir,omitempty"`
	Url    string `json:"url,omitempty"`
	Browse bool   `json:"browse,omitempty"`
}

type ApiConfig struct {
	Ak string `json:"ak,omitempty"`
}

type DbConfig struct {
	Driver string `yaml:"driver" json:"driver,omitempty"`
	Source string `yaml:"source,omitempty"`
}

type AppConfig struct {
	Addr   string           `yaml:"addr" json:"addr,omitempty"`
	MaxCpu int              `yaml:"max_cpu" json:"max_cpu,omitempty"`
	Static *AppStaticConfig `json:"static,omitempty"`
}

type ParamOpenai struct {
	Token int
}

type ParamConfig struct {
	Openai ParamOpenai `yaml:"openai,omitempty"`
}

type Config struct {
	App    *AppConfig             `yaml:"app"`
	ApiKey map[string][]ApiConfig `yaml:"api_key,omitempty"`
	Db     *DbConfig
	Param  ParamConfig
}

func parseYaml(path string) *Config {
	var c = new(Config)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	return c
}

func dumpYaml(c *Config) {
	data, _ := yaml.Marshal(c)
	os.WriteFile("app.yaml", data, 0644)
}
