package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	Product Product 
	Order int64
	CustomerID int64
	
}
type Product struct {
	gorm.Model
	Name, Description  string
	Price		       float64
	StoreID       	   int64
	Quantity		   int64
}

func (product Product) Migrate() {

	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&product)
}

func (product Product) Add() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&product)
}
	
func (product Product) Get(where ...interface{}) Product {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		panic(err.Error())
	}

	db.First(&product, where...)
	return product
}

func (product Product) GetAll(where ...interface{}) []Product{
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var products []Product
	db.Find(&products, where...)
	return products
}

func (product Product) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&product).Update(column,value)
}

func (product Product) Updates(data Product) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&product).Updates(data)
}

func (product Product) Delete() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&product, product.ID)
}