package dbOps

import (
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
		RoomDisplayID: internel.Uuid(),
		Owner: c.Account,
		Managers: []data.User{c},
		Users: []data.User{c},
	}
	if err := DB.Create(&room).Error; err != nil {
		return err
	}
	return  nil
}

//

func verificationOwner(room string, owner string) (data.User, data.Room, error) {
	var r data.Room
	var u data.User
	err := DB.Where(&data.Room{Owner: owner, RoomDisplayID: room }).First(&r).Error
	if err != nil  {
		return u, r, err
	}
	if  err = DB.Where(&data.User{Account: owner}).First(&u).Error; err != nil {
		return  u, r, err
	}
	return u, r, nil
}

func verificationMember(room string, member string) (data.User, data.Room, error) {
	var r data.Room
	var u data.User
	err := DB.Where(&data.Room{RoomDisplayID: room}).First(&r).Error
	if err != nil {
		return u, r, err
	}
	err = DB.Model(&r).Association("Users").Find(&r.Users).Error
	if err != nil {
		return u, r, err
	}
	if len(r.Users) > 0 {
		for _, _u := range r.Users {
			if _u.Account == member {
				u = _u
				goto EXITED
			}
		}
		return u, r, internel.RoomMemeberNotExited
	} else {
		return u, r, internel.RoomMemeberNotExited
	}
	EXITED:
	return u, r, nil
}

func verificationAdministrator(room string, administrator string) (data.User, data.Room, error) {
	var r data.Room
	var a data.User
	err := DB.Where(&data.Room{RoomDisplayID:room}).First(&r).Error
	if err != nil {
		return a, r, err
	}

	err = DB.Model(&r).Association("Managers").Find(&r.Managers).Error
	if err != nil {
		return a, r, err
	}
	if len(r.Managers) > 0 {
		for _, u := range r.Managers {
			if u.Account == administrator {
				a = u
				goto EXITED
			}
		}
		return a, r, internel.RoomAdministratorNotExited
	} else {
		return a, r, internel.RoomAdministratorNotExited
	}
	EXITED:
	return a, r, nil
}

// 添加群成员
func RoomAddMember(room string, member , administrator string) error {
	_, r, err := verificationAdministrator(room, administrator)
	if err != nil {
		return err
	}

	m, _, err := verificationMember(room, member)
	if err != nil && err != internel.RoomMemeberNotExited {
		return err
	}

	if err == internel.RoomMemeberNotExited {
		if err = DB.Where(&data.User{Account: member}).First(&m).Error; err != nil {
			return  err
		}
	}

	err = DB.Model(&r).Association("Users").Append(m).Error
	if err != nil {
		return err
	}
	return  nil
}

// 删除群成员
func RoomRemoveMember(room string, member , administrator string) error {
	_,_, err := verificationAdministrator(room, administrator)
	if err != nil {
		return err
	}

	u, r, err := verificationMember(room, member)
	if err != nil {
		return err
	}

	_,_, err = verificationAdministrator(room, member)
	if err == nil {
		return internel.RoomAdministratorUnremoveable
	}

	err = DB.Model(&r).Association("Users").Delete(u).Error
	if err != nil {
		return err
	}
	return nil
}

// 设置管理员
func RoomSetupAdministrator(room string, administrator , owner string) error {
	if  _, _, err := verificationOwner(room, owner); err != nil {
		return err
	}
	_, _, err := verificationAdministrator(room, administrator)
	if err == nil {
		return internel.RoomIsAlreadyAdministrator
	}

	u, r, err := verificationMember(room, administrator)
	if err != nil {
		return err
	}
	err = DB.Model(&r).Association("Managers").Append(u).Error
	if err != nil {
		return  err
	}
	return nil
}

// 取消管理员
func RoomCancelAdministrator(room string, administrator , owner string) error {
	if _,_, err := verificationOwner(room, owner); err != nil {
		return err
	}
	if administrator == owner {
		return internel.RoomCantCancelOwner
	}

	a, r, err := verificationAdministrator(room, administrator)
	if err != nil {
		return err
	}


	 err = DB.Model(&r).Association("Managers").Delete(a).Error
	 return err
}

// 转让群
func RoomTransferRoom(room string, transferee , owner string) error {
	// 确认是群主
	o, r, err := verificationOwner(room, owner)
	if err != nil {
		return err
	}
	// 确认被转让人是群中成员
	u, _, err := verificationMember(room , transferee)
	if err != nil {
		return err
	}

	// 判断 是否是被转让的是否是管理员
	_, _, err = verificationAdministrator(room, transferee)
	if err == nil || err == internel.RoomAdministratorNotExited { // 不是管理员 设置成管理员
		tx := DB.Begin()
		err = DB.Model(&r).Association("Managers").Append(u).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		err = DB.Model(&r).Update("owner", transferee).Error
		if err != nil {
			tx.Rollback()
			return  err
		}
		err = DB.Model(&r).Association("Managers").Delete(o).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	} else if err != nil { // 是管理员直接设置成群主
		tx := DB.Begin()
		err = DB.Model(&r).Update("owner", transferee).Error
		if err != nil {
			tx.Rollback()
			return  err
		}
		err = DB.Model(&r).Association("Managers").Delete(o).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	}
	return nil
}

func AllUserInRoom(room string) ([]data.User, error) {
	var r data.Room
	var users = []data.User{}
	if err := DB.Where(&data.Room{RoomDisplayID: room}).First(&r).Error; err != nil {
		return users, err
	}
	if err := DB.Model(&r).Association("Users").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}