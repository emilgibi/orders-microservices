package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emilgibi/orders-microservices/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (handler *Handler) Connect(host, user, pass, dbName, port string) {
	var err error
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	handler.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	handler.DB.AutoMigrate(models.Order{})
	if err != nil {
		panic(err)
	}

}

func (handler *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	handler.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func (handler *Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var addOrder models.Order
	_ = json.NewDecoder(r.Body).Decode(&addOrder)
	handler.DB.Create(&addOrder)
	json.NewEncoder(w).Encode(addOrder)
}
