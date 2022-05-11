package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here

	// buat router method GET dengan path /movie

	// ada 2 query parameter yaitu: genre dan country
	// buat defaultquery kalau genre tidak dikirim
	// buat query kalau parameter country tidak dikirim maka return StatusOK tanpa menampilkan data
	// buat query kalau parameter country dikirim maka menampilkan data

	router.GET("/movie", func(c *gin.Context) {
		genre := c.DefaultQuery("genre", "general")
		country := c.Query("country")

		if country == "" {
			c.String(http.StatusOK, "Here result of query of movie with genre %s", genre)
			return
		}
		c.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", genre, country)
	})

	return router

}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
