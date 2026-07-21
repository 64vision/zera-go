package data

import (
	"encoding/json"
	"fmt"
	"log"
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
	Remarks          string          `json:"remarks" pg:"remarks"`
	ProductItems     json.RawMessage `json:"product_items" pg:"product_items"`
	ServicePerson    json.RawMessage `json:"service_person" pg:"service_person"`
	ServiceLogs      json.RawMessage `json:"service_logs" pg:"service_logs"`
	DiscountRequest  string          `json:"discount_request" pg:"discount_request"`
	PaymentStatus    string          `json:"payment_status" pg:"payment_status"`
	Files            json.RawMessage `json:"files" pg:"files"`
	PaymentMethod    string          `json:"payment_method" pg:"payment_method"`
	PaymentRef       string          `json:"payment_ref" pg:"payment_ref"`
	DiscountItems    json.RawMessage `json:"discount_items" pg:"discount_items"`
}

type SalesItem struct {
	ID           int64   `json:"id" pg:"id"`
	ProductID    int     `json:"product_id" pg:"product_id"`
	ItemID       int     `json:"item_id" pg:"item_id"`
	Quantity     int     `json:"quantity" pg:"quantity"`
	SalesID      int     `json:"sales_id" pg:"sales_id"`
	SellingPrice float64 `json:"selling_price" pg:"selling_price"`
}
type ItemData struct {
	ID           int             `json:"id"`
	Code         string          `json:"code"`
	ProductID    int             `json:"product_id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Cost         float64         `json:"cost"`
	SellingPrice float64         `json:"selling_price"`
	Amount       float64         `json:"amount"`
	Stock        int             `json:"stock"`
	Sold         *int            `json:"sold"` // nullable
	Quantity     int             `json:"quantity"`
	Unit         string          `json:"unit"`
	Status       string          `json:"status"`
	Checked      bool            `json:"checked"`
	AddedBy      int             `json:"added_by"`
	AddedDT      time.Time       `json:"added_dt"`
	ModifiedOn   time.Time       `json:"modified_on"`
	Tags         []string        `json:"tags"`
	Images       []string        `json:"images"`
	Properties   json.RawMessage `json:"properties"`
	Track        string          `json:"track"`
}

func (sale *Sale) View() map[string]interface{} {
	fmt.Println("Sales View")
	var item Sale
	_, errdb := DBM.QueryOne(&item, "select * from sales where id=?", sale.ID)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, errdb.Error())
	}
	response := u.Message(true, "Ok!")
	response["sale"] = item
	return response

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

func (sale *Sale) Update() map[string]interface{} {
	fmt.Println("UpdateSales")

	sale.UpdatedAt = time.Now()
	errdb := DBM.Update(sale)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, errdb.Error())
	}
	if sale.Status == "Completed" {
		sale.TrackItems()
	}
	response := u.Message(true, "Sales changes has been saved!")
	response["sale"] = sale
	return response

}
func (sale *Sale) TrackItems() {
	var items []ItemData
	var errdb error
	err := json.Unmarshal(sale.ProductItems, &items)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		var sales_item SalesItem

		sales_item.ItemID = item.ID
		sales_item.ProductID = item.ProductID
		sales_item.Quantity = item.Quantity
		sales_item.SalesID = sale.ID
		sales_item.SellingPrice = item.SellingPrice
		errdb = DBM.Insert(&sales_item)
		if errdb != nil {
			panic(errdb)
		}

		if item.Track == "Track Inventory" {
			_, errdb = DBM.Exec("update product_items set stock=stock-?, sold=sold+? where id=?", item.Quantity, item.Quantity, item.ID)
			if errdb != nil {
				panic(errdb)
			}

		} else {
			_, errdb = DBM.Exec("update product_items set sold=sold+? where id=?", item.Quantity, item.ID)
			if errdb != nil {
				panic(errdb)
			}

		}
		fmt.Println(item.Track, item.Track)
	}
}
