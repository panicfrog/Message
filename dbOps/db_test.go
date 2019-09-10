package dbOps

import "testing"

func TestMain(m *testing.M) {
	ConfigDB()
	m.Run()
}

func TestAddUser(t *testing.T) {
	err := AddUser("yeyongping", "123456")
	if err != nil {
		t.Error(err)
	}
}

func TestFindUser(t *testing.T) {
	account := "yeyongping"
	user, err := FindUser(account)
	if err != nil {
		t.Error(err)
	}
	if user.Account != account {
		t.Error("查询错误")
	}
}

func TestChangePasswd(t *testing.T) {
	err := ChangePasswd("yeyongping", "123456", "123456abc")
	if err != nil {
		t.Error(err)
	}
}