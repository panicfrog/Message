package storage

import (
	"message/data"
	"message/internel"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	SetupRedis()
	m.Run()
}

func TestSetTokenAndVerificationToken(t *testing.T) {
	token := data.NewTokenPlayload("yeyongping", data.PlatformiOS)
	st, err := data.EncodeToken(&token)
	if err != nil {
		t.Error(err)
	}
	err = SetToken(token)
	if err != nil {
		t.Error(err)
	}

	err = VerificationToken(st)
	if err != nil {
		t.Error(err)
		return
	}

	timeout := time.After(time.Duration(internel.Configuration.MobileTokenExpire * 60 * 1e9 + 1))
	<- timeout
	err = VerificationToken(st)
	if err != internel.RedisTokenNotExited && err != internel.RedisTokenExpire {
		t.Error(err)
	}

}
