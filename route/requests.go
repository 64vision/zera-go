package main

import (
	"encoding/json"
	"net/http"
	"zerago/account"
	u "zerago/utils"
)

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
