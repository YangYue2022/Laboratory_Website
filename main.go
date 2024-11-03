package main

import (
	"UISwebsite_backend/models"
	"UISwebsite_backend/pkg/logging"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"

	"UISwebsite_backend/pkg/setting"
	"UISwebsite_backend/routers"
)

func main() {
	setting.Setup()
	logging.Setup()
	models.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
