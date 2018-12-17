package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	//ConfTypeGateway enum type for gateway config
	ConfTypeGateway = 0x002
	//ConfTypeAPI enum type for api config
	ConfTypeAPI = 0x003
	//ConfTypeConsole enum type for console
	ConfTypeConsole = 0x004
)

//InitialConf initial configuration
func InitialConf(t int32) {
	switch t {
	case ConfTypeGateway:
		viper.SetConfigName("gateway")
	case ConfTypeAPI:
		viper.SetConfigName("api")
	case ConfTypeConsole:
		viper.SetConfigName("console")
	default:
		panic("Illegal Conf type")
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}

}

//GetListeningPort return service listen port from conf
func GetListeningPort() int {
	return viper.GetInt("runtime.port")
}
