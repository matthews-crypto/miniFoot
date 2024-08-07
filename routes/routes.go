package routes

import (
	"miniFoot/controllers"
	"miniFoot/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.GET("/admin/dashboard", middleware.AuthRequired, func(c *gin.Context) {
		c.HTML(200, "admin_dashboard.html", nil)
	})
	router.GET("/reserviste/dashboard", middleware.AuthRequired, func(c *gin.Context) {
		c.HTML(200, "reserviste_dashboard.html", nil)
	})
	router.GET("/gerant/dashboard", middleware.AuthRequired, func(c *gin.Context) {
		c.HTML(200, "gerant_dashboard.html", nil)
	})
	router.GET("/proprio/dashboard", middleware.AuthRequired, func(c *gin.Context) {
		c.HTML(200, "proprio_dashboard.html", nil)
	})

	// Groupe pour les routes utilisateur authentifié
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthRequired)
	{
		userGroup.POST("/edit_password", controllers.EditPassword)
	}
}
