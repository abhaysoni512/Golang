package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)
type HTTPServer struct{
	Address string
}


type Config struct{
	Env string 	`yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`

}


// how to serialise pacakge

func Mustload() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH") 

	if configPath == ""{
		flags := flag.String("config","","path to config")
		flag.Parse()

		configPath = *flags

		if configPath == ""{
			log.Fatal("Config path isn't set")
		}
	}

	_,err := os.Stat(configPath)

	if err != nil{
		log.Fatalf("config file does not exist: %s",configPath)
	}

	var cfg Config

	err  = cleanenv.ReadConfig(configPath,&cfg)
	if err != nil{
		log.Fatalf("can not read confif file : %s", err.Error())
	}

	return &cfg
}