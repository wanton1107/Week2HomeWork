package middleware

import (
	"homework4/pkg/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		appErr := apperror.As(err)
		if appErr.Code >= http.StatusInternalServerError {
			log.WithError(err).Error("请求失败")
		} else {
			log.WithError(err).Warn("请求拒绝")
		}
		c.JSON(appErr.Code, errorResponse{Message: appErr.Message})
	}
}
