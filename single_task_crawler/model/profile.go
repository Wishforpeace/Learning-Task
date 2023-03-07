package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Gender    string
	Age       int
	Height    int
	Income    string
	Marriage  string
	Address   string
	Education string
}
