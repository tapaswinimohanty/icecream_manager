package lib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

type config struct {
	AppName    string `yaml:"app_name"`
	AppVersion string `yaml:"app_version"`
	AppPort    string `yaml:"app_port"`
	Author     string `yaml:"author"`

	Development struct {
		DBHost     string `yaml:"db_host"`
		DBPort     string `yaml:"db_port"`
		DBName     string `yaml:"db_name"`
		DBUser     string `yaml:"db_user"`
		DBPassword string `yaml:"db_password"`
	} `yaml:"development"`

	Staging struct {
		DBHost     string `yaml:"db_host"`
		DBPort     string `yaml:"db_port"`
		DBName     string `yaml:"db_name"`
		DBUser     string `yaml:"db_user"`
		DBPassword string `yaml:"db_password"`
	} `yaml:"staging"`

	Production struct {
		DBHost     string `yaml:"db_host"`
		DBPort     string `yaml:"db_port"`
		DBName     string `yaml:"db_name"`
		DBUser     string `yaml:"db_user"`
		DBPassword string `yaml:"db_password"`
	} `yaml:"production"`

	Secret string `yaml:"secret"`

	Environment string `yaml:"environment"`
	SQLLogMode  bool   `yaml:"sql_log_mode"`
	Seed        bool   `yaml:"seed"`
}

var Config *config

func LoadConfig(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}

	Config.Production.DBHost = getEnv("DBHost", Config.Production.DBHost)
	Config.Production.DBName = getEnv("DBName", Config.Production.DBName)
	Config.Production.DBPort = getEnv("DBPort", Config.Production.DBPort)
	Config.Production.DBUser = getEnv("DBUser", Config.Production.DBUser)
	Config.Production.DBPassword = getEnv("DBPassword", Config.Production.DBPassword)
	Config.Secret = getEnv("Secret", Config.Secret)
	Config.Environment = getEnv("Environment", Config.Environment)

	seed, _ := strconv.ParseBool(getEnv("Seed", strconv.FormatBool(Config.Seed)))
	Config.Seed = seed

	sqlLogMode, _ := strconv.ParseBool(getEnv("SQLLogMode", strconv.FormatBool(Config.SQLLogMode)))
	Config.SQLLogMode = sqlLogMode
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
