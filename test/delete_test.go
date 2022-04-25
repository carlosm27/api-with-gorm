package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carlosm27/apiwithgorm/grocery"
	"github.com/carlosm27/apiwithgorm/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteGrocery(t *testing.T) {

	model.Database()
	router := gin.Default()
	router.DELETE("/grocery/:id", grocery.DeleteGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/grocery/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"Grocery deleted"}`, w.Body.String())
}

func TestDeleteGroceryNotFound(t *testing.T) {
	model.Database()
	router := gin.Default()
	router.DELETE("/grocery/:id", grocery.DeleteGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/grocery/20", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Grocery not found!"}`, w.Body.String())
}
