package config

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/jalopezma/shops-api/common"
	"github.com/jalopezma/shops-api/providers"
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
		Hosts       []string    `yaml:"hosts"`
		User        string      `yaml:"user"`
		Password    string      `yaml:"password"`
		DbName      string      `yaml:"db_name"`
		ReplicaSet  string      `yaml:"replica_set"`
		Ssl         bool        `yaml:"ssl"`
		AuthSource  string      `yaml:"auth_source"`
		Collections Collections `yaml:"collections"`
	} `yaml:"mongo"`

	Redis struct {
		SentinelHosts []string `yaml:"sentinel_hosts"`
		SentinelGroup string   `yaml:"sentinel_group"`
		Password      string   `yaml:"password"`
	} `yaml:"redis"`
}

// Collections yaml struct
type Collections struct {
	Users string `yaml:"users"`
	Shops string `yaml:"shops"`
}

// GetConfig parses the file to config struct
func GetConfig(env, configPath string) Config {

	env = strings.ToLower(env)
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
	common.ConnectMongo(config.Mongo.Hosts, config.Mongo.DbName, config.Mongo.User, config.Mongo.Password, config.Mongo.ReplicaSet, config.Mongo.Ssl, config.Mongo.AuthSource)

	shopsSession := common.GetMongoSession()
	shops := providers.ShopsProvider{
		Session:    shopsSession,
		Collection: shopsSession.DB(config.Mongo.DbName).C(config.Mongo.Collections.Shops),
	}

	common.Deps = common.Dependencies{
		Shops: &shops,
	}
}
