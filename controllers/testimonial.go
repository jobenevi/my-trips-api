package controllers

import (
	"my-trips-api/database"
	"my-trips-api/models"
	"my-trips-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTestimonials(c *gin.Context) {
	testimonials, err := services.GetAllTestimonials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, testimonials)
}

func CreateTestimonial(c *gin.Context) {
	var testimonial models.Testimonial
	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTestimonial(&testimonial); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, testimonial)
}

func UpdateTestimonialByID(c *gin.Context) {
	var testimonial models.Testimonial
	id := c.Param("id")

	if err := database.DB.First(&testimonial, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}

	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateTestimonialByID(&testimonial); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testimonial)
}

func DeleteTestimonialByID(c *gin.Context) {
	var testimonial models.Testimonial

	if err := services.DeleteTestimonialByID(&testimonial); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Testimonial deleted successfully"})
}
