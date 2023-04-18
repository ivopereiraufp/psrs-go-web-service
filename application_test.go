package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/albums", getAlbums)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Blue Train")
	assert.Contains(t, w.Body.String(), "Jeru")
	assert.Contains(t, w.Body.String(), "Sarah Vaughan and Clifford Brown")
}

func TestGetAlbumByID_ValidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/albums/:id", getAlbumByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Blue Train")
}

func TestGetAlbumByID_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/albums/:id", getAlbumByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/nonexistent", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "album not found")
}

func TestPostAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/albums", postAlbums)

	newAlbum := `{
		"ID": "4",
		"Title": "Album 4",
		"Artist": "Artist 4",
		"Price": 40.0
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(newAlbum))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Album 4")
}
