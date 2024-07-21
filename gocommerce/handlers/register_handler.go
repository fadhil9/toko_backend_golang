package handlers

import (
	"main.go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message":"invalid input"})
			return
		}

		//check apakah username sudah ada?
		var existingUser models.User
		if err := db.Where("username = ?", input.username).First(&existingUser).Error; err == nil {
			c.JSON(400, gin.H{"message":"username already exist"})
		}

		//create new user
		newUser := models.User{
			Username : input.username,
			Password : input.Password, // harus dihash dlu sebelum distore
		}
		
		if err := db.Create(&newUser).Error; err !=nil {
			c.JSON(500, gin.H{"message":"internal server error"})
			return
		}

		token, err:= CreateToken(newUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"Message":"Internal server error"})
			return
		}

		c.JSON(200, gin.H{"Message":"User registered succesfully", "token":token})
	}
}