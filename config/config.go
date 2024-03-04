package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AppConfig struct {
	PeeringURL string `yaml:"peeringURL"`
	Port       int    `yaml:"port"`
	GrpcAddr   string `yaml:"grpcAddr"`
	HttpAddr   string `yaml:"httpAddr"`
}

// Config ...
type Config struct {
	AppConfig `yaml:"application"`
}

// ProvideConfig provides the standard configuration to fx
func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
