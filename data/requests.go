package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	u "zerago/utils"
)

func QueryData(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("QueryData")
	qry := &Query{}
	var resp map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(qry) //decode the request body into struct and failed if any error occur
	if err != nil {
		panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	for _, prohibited := range Prohibiteds {
		if strings.Contains(strings.ToLower(qry.Query), strings.ToLower(prohibited)) {
			u.Respond(w, u.Message(false, "Violated the data ruling!"))
			return
		}
	}
	resp = qry.ExecQuery()

	u.Respond(w, resp)
}

func InsertData(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	qry := &Query{}
	var resp map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(qry) //decode the request body into struct and failed if any error occur
	if err != nil {
		//panic(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	for _, prohibited := range Prohibiteds {
		if strings.Contains(strings.ToLower(qry.Query), strings.ToLower(prohibited)) {
			u.Respond(w, u.Message(false, "Violated the data ruling!"))
			return
		}
	}
	resp = qry.ExecInsert()
	u.Respond(w, resp)
}
