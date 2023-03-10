package routes

import (
	controller "pg-management/controllers"

	"github.com/gin-gonic/gin"
)

func pg_owner_routes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.Getowner_byID())
	incomingRoutes.GET("/foods/:food_id", controller.Get_Alluser())
	incomingRoutes.POST("/foods", controller.Sign_up())
	incomingRoutes.POST("/foods/:food_id", controller.Login())
	incomingRoutes.PATCH("/foods/:food_id", controller.Update_owner())
	incomingRoutes.DELETE("/foods/:food_id", controller.Delete_owner())
	incomingRoutes.GET("/foods/:food_id", controller.See_booking())
}
