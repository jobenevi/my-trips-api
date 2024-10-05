package controllers

import (
	"my-trips-api/database"
	"my-trips-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTestimonials(c *gin.Context) {
	var testimonials []models.Testimonial
	if err := database.DB.Find(&testimonials).Error; err != nil {
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

	if err := database.DB.Create(&testimonial).Error; err != nil {
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

	if err := database.DB.Save(&testimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, testimonial)

}

func DeleteTestimonialByID(c *gin.Context) {

	var testimonial models.Testimonial
	id := c.Param("id")

	database.DB.Delete(&testimonial, id)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Testimonial deleted successfully"})

}
