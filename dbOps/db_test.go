package dbOps

import "testing"

func TestMain(m *testing.M) {
	ConfigDB()
	m.Run()
}

func TestAddUser(t *testing.T) {
	err := AddUser("xiaohuang", "123456")
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
	err := CreateRoom("自己的群", "yeyongping")
	if err != nil {
		t.Error(err)
	}
}

func TestVerificationOwner(t *testing.T) {
	if _, _, err := verificationOwner("c60fb081837d4699a91f0e8054f31207", "yeyongping");err != nil {
		t.Error(err)
	}
}

func TestVerificatMember(t *testing.T) {
	if _, _, err := verificationMember("c60fb081837d4699a91f0e8054f31207", "yeyongping");err != nil {
		t.Error(err)
	}
}

func TestVerificatAdministrator(t *testing.T) {
	if _,_, err := verificationAdministrator("c60fb081837d4699a91f0e8054f31207", "yeyongping"); err != nil {
		t.Error(err)
	}
}

func TestRoomAddMember(t *testing.T) {
	if err := RoomAddMember("c60fb081837d4699a91f0e8054f31207", "xiaohuang", "yeyongping"); err != nil {
		t.Error(err)
	}
}

func TestRoomTransferRoom(t *testing.T) {
	if err := RoomTransferRoom("c60fb081837d4699a91f0e8054f31207", "yeyongping", "xiaohuang"); err != nil {
		t.Error(err)
	}
}

func TestRoomRemoveMember(t *testing.T) {
	if err := RoomRemoveMember("c60fb081837d4699a91f0e8054f31207", "xiaohuang", "yeyongping" ); err != nil {
		t.Error(err)
	}
}

func TestRoomSetupAdministrator(t *testing.T) {
	if err := RoomSetupAdministrator("c60fb081837d4699a91f0e8054f31207", "xiaohuang", "yeyongping"); err != nil {
		t.Error()
	}
}

func TestRoomCancelAdministrator(t *testing.T) {
	if err := RoomCancelAdministrator("c60fb081837d4699a91f0e8054f31207", "xiaohuang", "yeyongping"); err != nil {
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

func TestAllRooms(t *testing.T) {
	if _,err := AllRooms("yeyongping"); err != nil {
		t.Error(err)
	}
}

func TestAllUserInRoom(t *testing.T) {
	if _, err := AllUserInRoom("c60fb081837d4699a91f0e8054f31207"); err != nil {
		t.Error(err)
	}
}