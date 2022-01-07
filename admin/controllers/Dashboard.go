package controllers

import (
	"Ready2go/admin/helpers"
	"Ready2go/customer/models"
	"fmt"
	"html/template"
	"net/http"
	"crypto/sha256"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index_stores(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/storelist")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Stores"] = models.Store{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Index_customers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/customerlist")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Customers"] = models.Customer{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}


func (dashboard Dashboard) Delete_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	store := models.Store{}.Get(params.ByName("id"))
	store.Delete()
	http.Redirect(w,r,"/admin/stores",http.StatusSeeOther)
}

func (dashboard Dashboard) Edit_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/storeedit")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Store"] = models.Store{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	store := models.Store{}.Get("id = ?", params.ByName("id"))
	storename := r.FormValue("storename")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	city := r.FormValue("city")
	district := r.FormValue("district")
	store.Updates(models.Store{
		Storename: storename,
		Password: password,
		City: city,
		District: district,
	})
	http.Redirect(w, r, "/admin/stores", http.StatusSeeOther)

}

func (dashboard Dashboard) Delete_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customer := models.Customer{}.Get(params.ByName("id"))
	customer.Delete()
	http.Redirect(w,r,"/admin/customers",http.StatusSeeOther)
}

func (dashboard Dashboard) Edit_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/customeredit")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Customer"] = models.Customer{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	customer := models.Customer{}.Get("id = ?", params.ByName("id"))
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	city := r.FormValue("city")
	customer.Updates(models.Customer{
		Username: username,
		Password: password,
		City: city,
	})
	http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)

}