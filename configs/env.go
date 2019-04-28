package configs

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func readConfig(filename string) error {
	vi := viper.New()

	vi.SetConfigName(filename)
	vi.AddConfigPath(".")
	vi.AutomaticEnv()

	err := vi.ReadInConfig()
	if err != nil {
		return err
	}

	// environment := vi.GetString("environment")

	// get env from command flag
	envString := flag.String("ENV", "development", "ENVIRONMENT VARIABLE")
	flag.Parse()

	environment := strings.ToLower(*envString)

	if environment != "staging" && environment != "production" {
		environment = "development"
	}

	// Get DB Config by Environment
	dbConfig := vi.GetStringMapString(environment)
	for k, _ := range dbConfig {
		key := strings.ToUpper(k)

		err := os.Setenv(key, dbConfig[k])
		if err != nil {
			return err
		}
	}

	var config map[string]interface{}
	err = vi.Unmarshal(&config)
	if err != nil {
		return err
	}

	for k, v := range config {
		if reflect.TypeOf(v).String() == "string" || reflect.TypeOf(v).String() == "int" {
			err = os.Setenv(strings.ToUpper(k), vi.GetString(k))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SetEnv() {
	err := readConfig("app_config")

	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}
}
