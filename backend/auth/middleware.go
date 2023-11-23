package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func Middleware(c *cache.Cache) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := ctx.GetHeader("authorization")
		if h != "" {
			token := strings.Split(h, "Bearer ")[1]
			nc := SetSessionKey(ctx, token)
			ctx.Request = ctx.Request.WithContext(nc)
		}
		ctx.Next()
	}
}
