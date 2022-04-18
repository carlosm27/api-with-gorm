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

func TestGetGroceries(t *testing.T) {

	model.Database()
	router := gin.Default()
	router.GET("/groceries", controllers.GetGroceries)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/groceries", nil)
	router.ServeHTTP(w, req)

	expected := `[{"ID":1,"name":"Banana","quantity":45,"CreatedAt":"2022-04-17T22:12:00.1380124-04:00","UpdatedAt":"2022-04-17T22:12:00.1380124-04:00","DeletedAt":null}]`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
