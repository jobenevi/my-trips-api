package routes

import (
	config "my-trips-api/configs"
	"my-trips-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()
	r.Use(config.ConfigureCORS())
	r.GET("/testimonials", controllers.GetAllTestimonials)
	r.POST("/testimonials", controllers.CreateTestimonial)
	r.PUT("/testimonials/:id", controllers.UpdateTestimonialByID)
	r.DELETE("/testimonials/:id", controllers.DeleteTestimonialByID)
	r.GET("/testimonials-home", controllers.GetTestimonialsHome)
	r.Run()
}
