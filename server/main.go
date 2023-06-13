package main

import (
	"log"

	"github.com/Brekke-Green/beta-book/db"
	"github.com/Brekke-Green/beta-book/internal/router"
	"github.com/Brekke-Green/beta-book/internal/user"

//	routes "github.com/Brekke-Green/beta-book/routes"
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
        dbConn, err := db.NewDatabase()
        if err != nil {
            log.Fatalf("could not initialize db connection: %s", err)
        }
    
        userRep := user.NewRepository(dbConn.GetDB())
        userSvc := user.NewService(userRep)
        userHandler := user.NewHandler(userSvc)

        router.InitRouter(userHandler)
        router.Start("localhost:8080")
}

// func getClimbs(c *gin.Context) {
//        c.IndentedJSON(http.StatusOK, climbs)
//}

//func postClimb(c *gin.Context) {
    
//}

//func updateClimb(c *gin.Context) {

//}

//func deleteClimb(c *gin.Context) {

//}

//func patchClimb(c *gin.Context) {

//}
