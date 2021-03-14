package main

import (
	"github.com/gin-gonic/gin"
	"milito-golang/config"
	"milito-golang/entity"
	"milito-golang/store"
	"net/http"
)

func main() {
	db.SetInternal(config.InitialSetup())

	r := gin.Default()
	r.GET("/ping", processPing())
	r.GET("/setup", setupGame)
	_ = r.Run("localhost:8080")
}

var db store.Db

func setupGame(c *gin.Context) {
	db.SetInternal(config.InitialSetup())
	c.JSON(http.StatusOK, db.Internal())
}

func processPing() func(c *gin.Context) {

	argg := entity.UnitCard{UnitType: "okkfads"}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, argg)
	}
}
