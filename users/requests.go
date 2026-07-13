package users

import (
	"encoding/json"
	"net/http"
	u "zerago/utils"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	admin := &Administrator{}
	err := json.NewDecoder(r.Body).Decode(admin) //decode the request body into struct and failed if any error occur
	if err != nil {
		//panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := admin.Login()
	u.Respond(w, resp)
}
