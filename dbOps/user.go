package dbOps

import (
	"github.com/jinzhu/gorm"
	"message/data"
	"message/internel"
)

func AddUser(account string, passwd string) error {
	pw := internel.Md5(passwd)
	var user data.User
	dbErr := DB.Where(&data.User{Account:account}).First(&user).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		return dbErr
	}

	if dbErr == nil {
		return internel.DBErrorExited
	}

	user = data.User{ Account: account, Passwd: pw }
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return  nil
}

func VerificationUser(account string, passwd string) (data.User, error) {
	var user data.User
	dbErr := DB.Where(&data.User{Account:account, Passwd: internel.Md5(passwd)}).First(&user).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		return user, dbErr
	}
	if dbErr != nil && dbErr == gorm.ErrRecordNotFound {
		return user, internel.DBErrorExited
	}
	return  user, nil
}

func FindUser(accout string) (data.User, error) {
	var user data.User
	dbErr := DB.Where(&data.User{Account: accout}).First(&user).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		return user, dbErr
	}
	if dbErr != nil && dbErr == gorm.ErrRecordNotFound {
		return user, internel.DBErrorExited
	}
	return user, nil
}

func ChangePasswd(accout, oldPasswd, passwd string) error {
	var user data.User
	dbErr := DB.Where(&data.User{Account:accout, Passwd: internel.Md5(oldPasswd)}).First(&user).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		return  dbErr
	}
	if dbErr != nil && dbErr == gorm.ErrRecordNotFound {
		return internel.DBErrorExited
	}
	err := DB.Model(&user).Update("passwd", internel.Md5(passwd)).Error
	if err != nil {
		return err
	}
	return  nil
}

func AddFriend(account string, frientAccount string) error {
	var user data.User
	var friend data.User

	if account == frientAccount {
		return internel.UserCannotAddFriendWithSelf
	}
	if err := DB.Where(&data.User{Account: account}).First(&user).Error; err != nil {
		return err
	}
	if err := DB.Where(&data.User{Account: frientAccount}).First(&friend).Error; err != nil {
		return err
	}
	_user, _friend := user, friend

	err := DB.Model(&user).Association("Frients").Find(&_friend).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	if err == nil {
		return internel.UserYouAreAleadyFriend
	}
	err = DB.Model(&friend).Association("Frients").Find(&_user).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	if err == nil {
		return internel.UserYouAreAleadyFriend
	}

	tx := DB.Begin()
	if err := DB.Model(&user).Association("Frients").Append(friend).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := DB.Model(&friend).Association("Frients").Append(user).Error; err != nil {
		tx.Rollback()
		return  err
	}
	tx.Commit()
	return nil
}

func RemoveFriend(account string, friendAccount string) error {
	var user data.User
	var friend data.User

	if err := DB.Where(&data.User{Account: account}).First(&user).Error; err != nil {
		return err
	}
	if err := DB.Where(&data.User{Account: friendAccount}).First(&friend).Error; err != nil {
		return err
	}
	tx := DB.Begin()
	if err := DB.Model(&user).Association("Frients").Delete(friend).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := DB.Model(&friend).Association("Frients").Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func AllFriends(account string) ([]data.User, error) {
		var (
			user data.User
			users []data.User
	)

	if err := DB.Where(&data.User{Account: account}).First(&user).Error; err != nil {
		return  users, err
	}

	if err := DB.Model(&user).Association("Frients").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []data.User{}, nil
		}
		return users, err
	}
	return users, nil
}

func AllRooms(account string) ([]data.Room, error) {
	rooms := []data.Room{}
	var user data.User
	if err := DB.Where(&data.User{Account: account}).First(&user).Error; err != nil {
		return  rooms, err
	}
	if err := DB.Model(&user).Association("Rooms").Find(&rooms).Error; err != nil {
		return rooms, err
	}
	return rooms, nil
}
