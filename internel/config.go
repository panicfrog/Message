package internel

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiPort            int
	WebsocketPort      int
	AesKey             string
	DBHost             string
	DBPort             int
	DBUserName         string
	DBPasswd           string
	RedisAddr          string
	RedisPassword      string
	WebTokenExpire     int64
	MobileTokenExpire  int64
	DeskTopTokenExpire int64
}

var Configuration Config

func init() {
	Configuration = getConfig()
}

func getConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
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

	dbHost, ok := viper.Get("dbHost").(string)
	if !ok {
		panic("get config of key 'dbHost' error")
	}

	dbPort, ok := viper.Get("dbPort").(int)
	if !ok {
		panic("get config of key 'dbPort' error")
	}

	dbUserName, ok := viper.Get("dbUserName").(string)
	if !ok {
		panic("get config of key 'dbUserName' error")
	}

	dbPasswd, ok := viper.Get("dbPasswd").(string)
	if !ok {
		panic("get config of key 'dbPasswd' error")
	}

	redisAddr, ok := viper.Get("redisAddr").(string)
	if !ok {
		panic("get config of key 'reidsAddr' error")
	}

	redisPassword, ok := viper.Get("redisPassword").(string)
	if !ok {
		panic("get config of key 'redisPassword' error")
	}

	webTokenExpire, ok := viper.Get("webTokenExpire").(int)
	if !ok {
		panic("get config of key 'webTokenExpire' error")
	}

	mobileTokenExpire, ok := viper.Get("mobileTokenExpire").(int)
	if !ok {
		panic("get config of key 'mobileTokenExpire' error")
	}

	desktopTokenExpire, ok := viper.Get("desktopTokenExpire").(int)
	if !ok {
		panic("get config of key 'desktopTokenExpire' error")
	}

	return Config{
		apiPort,
		websocketPort,
		aesKey,
		dbHost,
		dbPort,
		dbUserName,
		dbPasswd,
		redisAddr,
		redisPassword,
		int64(webTokenExpire),
		int64(mobileTokenExpire),
		int64(desktopTokenExpire),
	}
}
