package middlewares

import (
	"github.com/Darklabel91/API_Names/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//ValidateName validates :name param. It must not contain numbers or spaces
func ValidateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to retrieve the ":name" parameter from the request context
		name := c.Param("name")

		if len(name) < 3 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name must be at least 3 characters"})
			return
		}

		// Check if the name contains whitespace
		if strings.Contains(name, " ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name must contain a single word with no spaces"})
			return
		}

		// Check if the name contains any numbers
		if _, err := strconv.Atoi(name); err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name must not contain any numbers"})
			return
		}

		c.Next()
	}
}

//ValidateNameJSON validates JSON on models.NameType body
func ValidateNameJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		var name models.NameType
		err := c.Bind(&name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid JSON request body"})
			return
		}
		c.Set("name", name)
		c.Next()
	}
}
