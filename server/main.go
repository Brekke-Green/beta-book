package main

import (
        "time"
        "net/http"

        routes "github.com/Brekke-Green/beta-book/routes" 
        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/cors"
)

type climb struct {
        Name   string `json:"name"`
        Setter string `json:"setter"`
        Grade  string `json:"grade"`
}

var climbs = []climb{
        {Name: "Brick", Setter: "Brekke", Grade: "V4"},
        {Name: "Slick", Setter: "Brekke", Grade: "V6"},
        {Name: "Trick", Setter: "Brekke", Grade: "V2"},
}

func main() {
        router := gin.Default()
        router.Use(cors.New(cors.Config{
                AllowOrigins:     []string{"http://localhost:5173"},
                AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
                AllowHeaders:     []string{"Content-Type"},
                ExposeHeaders:    []string{"Content-Length"},
                AllowCredentials: true,
                MaxAge: 12 * time.Hour,
        }))

        routes.AuthRoutes(router)
        routes.UserRoutes(router)


        router.Run("localhost:8080")
}

func getClimbs(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, climbs)
}

func postClimb(c *gin.Context) {
    
}

func updateClimb(c *gin.Context) {

}

func deleteClimb(c *gin.Context) {

}

func patchClimb(c *gin.Context) {

}
