package routers

import (
	"GinExample/api"

	"github.com/gin-gonic/gin"
)

func test(r *gin.Engine) {
	r.GET("/test", api.Test)
}
