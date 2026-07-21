package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zerago/data"
	u "zerago/utils"
)

func NewExpenseAccount(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	entry := &data.ExpensesAccount{}
	err := json.NewDecoder(r.Body).Decode(entry) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	entry.CreatedBy = r.Context().Value("user").(int)
	fmt.Println("entry.CreatedBy", entry.CreatedBy)
	resp := entry.Insert()
	u.Respond(w, resp)
}

func UpdateExpenseAccount(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	entry := &data.ExpensesAccount{}
	err := json.NewDecoder(r.Body).Decode(entry) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := entry.Update()
	u.Respond(w, resp)
}

func NewExpense(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	entry := &data.Expense{}
	err := json.NewDecoder(r.Body).Decode(entry) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	entry.CreatedBy = r.Context().Value("user").(int)
	resp := entry.Insert()
	u.Respond(w, resp)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	entry := &data.Expense{}
	err := json.NewDecoder(r.Body).Decode(entry) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := entry.Update()
	u.Respond(w, resp)
}
