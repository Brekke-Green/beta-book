package main

import (
        "time"
        "net/http"

        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/cors"
)

type boulder struct {
        Name   string `json:"name"`
        Setter string `json:"setter"`
        Grade  string `json:"grade"`
}

var boulders = []boulder{
        {Name: "Brick", Setter: "Brekke", Grade: "V4"},
        {Name: "Slick", Setter: "Brekke", Grade: "V6"},
        {Name: "Trick", Setter: "Brekke", Grade: "V2"},
}

func main() {
        router := gin.Default()


        router.Use(cors.New(cors.Config{
                AllowOrigins:     []string{"http://localhost:5173"},
                AllowMethods:     []string{"GET"},
                AllowHeaders:     []string{"Content-Type"},
                ExposeHeaders:    []string{"Content-Length"},
                AllowCredentials: true,
                MaxAge: 12 * time.Hour,
        }))
        router.GET("/boulders", getBoulders)
        router.Run("localhost:8080")
}

func getBoulders(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, boulders)
}
