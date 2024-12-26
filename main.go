package main

import (
	"blogapp/config"
	"blogapp/controller"
	model "blogapp/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &controller.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.JWTKey, nil
		})

		if err != nil {
			var errMsg string
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				errMsg = "Invalid token signature"
			} else {
				errMsg = "Error parsing token"
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func main() {
	// Database connection
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the schema
	config.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	// Initialize Gin router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Blog Home",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
		})
	})

	r.GET("/create_post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_post.html", gin.H{"title": "Create Post"})
	})

	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts.html", gin.H{"title": "Manage Posts"})
	})

	r.GET("/game_center", func(c *gin.Context) {
		c.HTML(http.StatusOK, "game_center.html", nil)
	})

	// Read more route
	r.GET("/read-more/:id", func(c *gin.Context) {
		id := c.Param("id")
		// Retrieve the post by ID (simulate or connect to database)
		post := map[string]string{
			"id":      id,
			"title":   "Sample Post Title " + id,
			"content": "This is the full content of the post with ID " + id,
		}
		c.HTML(http.StatusOK, "read_more.html", gin.H{
			"title":   post["title"],
			"content": post["content"],
		})
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		api.GET("/posts", controller.GetPosts)
		api.GET("/posts/:id", controller.GetPost)

		// Protected routes
		protected := api.Group("/")
		protected.Use(authMiddleware())
		{
			protected.POST("/posts", controller.CreatePost)
			protected.DELETE("/posts/:id", controller.DeletePost)
		}
	}

	r.Run(":9000")
}
