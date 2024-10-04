package models

type Testimonial struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	ProfileImage string `json:"profile_image"`
}
