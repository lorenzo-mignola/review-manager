package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	text string
	Star int
}
