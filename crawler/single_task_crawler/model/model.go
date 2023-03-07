package model

import "gorm.io/gorm"

var DB *gorm.DB

func CreateTable() error {
	if err := DB.AutoMigrate(&Profile{}); err != nil {
		return err
	}
	return nil
}
func (profile *Profile) CreateUser() error {
	if err := DB.Create(profile).Error; err != nil {
		return err
	}
	return nil
}
