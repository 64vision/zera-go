package data

import (
	"encoding/json"
	"time"
	u "zerago/utils"
)

type Customer struct {
	ID        int             `json:"id"`
	CreatedOn time.Time       `json:"created_on"`
	CreatedBy int             `json:"created_by"`
	Status    string          `json:"status"`
	Type      string          `json:"type"`
	Profile   json.RawMessage `json:"profile"`
	Contact   json.RawMessage `json:"contact"`
	Company   json.RawMessage `json:"company"`
	Address   json.RawMessage `json:"address"`
	Files     json.RawMessage `json:"files"`
}

func (entry *Customer) New() map[string]interface{} {
	entry.CreatedOn = time.Now()
	errdb := DBM.Insert(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "New customer has been created!")
	response["product_id"] = entry.ID
	return response
}
func (entry *Customer) Update() map[string]interface{} {
	errdb := DBM.Update(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Customer data has been modified!")
	return response
}
