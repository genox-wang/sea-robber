package controller

import (
	"fmt"
	"net/http"
	"sea-robber-api/app/model"

	"github.com/gin-gonic/gin"
)

// UserCreatePost is a function for create user
func UserCreatePost(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err == nil {
		if err := user.Create(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}

// UserUpdatePost is a function for update user
func UserUpdatePost(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "uuid should not be empty!",
		})
		return
	}

	var user model.User

	if err := c.BindJSON(&user); err == nil {
		if err := model.DB.Model(&user).Where("uuid = ?", uuid).Update(user).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}
