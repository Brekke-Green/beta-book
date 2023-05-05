package routes

import (
    "github.com/gin-gonic/gin"
)

func climbRoutes() {
    
        router.GET("/climbs", getBoulders)
        router.GET("/climbs/:climb", getClimb)
        router.POST("/climbs/:climb", postClimb)
        router.PUT("/climbs/:climb", updateClimb)
        router.PATCH("/climbs/:climb", patchClimb)
        router.DELETE("/climbs/:climb", deleteClimb)
    } 
