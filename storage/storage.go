package storage

import (
	"github.com/go-redis/redis"
	"message/data"
	"message/internel"
	"strconv"
	"time"
)

var redisClient *redis.Client

func SetupRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               internel.Configuration.RedisAddr,
		Password:           internel.Configuration.RedisPassword,
		DB:                 0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

}

func getField(token data.TokenPlayload) string {
	return internel.Md5(token.Account + strconv.Itoa(int(token.Platform)))
}

func SetToken(token data.TokenPlayload) error {
	f := getField(token)
	tokenStr, err := data.EncodeToken(&token)
	if err != nil {
		return err
	}
	if err = redisClient.Set(f, tokenStr, 0).Err(); err != nil {
		return err
	}
	sec := token.CreateTime/1e3
	nsec := token.CreateTime*1e6%1e9
	t := time.Unix(sec, nsec)
	var d time.Duration = 0
	if token.Platform == data.PlatformUnknow || token.Platform == data.PlatformWeb {
		d = time.Duration(internel.Configuration.WebTokenExpire * 60 * 1e9)
	} else if token.Platform == data.PlatformAndroid || token.Platform == data.PlatformiOS {
		d = time.Duration(internel.Configuration.MobileTokenExpire * 60 * 1e9)
	} else if token.Platform == data.PlatfromDesktop {
		d = time.Duration(internel.Configuration.DeskTopTokenExpire * 60 * 1e9)
	}
	t = t.Add(d)
	if err := redisClient.PExpireAt(f, t).Err(); err != nil {
		return err
	}
	return nil
}

func VerificationToken(token string) error {
	t, err := data.DecodeToken(token)
	if err != nil {
		return err
	}
	f := getField(t)
	d, err := redisClient.PTTL(f).Result()
	if err != nil {
		return internel.RedisTokenNotExited
	}

	if d == -2 {
		return internel.RedisTokenExpire
	}

	v, err := redisClient.Get(f).Result()
	if err != nil {
		return internel.RedisTokenNotExited
	}
	rt, err := data.DecodeToken(v)
	if err != nil {
		return err
	}
	if !rt.Equal(&t) {
		return internel.RedisTokenExpire
	}
	return nil
}

func VerficationToken(token string) bool {
	err := VerificationToken(token)
	return err == nil
}