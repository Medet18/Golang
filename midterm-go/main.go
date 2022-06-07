package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mutex sync.Mutex
)

var store = map[string]string{
	"1": "Ramazan",
	"2": "Medet",
	"3": "Kuanysh",
}

//get request
func getStore(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, store)
}

func getStoreByID(c *gin.Context) {
	mutex.Lock()
	id := c.Param("id")
	value := store[id]
	for _, a := range store {
		if a == value {
			c.IndentedJSON(http.StatusOK, value)
			mutex.Unlock()
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "name not found"})

}

func putStoreByID(c *gin.Context) {
	mutex.Lock()
	id := c.Param("id")
	name := c.Param("name")
	for b := range store {
		if b == id {
			store[id] = name
			c.IndentedJSON(http.StatusOK, store[id])
			mutex.Unlock()
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func main() {
	router := gin.Default()
	router.GET("/store", getStore)
	router.GET("/store/:id", getStoreByID)
	router.PUT("store/:id/:name", putStoreByID)
	router.Run("localhost:9090")
}
