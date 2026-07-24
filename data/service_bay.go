package data

import (
	u "zerago/utils"
)

type ServiceBay struct {
	ID          int    `json:"id" pg:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (entry *ServiceBay) Insert() map[string]interface{} {
	errdb := DBM.Insert(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Expense Account has been added")
	return response
}
func (entry *ServiceBay) Update() map[string]interface{} {

	errdb := DBM.Update(entry)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Expense Account has been modified!")
	return response
}
