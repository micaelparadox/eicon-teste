package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"test-eicon/controllers"
	"test-eicon/models"
	"test-eicon/repositories"
	"test-eicon/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Order{})

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	r := gin.Default()
	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders", orderController.GetOrders)

	return r
}

func TestCreateAndGetOrders(t *testing.T) {
	router := setupRouter()

	// Test creating orders (JSON)
	orders := []map[string]interface{}{
		{
			"control_number": "111111",
			"name":           "Produto A",
			"unit_price":     100.0,
			"quantity":       5,
			"customer_code":  1,
		},
		{
			"control_number": "222222",
			"name":           "Produto B",
			"unit_price":     200.0,
			"quantity":       10,
			"customer_code":  2,
		},
	}
	jsonValue, _ := json.Marshal(orders)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Test getting all orders
	req, _ = http.NewRequest("GET", "/orders", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, "111111", response[0]["control_number"])
	assert.Equal(t, "Produto A", response[0]["name"])
	assert.Equal(t, 500.0, response[0]["total_value"])
	assert.Equal(t, "222222", response[1]["control_number"])
	assert.Equal(t, "Produto B", response[1]["name"])
	assert.Equal(t, 1800.0, response[1]["total_value"]) // Verificar o desconto aplicado

	// Test getting order by control number
	req, _ = http.NewRequest("GET", "/orders?control_number=111111", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseByControlNumber []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &responseByControlNumber)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(responseByControlNumber))
	assert.Equal(t, "111111", responseByControlNumber[0]["control_number"])
	assert.Equal(t, "Produto A", responseByControlNumber[0]["name"])
	assert.Equal(t, 500.0, responseByControlNumber[0]["total_value"])

	// Test getting orders by date
	currentDate := time.Now().Format("2006-01-02")
	req, _ = http.NewRequest("GET", "/orders?date="+currentDate, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseByDate []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &responseByDate)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(responseByDate))
	assert.Equal(t, "111111", responseByDate[0]["control_number"])
	assert.Equal(t, "Produto A", responseByDate[0]["name"])
	assert.Equal(t, "222222", responseByDate[1]["control_number"])
	assert.Equal(t, "Produto B", responseByDate[1]["name"])

	// Test creating orders (XML)
	ordersXML := `<orders>
    <order>
        <control_number>333333</control_number>
        <name>Produto C</name>
        <unit_price>300.0</unit_price>
        <quantity>3</quantity>
        <customer_code>3</customer_code>
    </order>
    <order>
        <control_number>444444</control_number>
        <name>Produto D</name>
        <unit_price>400.0</unit_price>
        <quantity>7</quantity>
        <customer_code>4</customer_code>
    </order>
</orders>`
	req, _ = http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte(ordersXML)))
	req.Header.Set("Content-Type", "application/xml")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Test getting all orders again
	req, _ = http.NewRequest("GET", "/orders", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(response)) // Verificar o tamanho da resposta
	assert.Equal(t, "111111", response[0]["control_number"])
	assert.Equal(t, "Produto A", response[0]["name"])
	assert.Equal(t, "222222", response[1]["control_number"])
	assert.Equal(t, "Produto B", response[1]["name"])
	assert.Equal(t, "333333", response[2]["control_number"])
	assert.Equal(t, "Produto C", response[2]["name"])
	assert.Equal(t, "444444", response[3]["control_number"])
	assert.Equal(t, "Produto D", response[3]["name"])
}
