package dto

import "my-trips-api/internal/models"

type TestimonialHomeDTO struct {
	ID           int    `json:"id"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	ProfileImage string `json:"profile_image"`
}

func ConvertTestimoninalToTestimonialHomeDTO(testimonial *models.Testimonial) TestimonialHomeDTO {
	return TestimonialHomeDTO{
		ID:           int(testimonial.ID),
		Author:       testimonial.Name,
		Content:      testimonial.Content,
		ProfileImage: testimonial.ProfileImage,
	}
}
