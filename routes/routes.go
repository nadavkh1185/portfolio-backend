package routes

import (
	"portfolio-backend/controllers"
	"portfolio-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// PUBLIC ROUTES
	api.POST("/login", controllers.Login)

	api.GET("/profile", controllers.GetProfile)
	api.GET("/projects", controllers.GetProjects)
	api.GET("/projects/:id", controllers.GetProjectByID)
	api.GET("/about", controllers.GetAbout)
	api.GET("/contact", controllers.GetContact)
	api.POST("/contact/message", controllers.SendContactMessage)
	api.GET("/skills", controllers.GetSkills)
	api.GET("/skills/:id", controllers.GetSkillByID)
	api.GET("/experience", controllers.GetExperience)
	api.GET("/experience/:id", controllers.GetExperienceByID)

	// PROTECTED ROUTES
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// Project
	auth.POST("/projects", controllers.CreateProject)
	auth.PUT("/projects/:id", controllers.UpdateProject)
	auth.DELETE("/projects/:id", controllers.DeleteProject)

	// Profile
	auth.PUT("/profile", controllers.UpdateProfile)
	auth.PUT("/about", controllers.UpdateAbout)
	auth.PUT("/contact", controllers.UpdateContact)

	// Skills
	auth.POST("/skills", controllers.CreateSkill)
	auth.PUT("/skills/:id", controllers.UpdateSkill)
	auth.DELETE("/skills/:id", controllers.DeleteSkill)

	// Experience
	auth.POST("/experience", controllers.CreateExperience)
	auth.PUT("/experience/:id", controllers.UpdateExperience)
	auth.DELETE("/experience/:id", controllers.DeleteExperience)

	auth.POST("/upload", controllers.UploadImage)
}
