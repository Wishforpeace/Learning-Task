package main

import "github.com/jinzhu/gorm"

var (
	DB        *gorm.DB
	chanVideo chan
)
var Video struct {
	ChannelId
}