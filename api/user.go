package api

import (
	"message/dbOps"
	"message/internel"
)

func Register(account string, passwd string)  {
	err := dbOps.AddUser(account, passwd)
	if err == internel.DBErrorExited {
		// TODO: 发送错误的返回
	} else {
		// TODO: 发送请求成功
	}
}
