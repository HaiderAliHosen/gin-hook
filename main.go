package main

import (
	"fmt"

	"github.com/HaiderAliHosen/ginhook/hook"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("cannot connect to database")
		return
	}
	fmt.Println("server started............")

	initializationFunc(v1, session)
	//router.GET("/hook", func(c *gin.Context)

	router.Run(":3000")
}
func initializationFunc(router *gin.RouterGroup, dbs *mgo.Session) {
	hook.Init(router, dbs)
}
