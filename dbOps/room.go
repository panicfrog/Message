package dbOps

import (
	"github.com/jinzhu/gorm"
	"message/data"
	"message/internel"
)

// 创建房间
func CreateRoom(name string, creater string) error {
	c, err := FindUser(creater)
	if err != nil {
		return err
	}

	room := data.Room{
		RoomName: name,
		Owner: c.Account,
		Managers: []data.User{c},
		Users: []data.User{c},
	}
	if err := DB.Create(&room).Error; err != nil {
		return err
	}
	return  nil
}

func verificationOwner(room uint, owner string) (data.Room, error) {
	var r data.Room
	err := DB.Where(&data.Room{Owner: owner, Model: gorm.Model{ID: room} }).First(&r).Error
	if err != nil {
		return r, err
	}
	return r, nil
}

func verificationMember(room uint, member string) (data.Room, error) {
	var r data.Room
	err := DB.Where(&data.Room{Model:gorm.Model{ID: room}}).First(&r).Error
	if err != nil {
		return r, err
	}
	err = DB.Model(&r).Association("Users").Find(&r.Users).Error
	if err != nil {
		return r, err
	}
	if len(r.Users) > 0 {
		for _, u := range r.Users {
			if u.Account == member {
				goto EXITED
			}
		}
		return r, internel.RoomMemeberNotExited
	} else {
		return r, internel.RoomMemeberNotExited
	}
	EXITED:
	return r, nil
}

func vieificationAdministrator(room uint, administrator string) (data.Room, error) {
	var r data.Room
	err := DB.Where(&data.Room{Model:gorm.Model{ID: room}}).First(&r).Error
	if err != nil {
		return r, err
	}

	err = DB.Model(&r).Association("Managers").Find(&r.Managers).Error
	if err != nil {
		return r, err
	}
	if len(r.Managers) > 0 {
		for _, u := range r.Managers {
			if u.Account == administrator {
				goto EXITED
			}
		}
		return r, internel.RoomAdministratorNotExited
	} else {
		return r, internel.RoomAdministratorNotExited
	}
	EXITED:
	return r, nil
}

// 添加群成员
func RoomAddMember(room , member , creater string) {

}

// 删除群成员
func RoomRemoveMember(room , member , creater string) {

}

// 添加管理员
func RoomSetupAdministrator(room , administrator , creater string) {

}

// 移除管理员权限
func RoomCancelAdministrator(room , administrator , creater string) {

}

// 转让群
func RoomTransferRoom(room , transferee , creater string) {

}