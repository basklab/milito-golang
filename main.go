package main

import (
	"github.com/gin-gonic/gin"
	"milito-golang/config"
	"milito-golang/game"
	"milito-golang/store"
	"net/http"
)

func main() {
	db.SetInternal(game.InitialSetup())

	r := gin.Default()
	r.GET("/ping", processPing)
	r.GET("/setup", setupGame)
	_ = r.Run("localhost:8080")
}

var db store.Db

func setupGame(c *gin.Context) {
	db.SetInternal(game.InitialSetup())
	c.JSON(http.StatusOK, db.Internal())
}

func processPing(c *gin.Context) {
	c.JSON(http.StatusOK, config.NewAlexandrianMacedonian())
}
