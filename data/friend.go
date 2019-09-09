package data

import "github.com/jinzhu/gorm"

type Friend struct {
	gorm.Model        `json:"-"`
	User string       `gorm:"type:varchar(50);not null" json:"user"`
	Friend string     `gorm:"type:varchar(50);not null" json:"friend"`
}