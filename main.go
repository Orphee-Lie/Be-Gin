package main

import (
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware est une fonction qui retourne un gestionnaire Gin Middleware.
func AuthMiddleware() gin.HandlerFunc {
	// In a real-world application, you would perform proper authentication here.
	// For the sake of this example, we'll just check if an API key is present.
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// UserController représente un contrôleur lié à l'utilisateur
// UserController représente un contrôleur lié à l'utilisateur
type UserController struct{}

// GetUserInfo est une méthode de contrôleur pour obtenir des informations sur l'utilisateur
func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")

	// Validation de l'ID de l'utilisateur
	if userID == "" {
		c.JSON(400, gin.H{"error": "ID d'utilisateur non fourni"})
		return
	}

	// Accéder à la base de données pour obtenir les informations de l'utilisateur
	userInfo, err := getUserInfoFromDB(userID)
	if err != nil {
		// Gérer l'erreur lors de l'accès à la base de données
		c.JSON(500, gin.H{"error": "Erreur lors de la récupération des informations de l'utilisateur"})
		return
	}

	// Retourner les informations de l'utilisateur en JSON
	c.JSON(200, userInfo)
}

// DeleteUser est une méthode de contrôleur pour supprimer un utilisateur
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Validation de l'ID de l'utilisateur
	if userID == "" {
		c.JSON(400, gin.H{"error": "ID d'utilisateur non fourni"})
		return
	}

	// Supprimer l'utilisateur de la base de données
	err := deleteUserFromDB(userID)
	if err != nil {
		// Gérer l'erreur lors de la suppression de l'utilisateur
		c.JSON(500, gin.H{"error": "Erreur lors de la suppression de l'utilisateur"})
		return
	}

	// Retourner une réponse de succès
	c.JSON(200, gin.H{"message": "Utilisateur supprimé avec succès"})
}

// deleteUserFromDB est une fonction utilitaire pour supprimer un utilisateur de la base de données
func deleteUserFromDB(userID string) error {
	// Implémentez la logique pour supprimer l'utilisateur de la base de données
	// ...

	// Pour l'exemple, retournez nil en cas de réussite
	return nil
}

// getUserInfoFromDB est une fonction utilitaire pour obtenir les informations de l'utilisateur depuis la base de données
func getUserInfoFromDB(userID string) (gin.H, error) {
	// Implémentez la logique pour accéder à la base de données et récupérer les informations de l'utilisateur
	// ...

	// Pour l'exemple, retournez des informations statiques
	return gin.H{"id": userID, "name": "John Doe", "email": "john@example.com"}, nil
}

func main_() {
	router := gin.Default()

	public := router.Group("/public")
	{
		public.GET("/info", func(c *gin.Context) {
			c.String(200, "Public information")
		})
		public.GET("/products", func(c *gin.Context) {
			c.String(200, "Public product list")
		})
	}

	private := router.Group("/private")
	private.Use(AuthMiddleware())
	{
		private.GET("/data", func(c *gin.Context) {
			c.String(200, "Private data accessible after authentication")
		})
		private.POST("/create", func(c *gin.Context) {
			c.String(200, "Create a new resource")
		})
	}

	// Use our custom authentication middleware for a specific group of routes
	authGroup := router.Group("/api")
	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
	}

	userController := &UserController{}

	// Route using the UserController
	router.GET("/user/:id", userController.GetUserInfo)

	// Route with URL parameters
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "User ID: "+id)
	})

	// Route with query parameters
	router.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default-value")
		c.String(200, "Search query: "+query)
	})

	router.Run(":8080")
}
