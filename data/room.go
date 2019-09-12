package data

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	RoomName string         `gorm:"type:varchar(50);not null" json:"room_name"`
	Owner string            `gorm:"type:varchar(50);not null" json:"owner"`
	Managers []User         `gorm:"many2many:room_managers" json:"managers"`
	Users []User            `gorm:"many2many:room_users" json:"users"`
}