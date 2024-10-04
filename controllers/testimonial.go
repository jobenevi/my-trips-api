package controllers

import (
	"my-trips-api/database"
	"my-trips-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTestimonial(c *gin.Context) {

	var testimonial models.Testimonial

	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&testimonial)
	c.JSON(http.StatusCreated, testimonial)

}

func GetAllTestimonials(c *gin.Context) {

	var testimonials []models.Testimonial
	database.DB.Find(&testimonials)
	c.JSON(http.StatusOK, testimonials)

}

func UpdateTestimonialByID(c *gin.Context) {

	var testimonial models.Testimonial
	id := c.Param("id")

	database.DB.First(&testimonial, id)

	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&testimonial).UpdateColumns(testimonial)
	c.JSON(http.StatusOK, testimonial)
}

func DeleteTestimonialByID(c *gin.Context) {

	var testimonial models.Testimonial
	id := c.Param("id")

	database.DB.Delete(&testimonial, id)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Testimonial deleted successfully"})

}
