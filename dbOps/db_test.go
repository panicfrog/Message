package dbOps

import "testing"

func TestMain(m *testing.M) {
	ConfigDB()
	m.Run()
}

func TestAddUser(t *testing.T) {
	err := AddUser("xiaohong", "123456")
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
	err := CreateRoom("小黄的群", "xiaohuang")
	if err != nil {
		t.Error(err)
	}
}

func TestVerificationOwner(t *testing.T) {
	if _, _, err := verificationOwner(1, "yeyongping");err != nil {
		t.Error(err)
	}
}

func TestVerificatMember(t *testing.T) {
	if _, _, err := verificationMember(2, "xiaohuang");err != nil {
		t.Error(err)
	}
}

func TestVerificatAdministrator(t *testing.T) {
	if _,_, err := verificationAdministrator(1, "yeyongping"); err != nil {
		t.Error(err)
	}
}

func TestRoomAddMember(t *testing.T) {
	if err := RoomAddMember(1, "xiaohong", "yeyongping"); err != nil {
		t.Error(err)
	}
}

func TestRoomTransferRoom(t *testing.T) {
	if err := RoomTransferRoom(2, "xiaohuang", "yeyongping"); err != nil {
		t.Error(err)
	}
}

func TestRoomRemoveMember(t *testing.T) {
	if err := RoomRemoveMember(1, "xiaohong", "yeyongping" ); err != nil {
		t.Error(err)
	}
}

func TestRoomSetupAdministrator(t *testing.T) {
	if err := RoomSetupAdministrator(2, "yeyongping", "xiaohuang"); err != nil {
		t.Error()
	}
}

func TestRoomCancelAdministrator(t *testing.T) {
	if err := RoomCancelAdministrator(2, "yeyongping", "xiaohuang"); err != nil {
		t.Error(err)
	}
}

func TestAddFriend(t *testing.T) {
	if err := AddFriend("yeyongping", "xiaohuang"); err != nil {
		t.Error(err)
	}
}

func TestRemoveFriend(t *testing.T) {
	if err := RemoveFriend("yeyongping", "xiaohuang"); err != nil {
		t.Error(err)
	}
}