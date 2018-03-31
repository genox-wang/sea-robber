package controller

import (
	"net/http"
	"sea-robber-api/app/model"
	"sea-robber-api/app/utils/cache"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	// CACHE_RANK_ALL  rank cache key
	CACHE_RANK_ALL = "cRA"
	// CACHE_RANK_LOCK  rank cache lock
	CACHE_RANK_LOCK = "cRL"
)

// CacheRankAll struct for cache
type CacheRankAll struct {
	UsersMap map[int64]model.User
	Rank     []int64
}

// RankAllGet is a function for get rank for all users
func RankAllGet(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Request.Header.Get("Authorization"), 10, 64)
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "user id must > 0",
		})
		return
	}

	page, _ := strconv.ParseInt(c.Query("page"), 10, 32)
	size, _ := strconv.ParseInt(c.Query("size"), 10, 32)

	cR, found := cache.Cache.Get(CACHE_RANK_ALL)

	if found {
		if cRankAll, ok := cR.(CacheRankAll); ok {
			status, data := GetRankFromCache(cRankAll, id, int(page), int(size))
			c.JSON(status, data)
			return
		}
	}

	if cache.Cache.Add(CACHE_RANK_LOCK, true, time.Second*10) == nil {
		defer cache.Cache.Delete(CACHE_RANK_LOCK)
		users := make([]model.User, 0)

		if err := model.DB.Select("id, uuid, display_name, battle_score, battle_ship_id").Order("battle_score desc").Limit(10000).Find(&users).Error; err == nil {
			cRankAll := CacheRankAll{}
			cRankAll.UsersMap = make(map[int64]model.User, 0)
			cRankAll.Rank = make([]int64, 0)
			for idx, u := range users {
				u.Rank = idx + 1
				cRankAll.UsersMap[u.ID] = u
				cRankAll.Rank = append(cRankAll.Rank, u.ID)
			}
			cache.Cache.Set(CACHE_RANK_ALL, cRankAll, cache.CACHE_DEFAULT_EXPIRATION)
			// cache.Cache.Set(CACHE_RANK_ALL, cRankAll, time.Second*5)

			status, data := GetRankFromCache(cRankAll, id, int(page), int(size))
			c.JSON(status, data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusForbidden, nil)
	}
}

// GetRankFromCache  get rank from cache
func GetRankFromCache(cRankAll CacheRankAll, userID int64, page int, size int) (int, gin.H) {
	startIndex := page * size
	endIndex := (page + 1) * size
	rankLen := len(cRankAll.Rank)
	if rankLen < endIndex {
		endIndex = rankLen
	}
	if startIndex >= rankLen {
		return http.StatusBadRequest, gin.H{
			"msg": "already int last page",
		}
	}
	ranks := cRankAll.Rank[startIndex:endIndex]
	users := make([]model.User, 0)
	for _, r := range ranks {
		users = append(users, cRankAll.UsersMap[r])
	}
	userRank := -1
	if user, has := cRankAll.UsersMap[userID]; has {
		userRank = user.Rank
	}

	return http.StatusOK, gin.H{
		"user_rank": userRank,
		"rank":      users,
	}
}
