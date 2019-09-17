package data

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model              `json:"-"`
	RoomDisplayID string    `gorm:"type:varchar(50);not null;unique" json:"room_identifier"`
	RoomName string         `gorm:"type:varchar(50);not null" json:"room_name"`
	Owner string            `gorm:"type:varchar(50);not null" json:"owner"`
	Managers []User         `gorm:"many2many:room_managers" json:"managers,omitempty"`
	Users []User            `gorm:"many2many:room_users" json:"users,omitempty"`
}