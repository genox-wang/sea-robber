package controller

import (
	"net/http"
	"sea-robber-api/app/model"

	"github.com/gin-gonic/gin"
)

// RankAllGet is a function for get rank for all users
func RankAllGet(c *gin.Context) {
	users := make([]model.User, 0)

	model.DB.Select("uuid, display_name, battle_score, battle_ship_id").Order("battle_score desc").Limit(100).Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"rank": users,
	})
}
