package middleware

import (
	"bee-pod-master/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
)

func GlobalRecover() gin.HandlerFunc {
	return RecoveryWithWriter()
}

func RecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				e := err.(error)
				logger.Logger().Error(fmt.Sprintf("Unhandled error: %v", string(debug.Stack())), zap.Error(e))
				recoveryHandler(c, e)
			}
		}()
		c.Next()
	}
}

func recoveryHandler(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err.Error())
}
