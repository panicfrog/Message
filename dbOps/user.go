package dbOps

import (
	"github.com/jinzhu/gorm"
	"log"
	"message/data"
	"message/internel"
)

func AddUser(account string, passwd string) error {
	pw := internel.Md5(passwd)
	var user data.User
	dbErr := DB.Where(&data.User{Account:account}).First(&user).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		log.Fatal(dbErr)
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

func FindUser(accout string) (data.User, error) {
	var user data.User
	dbErr := DB.Where(&data.User{Account: accout}).First(&user).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		return user, dbErr
	}
	if dbErr != nil && dbErr == gorm.ErrRecordNotFound {
		log.Fatal(dbErr)
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
		log.Fatal(dbErr)
	}
	err := DB.Model(&user).Update("passwd", internel.Md5(passwd)).Error
	if err != nil {
		return err
	}
	return  nil
}