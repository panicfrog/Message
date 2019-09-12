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

func TestCreateRoom(t *testing.T) {
	err := CreateRoom("第一个群", "yeyongping")
	if err != nil {
		t.Error(err)
	}
}

func TestVerificationOwner(t *testing.T) {
	if _, err := verificationOwner(1, "yeyongping");err != nil {
		t.Error(err)
	}
}

func TestVerificatMember(t *testing.T) {
	if _, err := verificationMember(1, "yeyongping");err != nil {
		t.Error(err)
	}
}

func TestVerificatAdministrator(t *testing.T) {
	if _, err := vieificationAdministrator(1, "yeyongping"); err != nil {
		t.Error(err)
	}
}