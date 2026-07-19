package main

import (
	"encoding/json"
	"net/http"
	"zerago/account"
	"zerago/data"
	u "zerago/utils"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &account.User{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err.Error())
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.ChangePassword()
	u.Respond(w, resp)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &account.User{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := item.Create()
	u.Respond(w, resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &account.User{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.Update()
	u.Respond(w, resp)
}

func SalesView(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Sale{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.View()
	u.Respond(w, resp)
}

func NewCustomer(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Customer{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	item.CreatedBy = r.Context().Value("user").(int)
	resp := item.New()
	u.Respond(w, resp)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Customer{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.Update()
	u.Respond(w, resp)
}

func NewPurchase(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Purchase{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	item.CreatedBy = r.Context().Value("user").(int)
	resp := item.New()
	u.Respond(w, resp)
}

func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Purchase{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.Update()
	u.Respond(w, resp)
}

func UpdateVendor(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Vendor{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := item.Update()
	u.Respond(w, resp)
}

func NewVendor(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	item := &data.Vendor{}
	err := json.NewDecoder(r.Body).Decode(item) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	item.AddedBy = r.Context().Value("user").(int)
	resp := item.Insert()
	u.Respond(w, resp)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	product := &data.Product{}
	err := json.NewDecoder(r.Body).Decode(product) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := product.Update()
	u.Respond(w, resp)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	product := &data.Product{}
	err := json.NewDecoder(r.Body).Decode(product) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	product.AddedBy = r.Context().Value("user").(int)
	resp := product.Insert()
	u.Respond(w, resp)
}

func UpdateSales(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	sales := &data.Sale{}
	err := json.NewDecoder(r.Body).Decode(sales) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := sales.Update()
	u.Respond(w, resp)
}

func AddSales(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	sales := &data.Sale{}
	err := json.NewDecoder(r.Body).Decode(sales) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	sales.SalesPerson = r.Context().Value("user").(int)
	resp := sales.Insert()
	u.Respond(w, resp)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	user := &account.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := user.Login()
	u.Respond(w, resp)
}
