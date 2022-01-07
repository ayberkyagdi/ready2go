package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Username,Password,City  string
}

type Store struct {
	gorm.Model
	Storename,Password,City,District  string
}

func (customer Customer) Migrate() {

	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&customer)
}

func (customer Customer) Add() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&customer)
}
	
func (customer Customer) Get(where ...interface{}) Customer {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		panic(err.Error())
	}

	db.First(&customer, where...)
	return customer
}

func (customer Customer) GetAll(where ...interface{}) []Customer{
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var customers []Customer
	db.Find(&customers, where...)
	return customers
}

func (customer Customer) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&customer).Update(column,value)
}

func (customer Customer) Updates(data Customer) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&customer).Updates(data)
}

func (customer Customer) Delete() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&customer, customer.ID)
}

func (store Store) Migrate() {

	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&store)
}

func (store Store) Add() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&store)
}
	
func (store Store) Get(where ...interface{}) Store {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		panic(err.Error())
	}

	db.First(&store, where...)
	return store
}

func (store Store) GetAll(where ...interface{}) []Store{
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var stores []Store
	db.Find(&stores, where...)
	return stores
}

func (store Store) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&store).Update(column,value)
}

func (store Store) Updates(data Store) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&store).Updates(data)
}

func (store Store) Delete() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&store, store.ID)
}