package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) { // TODO: replace this
		param := c.Param("id")
		id, _ := strconv.Atoi(param)

		if _, ok := data[id]; !ok {
			c.String(http.StatusNotFound, "data not found")
			return
		}

		output := fmt.Sprintf("Name: %s, Country: %s, Age: %d", data[id].Name, data[id].Country, data[id].Age)
		c.String(http.StatusOK, output)
	}
}

func GetRouter() *gin.Engine {
	//return &gin.Engine{} // TODO: replace this
	router := gin.Default()
	router.GET("/profile/:id", ProfileHandler())
	return router

}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
