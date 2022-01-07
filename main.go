package main

import (
	"Ready2go/config"
	"net/http"
	"gorm.io/gorm"
	user_models "Ready2go/customer/models"
)

type User struct {
	gorm.Model
	Username, Password string
}

func main() {
	user_models.Product{}.Migrate()
	user_models.Customer{}.Migrate()
	user_models.Store{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())

}
