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
	// Routes consist of a path and a handler function.
	// router.HandleFunc("/account/register", users.Register).Methods("POST")
	// router.HandleFunc("/account/verify", users.Verify).Methods("POST")
	router.HandleFunc("/app/user/login", UserLogin).Methods("POST")
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

	// router.HandleFunc("/account/get", account.GetAccount).Methods("POST")
	// router.HandleFunc("/account/resendcode", account.ResendCode).Methods("POST")
	// router.HandleFunc("/account/forgot", account.ForgotPassword).Methods("POST")
	// router.HandleFunc("/account/balance", account.BalanceInquire).Methods("POST")
	// router.HandleFunc("/account/qry", account.CustomQry).Methods("POST")
	// router.HandleFunc("/account/gateways", account.GetGatewayEnabled).Methods("GET")
	// router.HandleFunc("/credits/buy", account.DoBuyCredits).Methods("POST")
	// router.HandleFunc("/credits/cashout", account.DoCashout).Methods("POST")
	// router.HandleFunc("/credits/cashout_cancel", account.UpdateCashout).Methods("POST")
	// router.HandleFunc("/credits/get_cashout", account.GetCashout).Methods("POST")
	// router.HandleFunc("/credits/callback", account.Callback).Methods("POST")
	// router.HandleFunc("/credits/maya_callback", account.MayaCallback).Methods("POST")

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
