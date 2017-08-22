package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name  string  `gorm:"not null" json:"name"`
	Image []Image `gorm:"many2many:images_tags" json:"-"`
}
