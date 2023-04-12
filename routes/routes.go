package routes

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Darklabel91/API_Names/controllers"
	"github.com/Darklabel91/API_Names/middlewares"
	"github.com/Darklabel91/API_Names/models"
	"github.com/gin-gonic/gin"
)

const DOOR = ":8080"
const FILENAME = "Logs.txt"
const MICROSECONDS = 300

func HandleRequests() {
	// Set Gin to release mode.
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router.
	r := gin.Default()

	// Use the OnlyAllowIPs middleware on all routes.
	err := r.SetTrustedProxies(models.IPs)
	if err != nil {
		return
	}

	// Create a file to store the logs.
	file, err := os.OpenFile(FILENAME, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	// Use the Gin logger with a custom log file.
	r.Use(gin.LoggerWithWriter(file))

	// Upload the log file from time to time.
	var log models.Log
	ticker := time.NewTicker(MICROSECONDS * time.Microsecond)
	defer ticker.Stop()
	log.UploadLog(ticker, FILENAME)

	// Routes without middleware.
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Main middleware validation.
	r.Use(middlewares.ValidateAuth())

	// Rate limiter middleware validation
	r.Use(middlewares.RateLimit())

	// Cache the name types.
	cache := &sync.Map{}
	r.Use(cachingNameTypes(cache))

	// CRUD routes.
	r.POST("/name", middlewares.ValidateName(), middlewares.ValidateNameJSON(), controllers.CreateName)
	r.GET("/:id", middlewares.ValidateID(), controllers.GetID)
	r.GET("/name/:name", middlewares.ValidateName(), controllers.GetName)
	r.GET("/metaphone/:name", middlewares.ValidateName(), controllers.GetMetaphoneMatch)
	r.PATCH("/:id", middlewares.ValidateID(), middlewares.ValidateNameJSON(), controllers.UpdateName)
	r.DELETE("/:id", middlewares.ValidateID(), controllers.DeleteName)

	// Start the server.
	err = r.Run(DOOR)
	if err != nil {
		return
	}
}

// Caches the name types.
func cachingNameTypes(cache *sync.Map) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check the cache.
		cacheData, existKey := cache.Load("nameTypes")
		if existKey {
			c.Set("nameTypes", cacheData)
		} else {
			var nameType models.NameType
			allNames, err := nameType.GetAllNames()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error on caching all name types"})
				return
			}
			cache.Store("nameTypes", allNames)
			c.Set("nameTypes", allNames)
		}
		c.Next()
	}
}
