package main

import (
	"bee-pod-master/application/controller"
	"bee-pod-master/application/service"
	_ "bee-pod-master/docs"
	"bee-pod-master/pkg/client"
	"bee-pod-master/pkg/config"
	"bee-pod-master/pkg/kubernetes"
	"bee-pod-master/pkg/logger"
)

func main() {
	applicationConfig, _ := config.NewApplicationConfig()
	redis := client.NewRedisService(applicationConfig.Redis)
	ks := kubernetes.NewKubernetesService(applicationConfig)
	gin := controller.NewGinServer()
	s := service.NewMasterService(redis, ks)
	cont := controller.NewMasterController(s)
	router := gin.SetupRouter(cont)
	err := router.Run(":8082")
	if err != nil {
		logger.Logger().Error("A gin error occurred" + err.Error())
		panic(err)
	}
}
