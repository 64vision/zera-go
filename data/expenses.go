package data

import (
	"encoding/json"
	"time"
	u "zerago/utils"
)

type Expense struct {
	ID            int             `json:"id" pg:"id"`
	Date          string          `json:"date" pg:"date"`
	Category      string          `json:"category" pg:"category"`
	PaymentMethod string          `json:"payment_method" pg:"payment_method"`
	Amount        float64         `json:"amount" pg:"amount"`
	PayeeTo       string          `json:"payee_to" pg:"payee_to"`
	RefNo         string          `json:"ref_no" pg:"ref_no"`
	Description   string          `json:"description" pg:"description"`
	CreatedBy     int             `json:"created_by" pg:"created_by"`
	CreatedAt     time.Time       `json:"created_at" pg:"created_at"`
	Status        string          `json:"status" pg:"status"`
	AccountName   string          `json:"account_name" pg:"account_name"`
	AccountID     int             `json:"account_id" pg:"account_id"`
	Attachments   json.RawMessage `json:"attachments" pg:"attachments"`
}

type ExpensesAccount struct {
	ID          int       `json:"id" pg:"id"`
	Name        string    `json:"name" pg:"name"`
	Category    string    `json:"category" pg:"category"`
	Description string    `json:"description" pg:"description"`
	Status      string    `json:"status" pg:"status"`
	CreatedBy   int       `json:"created_by" pg:"created_by"`
	CreatedAt   time.Time `json:"created_at" pg:"created_at"`
}

func (entry *ExpensesAccount) Insert() map[string]interface{} {
	entry.CreatedAt = time.Now()
	errdb := DBM.Insert(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Expense Account has been added")
	return response
}
func (entry *ExpensesAccount) Update() map[string]interface{} {

	errdb := DBM.Update(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Expense Account has been modified!")
	return response
}

func (entry *Expense) Insert() map[string]interface{} {
	entry.CreatedAt = time.Now()
	errdb := DBM.Insert(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "New Expense has been added")
	return response
}
func (entry *Expense) Update() map[string]interface{} {

	errdb := DBM.Update(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Expense has been modified!")
	return response
}
