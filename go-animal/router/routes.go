package Routes

import (
	"../api"
	"../config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/dogs", config.DOGIMAGES)

	v1 := r.Group("/v1")
	{
		v1.GET("dog", api.GetDogs)
	}

	return r
}
