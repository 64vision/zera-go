package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"zerago/auth"
	"zerago/data"
)

const (
	PORT = "8100"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/app/user/login", UserLogin).Methods("POST")
	router.HandleFunc("/app/new_user", NewUser).Methods("POST")
	router.HandleFunc("/app/update_user", UpdateUser).Methods("POST")
	router.HandleFunc("/app/update_password", UpdatePassword).Methods("POST")

	router.HandleFunc("/app/query", data.QueryData).Methods("POST")
	router.HandleFunc("/app/new_data", data.InsertData).Methods("POST")
	router.HandleFunc("/app/add_sales", AddSales).Methods("POST")
	router.HandleFunc("/app/update_sales", UpdateSales).Methods("POST")
	router.HandleFunc("/app/new_product", NewProduct).Methods("POST")
	router.HandleFunc("/app/update_product", UpdateProduct).Methods("POST")
	router.HandleFunc("/app/new_vendor", NewVendor).Methods("POST")
	router.HandleFunc("/app/update_vendor", UpdateVendor).Methods("POST")
	router.HandleFunc("/app/new_purchase", NewPurchase).Methods("POST")
	router.HandleFunc("/app/update_purchase", UpdatePurchase).Methods("POST")
	router.HandleFunc("/app/new_customer", NewCustomer).Methods("POST")
	router.HandleFunc("/app/update_customer", UpdateCustomer).Methods("POST")
	router.HandleFunc("/app/view_sales", SalesView).Methods("POST")

	//expenses
	router.HandleFunc("/app/new_expense", NewExpense).Methods("POST")
	router.HandleFunc("/app/update_expense", UpdateExpense).Methods("POST")
	router.HandleFunc("/app/new_expenses_account", NewExpenseAccount).Methods("POST")
	router.HandleFunc("/app/update_expenses_account", UpdateExpenseAccount).Methods("POST")

	router.Use(auth.JwtAuthentication)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	handler := cors.Default().Handler(router)
	handler = c.Handler(handler)
	rand.Seed(time.Now().UnixNano())

	/*--------------------------------------------------
		Run Server
	-----------------------------------------------------*/
	fmt.Println("HYPERBALL server run at port: " + PORT)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
