package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

var db *gorm.DB

func main() {
	var err error
	// Connect to the database
	dsn := "host=localhost user=postgres password=postgres dbname=mydatabase port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Resource{})

	// Set up Gin
	r := gin.Default()
	r.GET("/resources/:name", getResource)
	r.POST("/resources", createResource)
	r.PUT("/resources/:name", updateResource)
	r.DELETE("/resources/:name", deleteResource)

	r.Run() // By default, it serves on :8080
}

func getResource(c *gin.Context) {
	name := c.Param("name")
	var resource Resource
	result := db.Where("name = ?", name).First(&resource)

	if result.Error != nil {
		c.JSON(404, gin.H{"message": "Resource not found"})
		return
	}

	c.JSON(200, resource)
}

func createResource(c *gin.Context) {
	var newResource Resource
	if err := c.BindJSON(&newResource); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newResource)
	c.JSON(201, newResource)
}

func updateResource(c *gin.Context) {
	name := c.Param("name")
	var updatedResource Resource

	if err := c.BindJSON(&updatedResource); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var resource Resource
	if db.Where("name = ?", name).First(&resource).Error != nil {
		c.JSON(404, gin.H{"message": "Resource not found"})
		return
	}

	db.Model(&resource).Updates(updatedResource)
	c.JSON(200, updatedResource)
}

func deleteResource(c *gin.Context) {
	name := c.Param("name")
	var resource Resource
	if db.Where("name = ?", name).First(&resource).Error != nil {
		c.JSON(404, gin.H{"message": "Resource not found"})
		return
	}

	db.Delete(&resource)
	c.JSON(200, gin.H{"message": "Resource deleted"})
}
