package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carlosm27/apiwithgorm/grocery"
	"github.com/carlosm27/apiwithgorm/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateGrocery(t *testing.T) {

	model.Database()
	router := gin.Default()

	values := map[string]interface{}{"name": "Banana", "quantity": 45}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/grocery/1", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, GetGroceryUpdatedTest("1"), w.Body.String())
}

func TestUpdateGroceryNotFound(t *testing.T) {
	model.Database()
	router := gin.Default()
	values := map[string]interface{}{"name": "Banana", "quantity": 45}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/grocery/20", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Grocery not found!"}`, w.Body.String())
}
func TestInvalidField(t *testing.T) {
	model.Database()
	router := gin.Default()
	values := map[string]interface{}{"badname": "Banana", "badquantity": 45}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/grocery/16", bytes.NewBuffer(json_data))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, GetGroceryUpdatedTest("16"), w.Body.String())
}

func GetGroceryUpdatedTest(id string) string {
	model.Database()
	router := gin.Default()
	router.GET("/grocery/:id", grocery.GetGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/grocery/"+id, nil)
	router.ServeHTTP(w, req)

	return w.Body.String()
}
