package util

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

type ParamOpenAI struct {
	Count int
	Token int
}

type ParamProxy struct {
	URL string `yaml:"url,omitempty"`
}

type ParamConfig struct {
	Openai ParamOpenAI `yaml:"openai,omitempty" json:"openai,omitempty"`
	Proxy  ParamProxy  `yaml:"proxy,omitempty"`
}

type Config struct {
	App    *AppConfig             `yaml:"app"`
	ApiKey map[string][]ApiConfig `yaml:"api_key,omitempty"`
	Db     *DbConfig
	Param  ParamConfig
}

func ParseYaml(path string) {
	var c = new(Config)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	YamlConfig = c
}

func DumpYaml(c *Config) {
	data, _ := yaml.Marshal(c)
	os.WriteFile("app.yaml", data, 0644)
}

var YamlConfig *Config
