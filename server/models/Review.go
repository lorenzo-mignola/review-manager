package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Text string
	Star int
}
