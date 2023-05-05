package routes

import (
    controller "github.com/Brekke-Green/beta-book/controllers" 
//    middleware "github.com/Brekke-Green/beta-book/middleware"
    "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
    //    incomingRoutes.Use(middleware.Authenticate())
    incomingRoutes.GET("/users", controller.GetUsers())
    incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
