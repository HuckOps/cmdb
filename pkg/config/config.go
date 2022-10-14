package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Server
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Mysql
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// Mongos
type Mongos struct {
	Server   []string `yaml:"server"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

// Redis
type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type ConfigStruct struct {
	Redis  Redis  `yaml:"redis"`
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Mongos Mongos `yaml:"mongos"`
}

var Config ConfigStruct

func ParserConfig(configPath string)  {
	if configFile, err := os.Open(configPath); err != nil {
		panic("Open config file failed")
	}else {
		 if err := yaml.NewDecoder(configFile).Decode(&Config); err != nil {
		 	panic("Parser config file failed")
		 }
	}
}



