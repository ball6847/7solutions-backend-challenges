package api

import (
	"7solution/api/controller"
	"7solution/api/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
)

func Serve() {
	router := gin.New()
	router.Use(gin.Logger())

	// services
	http := req.C()
	baconipsum := service.NewBaconipsumClient(http)
	wordCounter := service.WordCounter{}

	// controllers
	beefController := controller.NewBeefController(&wordCounter, baconipsum)

	// routings
	router.GET("/beef/summary", beefController.SummaryHandler)

	err := router.Run(":5555")
	if err != nil {
		log.Fatalln("Error starting api server ", err)
	}
}
