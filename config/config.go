package config

import (
	"log"

	"github.com/BurntSushi/toml"
	rest "github.com/node-a-team/irisnet-validator_exporter/getData/rest"
	rpc "github.com/node-a-team/irisnet-validator_exporter/getData/rpc"
)

const ()

var (
	ConfigPath string
	Config     configType
)

type configType struct {
	Title string `json:"title"`

	Servers struct {
		Addr struct {
			RPC  string `json:"rpc"`
			REST string `json:"rest"`
		}
	}

	Validator struct {
		OperatorAddr string `json:"operatorAddr"`
	}

	Options struct {
		ListenPort string `json:"listenPort"`
	}
}

func Init() {

	Config = readConfig()

	rpc.Addr = Config.Servers.Addr.RPC
	rest.Addr = Config.Servers.Addr.REST

	rest.OperAddr = Config.Validator.OperatorAddr

}

func readConfig() configType {

	var config configType

	if _, err := toml.DecodeFile(ConfigPath+"/config.toml", &config); err != nil {

		log.Fatal("Config file is missing: ", config)
	}

	return config

}
