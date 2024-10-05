package services

import (
	"my-trips-api/database"
	"my-trips-api/models"
)

func GetAllTestimonials() ([]models.Testimonial, error) {
	var testimonials []models.Testimonial
	if err := database.DB.Find(&testimonials).Error; err != nil {
		return nil, err
	}
	return testimonials, nil
}

func CreateTestimonial(testimonial *models.Testimonial) error {
	if err := database.DB.Create(testimonial).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTestimonialByID(testimonial *models.Testimonial) error {
	var id = testimonial.ID

	if err := database.DB.First(&testimonial, id).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&testimonial).UpdateColumns(testimonial).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTestimonialByID(testimonial *models.Testimonial) error {

	var id = testimonial.ID

	if err := database.DB.First(&testimonial, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(testimonial).Error; err != nil {
		return err
	}
	return nil
}
