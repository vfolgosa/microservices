package main

import (
	"github.com/vfolgosa/microservices/back/src/config"
	"github.com/vfolgosa/microservices/back/src/config/dao"
	"github.com/vfolgosa/microservices/back/src/models"

	"github.com/gin-gonic/gin"
)

var mdao = dao.MoviesDAO{}
var mconf = config.Config{}

func init() {
	mconf.Read()

	mdao.Server = mconf.Server
	mdao.Database = mconf.Database
	mdao.Connect()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/v1/movies", func(c *gin.Context) {
		movies, err := mdao.GetAll()
		if err != nil {
			AbortMsg(500, err, c)
		}
		c.JSON(200, gin.H{
			"movies": movies,
		})

	})

	r.POST("/api/v1/movies", func(c *gin.Context) {
		var movie models.Movie
		c.BindJSON(&movie)
		err := mdao.Create(movie)
		if err != nil {
			AbortMsg(500, err, c)
		}
	})

	r.Run()
}

func AbortMsg(code int, err error, c *gin.Context) {
	c.String(code, "Oops! Please retry.")
	c.Error(err)
	c.Abort()
}
