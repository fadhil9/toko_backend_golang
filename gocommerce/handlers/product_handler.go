package handlers

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main.go/models"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		//menambahkan satu goroutine ke waitgroup
		wg.Add(1)

		//memulai goroutine  untuk melakukan  operasi yang membutuhkan  waktu lama
		go func() {
			defer wg.Done() //menandai bahwa goroutine telah selesai
			db.Find(&products)
		}()

		//menunggu goroutine selesai
		wg.Wait()

		c.JSON(200, products)
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product Not Found"})
			return
		}
		c.JSON(200, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "invalid input"})
			return
		}
		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(400, gin.H{"Message": "Product Not Found"})
			return
		}

		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"Message": "invalid input"})
			return
		}

		db.Model(&product).Updates(input)
		c.JSON(200, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product Not Found"})
		}

		db.Delete(&product)
		c.JSON(200, gin.H{"message": "Product Deleted"})
	}
}
