package data

import "github.com/jinzhu/gorm"

type InvitationType int
type Invitation struct {
	gorm.Model
	Inviter string          `gorm:"type:varchar(100);not null"`
	Invitee string          `gorm:"type:varchar(100);not null"`
	Type InvitationType     `gorm:"type:integer;not null"`
}