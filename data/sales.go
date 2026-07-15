package data

import (
	"encoding/json"
	"fmt"
	"time"
	u "zerago/utils"
)

type Sale struct {
	ID               int             `json:"id" pg:"id"`
	InvoiceNo        string          `json:"invoice_no" pg:"invoice_no"`
	CustomerID       int             `json:"customer_id" pg:"customer_id"`
	CustomerData     json.RawMessage `json:"customer_data" pg:"customer_data"`
	TotalAmount      float64         `json:"total_amount" pg:"total_amount"`
	Discount         float64         `json:"discount" pg:"discount"`
	SalesAmount      float64         `json:"sales_amount" pg:"sales_amount"`
	Tax              float64         `json:"tax" pg:"tax"`
	Type             string          `json:"type" pg:"type"`
	InitPayment      float64         `json:"init_payment" pg:"init_payment"`
	PaymentTerms     string          `json:"payment_terms" pg:"payment_terms"`
	PaymentFrequency string          `json:"payment_frequency" pg:"payment_frequency"`
	PaymentLength    int             `json:"payment_length" pg:"payment_length"`
	DueAmount        float64         `json:"due_amount" pg:"due_amount"`
	DueDate          string          `json:"due_date" pg:"due_date"`
	SalesPerson      int             `json:"sales_person" pg:"sales_person"`
	ReferralPerson   int             `json:"referral_person" pg:"referral_person"`
	CreatedAt        time.Time       `json:"created_at" pg:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" pg:"updated_at"`
	Status           string          `json:"status" pg:"status"`
	Remarks          string          `json:"remarks" pg:"Remarks"`
	ProductItems     json.RawMessage `json:"product_items" pg:"product_items"`
	ServicePerson    json.RawMessage `json:"service_person" pg:"service_person"`
	ServiceLogs      json.RawMessage `json:"service_logs" pg:"service_logs"`
}

func (sale *Sale) Insert() map[string]interface{} {
	fmt.Println("InsertSales")
	sale.CreatedAt = time.Now()
	errdb := DBM.Insert(sale)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, errdb.Error())
	}

	response := u.Message(true, "Saved!")
	response["sale"] = sale
	return response

}
