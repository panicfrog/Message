package data

import "github.com/jinzhu/gorm"

type UserStatus int
const (
	UserStatusOnline  = 0x01
	UserStatusOffLine = 0x02
	UserStatusBusy    = 0x03
)

type User struct {
	gorm.Model
	Account string      `gorm:"type:varchar(50);not null; unique"`
	Passwd string       `gorm:"type:varchar(50);not null"`
	Frients []User      `gorm:"many2many:friendship;association_jointable_foreignkey:friend_id"`
	Status UserStatus   `gorm:"type:integer;default:'2'"`
	Rooms []Room        `gorm:"many2many:room_users"`
}