package config

import (
	admin "Ready2go/admin/controllers"
	customer "Ready2go/customer/controllers"
	"net/http"
	

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router{
	r := httprouter.New()

	// ADMIN 
	r.GET("/admin/stores", admin.Dashboard{}.Index_stores)
	r.GET("/admin/delete/store/:id", admin.Dashboard{}.Delete_store)
	r.GET("/admin/edit/store/:id", admin.Dashboard{}.Edit_store)
	r.POST("/admin/update/store/:id", admin.Dashboard{}.Update_store)
	r.GET("/admin/customers", admin.Dashboard{}.Index_customers)
	r.GET("/admin/delete/customer/:id", admin.Dashboard{}.Delete_customer)
	r.GET("/admin/edit/customer/:id", admin.Dashboard{}.Edit_customer)
	r.POST("/admin/update/customer/:id", admin.Dashboard{}.Update_customer)
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))

	//Customers
	r.GET("/", customer.Userops{}.Login_customer)
	r.GET("/customer/createuser", customer.Userops{}.Register_customer)
	r.GET("/customer/dashboard", customer.Userops{}.Dashboard_customer)
	r.POST("/customer/dashboard", customer.Userops{}.Dashboard_customer)
	r.POST("/customer/registereduser", customer.Userops{}.Db_register_customer)
	r.GET("/customer/buy/:id", customer.Userops{}.Customer_buy_product)
	r.GET("/customer/cart/:id", customer.Userops{}.Customer_add_to_cart)
	r.GET("/cart/list", customer.Userops{}.Customer_list_cart)
	r.GET("/cart/done", customer.Userops{}.Customer_done_cart)
	r.GET("/logout", customer.Userops{}.Logout)

	//Stores
	r.GET("/login/store", customer.Userops{}.Login_store)
	r.GET("/store/createstore", customer.Userops{}.Register_store)
	r.POST("/store/dashboard", customer.Userops{}.Dashboard_store)
	r.GET("/store/dashboard", customer.Userops{}.Dashboard_store)
	r.GET("/store/newproduct/:id", customer.Userops{}.Store_add_product)
	r.POST("/store/registeredstore", customer.Userops{}.Db_register_store)


	return r
}
