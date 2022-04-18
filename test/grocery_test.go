package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carlosm27/go_projects/apiwithgorm/controllers"
	"github.com/carlosm27/go_projects/apiwithgorm/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGrocery(t *testing.T) {

	model.Database()
	router := gin.Default()
	router.GET("/grocery/:id", controllers.GetGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/grocery/1", nil)
	router.ServeHTTP(w, req)

	expected := `{"ID":1,"name":"Banana","quantity":45,"CreatedAt":"2022-04-17T22:12:00.1380124-04:00","UpdatedAt":"2022-04-17T22:12:00.1380124-04:00","DeletedAt":null}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestGetGroceryNotFound(t *testing.T) {
	model.Database()
	router := gin.Default()
	router.GET("/grocery/:id", controllers.GetGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/grocery/20", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Grocery not found"}`, w.Body.String())

}
