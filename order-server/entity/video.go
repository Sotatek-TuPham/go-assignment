package entity

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Rating      uint   `json:"rating"`
}
