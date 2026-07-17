package data

import (
	"encoding/json"
	"time"
	u "zerago/utils"
)

type Purchase struct {
	ID          int             `json:"id"`
	Vendor      json.RawMessage `json:"vendor"`
	Company     json.RawMessage `json:"company"`
	Terms       json.RawMessage `json:"terms"`
	Taxes       json.RawMessage `json:"taxes"`
	Credits     json.RawMessage `json:"credits"`
	Items       json.RawMessage `json:"items"`
	Signatories json.RawMessage `json:"signatories"`
	PoDate      string          `json:"po_date"`
	CreatedOn   time.Time       `json:"created_on"`
	CreatedBy   int             `json:"created_by"`
	ModifiedOn  time.Time       `json:"modified_on"`
	Notes       string          `json:"notes"`
	Status      string          `json:"status"`
	Amount      float64         `json:"amount"`
	Logs        json.RawMessage `json:"logs"`
	Discounts   json.RawMessage `json:"discounts"`
}

func (entry *Purchase) New() map[string]interface{} {
	entry.CreatedOn = time.Now()
	errdb := DBM.Insert(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "New Purchase has been requested")
	response["product_id"] = entry.ID
	return response
}
func (entry *Purchase) Update() map[string]interface{} {
	entry.ModifiedOn = time.Now()
	errdb := DBM.Update(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Purchase has been modified!")
	return response
}
