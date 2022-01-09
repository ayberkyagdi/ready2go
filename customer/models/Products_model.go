package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	Product    Product 
	Order 	   int64
	CustomerID int64
	
}
type Product struct {
	gorm.Model
	Name, Description  string
	Price		   float64
	StoreID       	   int64
	Quantity	   int64
}

func (product Product) Open() *gorm.DB {
	Create_db()
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}
	
	return db
}


func (product Product) Migrate() {
	db := product.Open()
	
	db.AutoMigrate(&product)
}

func (product Product) Add() {
	db := product.Open()
	
	db.Create(&product)
}
	
func (product Product) Get(where ...interface{}) Product {
	db := product.Open()

	db.First(&product, where...)
	return product
}

func (product Product) GetAll(where ...interface{}) []Product{
	db := product.Open()

	var products []Product
	db.Find(&products, where...)
	return products
}

func (product Product) Update(column string, value interface{}) {
	db := product.Open()

	db.Model(&product).Update(column,value)
}

func (product Product) Updates(data Product) {
	db := product.Open()

	db.Model(&product).Updates(data)
}

func (product Product) Delete() {
	db := product.Open()

	db.Delete(&product, product.ID)
}
