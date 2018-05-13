package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var config Config

// MongoAccess yaml struct
type MongoAccess struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	AuthDB   string `yaml:"authDb"`
}

// Config settings file
type Config struct {
	Mongo struct {
		Hosts      []string `yaml:"hosts"`
		DbName     string   `yaml:"db_name"`
		ReplicaSet string   `yaml:"replica_set,omitempty"`
	} `yaml:"mongo"`

	Redis struct {
		SentinelHosts []string `yaml:"sentinel_hosts"`
		SentinelGroup string   `yaml:"sentinel_group"`
		Password      string   `yaml:"password"`
	} `yaml:"redis"`
}

// GetConfig parses the file to config struct
func GetConfig(env, configPath string) Config {

	configFile := configPath + "settings_" + env + ".yml"

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("[Config] yamlFile.Get err #%v", err)
	}

	config = Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("[Config] Unmarshal: %v", err)
	}
	return config
}

// InitDeps - Reads settings file and sets dependencies
func InitDeps(env string) {
	GetConfig(env, "./")
	log.Printf("\n%#v\n", config)
}
