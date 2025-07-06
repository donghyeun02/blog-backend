package main

import (
	"fmt"
	"log"
	"net/http"

	"blog-backend/internal/infrastructure/config"
	"blog-backend/internal/infrastructure/database"
	"blog-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	gin.SetMode(cfg.Server.Mode)

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	defer func() {
		if err := database.Close(db); err != nil {
			log.Printf("Failed to close database: %v", err)
		}
	}()

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	if cfg.Server.Mode == "debug" {
		router.Use(middleware.DevelopmentCORS())
	} else {
		router.Use(middleware.CORS())
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Blog Backend API Server",
			"version": "v1.0.0",
			"status":  "running",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := router.Group("/api/v1")
	{
		api.GET("/articles", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Articles API - Coming Soon",
			})
		})
	}

	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server on %s", serverAddr)
	log.Printf("Server mode: %s", cfg.Server.Mode)
	log.Printf("Database connected: %s", cfg.Database.Host)
	
	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
