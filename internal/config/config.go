package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string 	`yaml:"address"`
}

type Config struct {
	Env string `yml:"env" yaml:"env" env-required:"true"`
	StoragePath string 	`yml:"storage_path" yaml:"storage_path"`
	HTTPServer `yaml:"http_server"`
}

func MustLoad(){
var configpath string	
configpath = os.Getenv("CONFIG_PATH")

if configpath == "" {
	flags:= flag.String("config", "", "path to the config file")
	flag.Parse()
	configpath = *flags

	if configpath == "" {
		log.Fatal("config path is required")
	}

}
}