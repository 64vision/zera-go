package data

import (
	"encoding/json"
	"fmt"
	u "zerago/utils"
)

type Query struct {
	Table string `json:"table"`
	Query string `json:"query"`
	Type  string `json:"type"`
}
type Item struct {
	Data json.RawMessage `json:"data"`
}

func (qry *Query) ExecQuery() map[string]interface{} {
	fmt.Println("ExecQuery", qry.Query)
	response := u.Message(true, "Result")
	var results Item
	_, err := DBM.Query(&results, `SELECT json_agg(t) as data
							FROM (
								`+qry.Query+`
							) t;`)
	if err != nil {
		//panic(err)
		return u.Message(false, err.Error())
	}
	if len(results.Data) == 0 {
		return u.Message(false, "No rows found!")
	}
	response = u.Message(true, "Ok!")
	response["results"] = results.Data
	return response

}

func (qry *Query) ExecInsert() map[string]interface{} {
	fmt.Println("ExecInsert", qry.Query)
	res, errdb := DBM.Exec(qry.Query)
	if errdb != nil {
		//panic(errdb)
		return u.Message(false, errdb.Error())
	}
	fmt.Println(res.Model())
	if res.RowsAffected() == 0 {
		return u.Message(false, "No rows affected!")
	}
	response := u.Message(true, "Save")
	response[qry.Table] = res
	return response

}
