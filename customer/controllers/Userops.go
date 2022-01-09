package controllers

import (
	"Ready2go/customer/helpers"
	"Ready2go/customer/models"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

type Userops struct{}

var store = sessions.NewCookieStore([]byte("mysession"))

func GetSessionLoggedID(r *http.Request) int {
	storeAuth, _ := store.Get(r, "authentication")
	if auth, ok := storeAuth.Values["loggedID"]; ok {
		return auth.(int)
	}
	fmt.Println("none found")
	return 0
}

func SetSessionLoggedID(w http.ResponseWriter, r *http.Request, id int) {
	storeAuth, err := store.Get(r, "authentication")
	if err != nil {
		fmt.Println(err.Error())
	}
	storeAuth.Options = &sessions.Options{HttpOnly: true, Secure: true, MaxAge: 2628000}
	storeAuth.Values["loggedID"] = id
	storeAuth.Save(r, w)
}

func (userops Userops) Login_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("customer")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "login", nil)

}

func (userops Userops) Login_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("store")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "login", nil)
}
func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session, _ := store.Get(r, "mysession")
    session.Options.MaxAge = -1
	session.Values = nil
	session.Save(r, w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (userops Userops) Register_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("customer")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "register", nil)
}

func (userops Userops) Register_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("store")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "register", nil)
}

func (userops Userops) Dashboard_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := make(map[string]interface{})
	if r.Method =="POST"{
		username := r.FormValue("username")
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
		admin_pass := fmt.Sprintf("%x", sha256.Sum256([]byte("admin")))

		if username == "admin" && password == admin_pass {
			http.Redirect(w, r, "/admin/customers", http.StatusSeeOther)
			
		}else{
		user := models.Customer{}.Get("username = ? AND password = ?", username, password)
		if user.Username == username && user.Password == password {

			view, _ := template.ParseFiles(helpers.Include("customer")...)
			data["Stores"] = models.Store{}.GetAll("city = ?", user.City)
			SetSessionLoggedID(w, r, int(user.ID))
			view.ExecuteTemplate(w, "dashboard", data)
		}else{
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
		}
		
	}else{
		view, _ := template.ParseFiles(helpers.Include("customer")...)
		customer_id := GetSessionLoggedID(r)
		user := models.Customer{}.Get("id  = ?", customer_id)
		data["Stores"] = models.Store{}.GetAll("city = ?", user.City)
		view.ExecuteTemplate(w, "dashboard", data)

	}	

	
}

func (userops Userops) Dashboard_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("store")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})

	if r.Method == "POST" && len(r.FormValue("storename")) > 0 {
		username := r.FormValue("storename")
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
		user := models.Store{}.Get("storename = ? AND password = ?", username, password)
		SetSessionLoggedID(w, r, int(user.ID))
		store_id := GetSessionLoggedID(r)
		data["Products"] = models.Product{}.GetAll("store_id = ?",store_id)
		data["ID"] = store_id
		view.ExecuteTemplate(w, "dashboard", data)
	} else if r.Method == "POST" && len(r.FormValue("name")) > 0 {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price,_ := strconv.ParseFloat(r.FormValue("price"),64)
		storeid := GetSessionLoggedID(r)
		quantity,_ := strconv.Atoi(r.FormValue("quantity"))
		models.Product{
			Name:        name,
			Description: description,
			Price:       float64(price),
			StoreID:     int64(storeid),
			Quantity:    int64(quantity),
		}.Add()
		store_id := GetSessionLoggedID(r)
		data["Products"] = models.Product{}.GetAll("store_id = ?",store_id)
		data["ID"] = store_id
		view.ExecuteTemplate(w, "dashboard", data)
	} else {
		store_id := GetSessionLoggedID(r)
		data["Products"] = models.Product{}.GetAll("store_id = ?",store_id)
		data["ID"] = store_id
		view.ExecuteTemplate(w, "dashboard", data)
	}
}

func (userops Userops) Store_add_product(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("store")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["ID"] = params.ByName("id")
	view.ExecuteTemplate(w, "newproduct", data)
}

func (userops Userops) Customer_buy_product(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("customer")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	store_id := params.ByName("id")
	data["Products"] = models.Product{}.GetAll("store_id = ?",store_id)
	
	view.ExecuteTemplate(w, "buyproduct", data)
}

func (userops Userops) Customer_add_to_cart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	product_id := params.ByName("id")
	product := models.Product{}.Get(product_id)
	session, _ := store.Get(r, "mysession")
	customer_id := GetSessionLoggedID(r)
	cart := session.Values["cart"]
	if cart == nil{
		var cart []models.Item
		cart = append(cart, models.Item{
			Product: product,
			Order: 1,
			CustomerID:int64(customer_id),
		} )
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
		
		
	}else {
		strCart := session.Values["cart"].(string)
		var cart []models.Item
		json.Unmarshal([]byte(strCart), &cart)
		int_id,_ := strconv.Atoi(product_id)
		index := item_exist(int_id, cart)

		if index == -1 {
			cart = append(cart, models.Item{
				Product: product,
				Order: 1,
			} )
		}else{
			product := models.Product{}.Get("id = ?", int_id)
			if cart[index].Order+1 <= product.Quantity{
				cart[index].Order++
			}
		}
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
	
	}
	session.Save(r,w)
	http.Redirect(w, r, "/cart/list", http.StatusSeeOther)
	
}

func item_exist(id int, cart []models.Item) int {
	for i:= 0; i<len(cart); i++{
		if int(cart[i].Product.ID) == id {
			return i
		}
	}
	return -1
}


func (userops Userops) Customer_list_cart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var cart []models.Item
	view, _ := template.ParseFiles(helpers.Include("customer")...)
	session, _ := store.Get(r, "mysession")
	strCart := session.Values["cart"].(string)

	json.Unmarshal([]byte(strCart), &cart)
	var Totalprice float64 = 0
	for i := range cart {
		Totalprice += float64(cart[i].Order) * cart[i].Product.Price
	}
	data := map[string]interface{}{
		"cart":cart,
		"total":Totalprice,
		"storeID":cart[len(cart)-1].Product.StoreID,
	}
	
	view.ExecuteTemplate(w,"cart",data)

}

func (userops Userops) Db_register_customer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	city := r.FormValue("city")
	models.Customer{
		Username: username,
		Password: password,
		City:     city,
	}.Add()
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (userops Userops) Db_register_store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	storename := r.FormValue("storename")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	city := r.FormValue("city")
	district := r.FormValue("district")
	models.Store{
		Storename: storename,
		Password:  password,
		City:      city,
		District:  district,
	}.Add()
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (userops Userops) Product_adding(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price,_ := strconv.Atoi(r.FormValue("price"))
	storeid := GetSessionLoggedID(r)
	quantity,_ := strconv.Atoi(r.FormValue("quantity"))
	models.Product{
		Name:        name,
		Description: description,
		Price:       float64(price),
		StoreID:     int64(storeid),
		Quantity:    int64(quantity),
	}.Add()
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (userops Userops) Customer_done_cart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var cart []models.Item
	var customer models.Customer
    type Dictionary map[string]interface{}
	d1 := []Dictionary{}
	view, _ := template.ParseFiles(helpers.Include("customer")...)
	session, _ := store.Get(r, "mysession")
	strCart := session.Values["cart"].(string)
	
	json.Unmarshal([]byte(strCart), &cart)
	var totalcost float64 = 0
	order_time := time.Now().Format("Mon Jan _2 15:04:05 2006")
	for i := range cart {
		store := models.Store{}.Get("id = ?", cart[i].Product.StoreID)
		cost := strconv.Itoa(int(cart[i].Product.Price))+"₺"
		totalcost += float64(cart[i].Order) * cart[i].Product.Price
		customer = models.Customer{}.Get("id = ?", cart[0].CustomerID)
		d2 := Dictionary{"Customer":customer.Username,"Customer_City":customer.City,"Product":cart[i].Product.Name,"Price":cost,"Quantity":strconv.Itoa(int(cart[i].Order)),"Store":store.Storename,
						"Store_City":store.City,"Store_District":store.District,"Order_time":order_time,"Total_Cost":totalcost}
		product := models.Product{}.Get("name = ? AND store_id = ?", cart[i].Product.Name,cart[i].Product.StoreID)
		product.Update("quantity", (cart[i].Product.Quantity - cart[i].Order))
		d1 = append(d1,d2)
	}

	file, _ := json.MarshalIndent(d1, "", " ")
	now := time.Now()
	file_time := now.Format(time.RFC3339Nano)
	file_time = file_time[:len(file_time)-14]
	file_time = strings.Replace(file_time, ":", "_", -1 )

	path := "orders/"
	filename := path+customer.Username+"-"+file_time+".json"
	
	_ = ioutil.WriteFile(filename, file, 0644)
	session.Values["cart"] = nil
	session.Save(r, w)
	view.ExecuteTemplate(w,"cart",nil)

}