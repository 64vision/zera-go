package data

import (
	"time"
	u "zerago/utils"
)

type Vendor struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	AddedOn time.Time   `json:"added_on"`
	AddedBy int         `json:"added_by"`
	Profile interface{} `json:"profile"`
	Status  string      `json:"status"`
}

func (entry *Vendor) Insert() map[string]interface{} {
	entry.AddedOn = time.Now()
	errdb := DBM.Insert(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "New vendor has been added")
	return response
}
func (entry *Vendor) Update() map[string]interface{} {

	errdb := DBM.Update(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Vendor has been modified!")
	return response
}
