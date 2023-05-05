package routes

import (
    "github.com/gin-gonic/gin"
)

func boulderRoutes() {
    
        router.GET("/boulders", getBoulders)
        router.GET("/boulders/:boulder", getBoulder)
        router.POST("/boulders/:boulder", postBoulder)
        router.PUT("/boulders/:boulder", updateBoulder)
        router.PATCH("/boulders/:boulder", patchBoulder)
        router.DELETE("/boulders/:boulder", deleteBoulder)
    } 
