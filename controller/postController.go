package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"example.com/todogo/config"
	"example.com/todogo/models"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := config.Database.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "client error",
		})
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	config.Database.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	config.Database.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PutPost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post

	config.Database.First(&post, id)

	config.Database.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	config.Database.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post have been deleted",
	})
}
