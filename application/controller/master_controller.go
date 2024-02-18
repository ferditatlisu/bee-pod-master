package controller

import (
	"bee-pod-master/application/service"
	"bee-pod-master/pkg/config"
	"bee-pod-master/pkg/logger"
	"bee-pod-master/pkg/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type masterController struct {
	ms service.MasterService
}

type MasterController interface {
	Copy() gin.HandlerFunc
	Search() gin.HandlerFunc
	Partition() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetConfig() gin.HandlerFunc
	GetEnv() gin.HandlerFunc
}

func NewMasterController(s service.MasterService) MasterController {
	return &masterController{ms: s}
}

// Copy  @Tags pod-master-controller
// @Accept json
// @Produce json
// @Success 200
// @Param 	podBody body 	dto.CopyRequest 	true 	"CopyRequest"
// @Router /master/copy [post]
func (m *masterController) Copy() gin.HandlerFunc {
	return func(context *gin.Context) {
		copyReq, err := GetRequestBody[dto.CopyRequest](context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		copyReq.UpdateCopyRequest()
		err = m.ms.Copy(copyReq)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, "OK")
	}
}

// Search  @Tags pod-master-controller
// @Accept json
// @Produce json
// @Success 200
// @Param 	podBody body 	dto.SearchRequest 	true 	"SearchRequest"
// @Router /master/search [post]
func (m *masterController) Search() gin.HandlerFunc {
	return func(context *gin.Context) {
		searchReq, err := GetRequestBody[dto.SearchRequest](context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		searchReq.UpdateSearchRequest()
		err = m.ms.Search(searchReq)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, "OK")
	}
}

// Search  @Tags pod-master-controller
// @Accept json
// @Produce json
// @Success 200
// @Param 	podBody body 	dto.PartitionRequest 	true 	"PartitionRequest"
// @Router /master/partition [post]
func (m *masterController) Partition() gin.HandlerFunc {
	return func(context *gin.Context) {
		req, err := GetRequestBody[dto.PartitionRequest](context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		req.UpdateRequest()
		err = m.ms.Partition(req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, "OK")
	}
}

// Delete @Tags pod-master-controller
// @Accept json
// @Produce json
// @Success 200
// @Param   key      query     string     true  "pod name"
// @Router /master [delete]
func (m *masterController) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		podName := context.Query("key")
		logger.Logger().Info(fmt.Sprintf("Delete request is incoming with this key: %s", podName))
		err := m.ms.Delete(podName)
		if err != nil {
			logger.Logger().Error("delete operation failed: ", zap.Error(err))
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, "OK")
	}
}

// GET @Tags pod-master-controller
// @Accept json
// @Produce json
// @Success 200
// @Router /master/getenv [get]
func (m *masterController) GetEnv() gin.HandlerFunc {
	return func(context *gin.Context) {
		envValue := os.Getenv("KAFKA_CONFIGS")
		context.JSON(http.StatusOK, envValue)
		return
	}
}

// GET @Tags pod-master-controller
// @Accept json
// @Produce json
// @Success 200
// @Router /master/getconfig [get]
func (m *masterController) GetConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		envValue := os.Getenv("KAFKA_CONFIGS")
		var kafkaConfigs []config.KafkaConfig
		err := json.Unmarshal([]byte(envValue), &kafkaConfigs)
		if err != nil {
			logger.Logger().Error("marshall problem, ", zap.Error(err))
		}
		context.JSON(http.StatusOK, kafkaConfigs)
		return
	}
}

//a := os.Getenv("KAFKA_CONFIGS")
//
