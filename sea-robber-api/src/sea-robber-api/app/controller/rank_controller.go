package controller

import (
	"net/http"
	"sea-robber-api/app/model"
	"sea-robber-api/app/utils/cache"

	"github.com/gin-gonic/gin"
)

const (
	// CACHE_RANK_ALL  rank cache key
	CACHE_RANK_ALL = "cRA"
)

// RankAllGet is a function for get rank for all users
func RankAllGet(c *gin.Context) {
	u, found := cache.Cache.Get(CACHE_RANK_ALL)

	if found {
		if users, ok := u.([]model.User); ok {
			c.JSON(http.StatusOK, gin.H{
				"rank": users,
			})
			return
		}
	}

	users := make([]model.User, 0)

	if err := model.DB.Select("uuid, display_name, battle_score, battle_ship_id").Order("battle_score desc").Limit(100).Find(&users).Error; err == nil {
		cache.Cache.Set(CACHE_RANK_ALL, users, cache.CACHE_DEFAULT_EXPIRATION)
	}

	c.JSON(http.StatusOK, gin.H{
		"rank": users,
	})
}
