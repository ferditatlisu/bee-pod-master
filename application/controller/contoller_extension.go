package controller

import (
	"bee-pod-master/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRequestBody[T any](context *gin.Context) (*T, error) {
	var body T
	err := context.ShouldBindJSON(&body)

	if err != nil {
		logger.GetLogger(context).Error("Invalid Message", zap.Error(err))
		return nil, err
	}
	return &body, nil
}
