package routes

import (
	"github.com/HironixRotifer/mongodb-service-advertisements/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/advertisement/view", controllers.GetAdvertisementById())
	incomingRoutes.GET("/advertisement/all", controllers.GetAdvertisements())
	incomingRoutes.POST("/advertisement/create", controllers.CreateAdvertisement())
}
