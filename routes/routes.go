package routes

import (
	"my-trips-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	router := gin.Default()
	router.GET("/testimonials", controllers.GetAllTestimonials)
	router.POST("/testimonials", controllers.CreateTestimonial)
	router.PATCH("/testimonials/:id", controllers.UpdateTestimonialByID)
	router.DELETE("/testimonials/:id", controllers.DeleteTestimonialByID)
	router.Run()
}
