package models

import "gorm.io/gorm"

type Testimonial struct {
	gorm.Model
	Name         string `json:"name"`
	Content      string `json:"content"`
	ProfileImage string `json:"profile_image"`
}
