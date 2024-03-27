package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// assign handler function to an endpoint path
func main() {
	// init a gin router using Default
	router := gin.Default()
	// associate the GET HTTP method and /albums path with a handler function
	router.GET("/albums", getAlbums)
	// GET album by ID
	router.GET("albums/:id", getAlbumByID)
	// POST to add new album
	router.POST("/albums", postAlbums)

	// run the function to attach the router to an http.Server and start server
	router.Run("0.0.0.0:8080")
}

// write handler to return all items
// prepare a response
// getAlbums responds with a list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add the new album to the slice.
	albums = append(albums, newAlbum)
	// add 201 status code to response, along with JSON representing the album added
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for album that matches ID value
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}
