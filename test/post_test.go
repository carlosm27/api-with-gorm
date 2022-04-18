package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carlosm27/go_projects/apiwithgorm/controllers"
	"github.com/carlosm27/go_projects/apiwithgorm/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostGrocery(t *testing.T) {

	model.Database()
	router := gin.Default()

	values := map[string]interface{}{"name": "Banana", "quantity": 45}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/grocery", bytes.NewBuffer(json_data))
	router.POST("/grocery", controllers.PostGrocery)

	var res map[string]interface{}

	w := httptest.NewRecorder()

	json.NewDecoder(w.Body).Decode(&res)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, GetGroceryTest("1"), w.Body.String())
	fmt.Println(GetGroceryTest("1"))
	fmt.Println(w.Body.String())

}
func TestPostInvalidGrocery(t *testing.T) {
	model.Database()
	router := gin.Default()

	values := map[string]interface{}{"badname": "Banana", "badquantity": 45}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/grocery", bytes.NewBuffer(json_data))
	router.POST("/grocery", controllers.PostGrocery)

	var res map[string]interface{}

	w := httptest.NewRecorder()

	json.NewDecoder(w.Body).Decode(&res)

	router.ServeHTTP(w, req)
	expected := `{"error":"Key: 'NewGrocery.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'NewGrocery.Quantity' Error:Field validation for 'Quantity' failed on the 'required' tag"}`
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func GetGroceryTest(id string) string {
	model.Database()
	router := gin.Default()
	router.GET("/grocery/:id", controllers.GetGrocery)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/grocery/"+id, nil)
	router.ServeHTTP(w, req)

	return w.Body.String()
}
