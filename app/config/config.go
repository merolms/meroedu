package config

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Host                 string
		Port                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
	}
	Server struct {
		Address string
	}
}

// C ...
var C config

// ReadConfig ...
func ReadConfig() {
	Config := &C
	viper.SetConfigFile(`config.yml`)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error Reading Config %s\n", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Error unmarshal config '%s'\n", err)
	}

	spew.Dump(C)
}
