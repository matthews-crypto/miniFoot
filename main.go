package main

import (
	"miniFoot/routes"
	"miniFoot/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connexion à la base de données
	utils.ConnectDB()

	// Initialisation du routeur Gin
	router := gin.Default()

	// Configuration de la session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Chargement des templates HTML
	router.LoadHTMLGlob("templates/*")

	// Routes
	routes.InitializeRoutes(router)

	// Route pour servir les fichiers statiques (CSS, JS, images, etc.)
	router.Static("/public", "./public")

	// Route pour afficher le formulaire d'inscription
	router.GET("/register.html", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	// Route pour afficher le formulaire de connexion
	router.GET("/login.html", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	// Démarrage du serveur
	router.Run(":8080")
}
