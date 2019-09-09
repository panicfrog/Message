package internel

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiPort       int
	WebsocketPort int
	AesKey        string
}

var Configuration Config

func init() {
	Configuration = getConfig()
}

func getConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	apiPort, ok := viper.Get("apiPort").(int)
	if !ok {
		panic("get config of key 'apiPort' error")
	}

	websocketPort, ok := viper.Get("websocketPort").(int)
	if !ok {
		panic("get config of key 'websocketPort' error")
	}

	aesKey, ok := viper.Get("aesKey").(string)
	if !ok {
		panic("get config of key 'aesKey' error")
	}

	return Config{
		apiPort,
		websocketPort,
		aesKey,
	}
}
