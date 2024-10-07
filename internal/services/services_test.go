package services

import (
	"my-trips-api/internal/database"
	"my-trips-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	database.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB.AutoMigrate(&models.Testimonial{})
}

func TestGetAllTestimonials(t *testing.T) {
	setupTestDB()

	database.DB.Create(&models.Testimonial{Content: "Testimonial 1"})
	database.DB.Create(&models.Testimonial{Content: "Testimonial 2"})

	testimonials, err := GetAllTestimonials()
	assert.NoError(t, err)
	assert.Len(t, testimonials, 2)
}

func TestCreateTestimonial(t *testing.T) {
	setupTestDB()

	testimonial := &models.Testimonial{Content: "New Testimonial"}
	err := CreateTestimonial(testimonial)
	assert.NoError(t, err)

	var result models.Testimonial
	database.DB.First(&result, testimonial.ID)
	assert.Equal(t, "New Testimonial", result.Content)
}

func TestUpdateTestimonialByID(t *testing.T) {
	setupTestDB()

	testimonial := &models.Testimonial{Content: "Original Content"}
	database.DB.Create(testimonial)

	testimonial.Content = "Updated Content"
	err := UpdateTestimonialByID(testimonial)
	assert.NoError(t, err)
}

func TestDeleteTestimonialByID(t *testing.T) {
	setupTestDB()

	testimonial := &models.Testimonial{Content: "Testimonial to be deleted"}
	database.DB.Create(testimonial)

	err := DeleteTestimonialByID(testimonial)
	assert.NoError(t, err)

	var result models.Testimonial
	err = database.DB.First(&result, testimonial.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestGetTestimonialsHomeRandom(t *testing.T) {
	setupTestDB()

	database.DB.Create(&models.Testimonial{Content: "Testimonial 1"})
	database.DB.Create(&models.Testimonial{Content: "Testimonial 2"})
	database.DB.Create(&models.Testimonial{Content: "Testimonial 3"})
	database.DB.Create(&models.Testimonial{Content: "Testimonial 4"})

	testimonials, err := GetTestimonialsHomeRandom()
	assert.NoError(t, err)
	assert.Len(t, testimonials, 3)

	for _, testimonial := range testimonials {
		assert.NotEmpty(t, testimonial.Content)
	}
}
