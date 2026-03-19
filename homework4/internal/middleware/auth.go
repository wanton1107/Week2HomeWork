package middleware

import (
	"homework4/pkg/apperror"
	"strings"

	"homework4/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
)

func AuthMiddleware(jwtMgr *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(apperror.Unauthorized("未传令牌", nil))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.Error(apperror.Unauthorized("令牌格式异常", nil))
			c.Abort()
			return
		}

		claims, err := jwtMgr.ParseToken(parts[1])
		if err != nil {
			c.Error(apperror.Unauthorized("令牌已过期", err))
			c.Abort()
			return
		}

		// 当前登录用户ID
		c.Set(ContextUserIDKey, claims.UserID)
		c.Next()
	}
}
