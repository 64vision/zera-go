package data

import (
	"encoding/json"
	"fmt"
	"time"
	u "zerago/utils"
)

type Product struct {
	ID          int             `json:"id" pg:"id"`
	Name        string          `json:"name" pg:"name"`
	Code        string          `json:"code" pg:"code"`
	Description string          `json:"description" pg:"description"`
	AddedDT     time.Time       `json:"added_dt" pg:"added_dt"`
	ModifiedOn  time.Time       `json:"modified_on" pg:"modified_on"`
	Status      string          `json:"status" pg:"status"`
	AddedBy     int             `json:"added_by" pg:"added_by"`
	Type        string          `json:"type" pg:"type"`
	Properties  json.RawMessage `json:"properties" pg:"properties"`
	Tags        json.RawMessage `json:"tags" pg:"tags"`
	Images      json.RawMessage `json:"images" pg:"images"`
}
type ProductItem struct {
	ID           int             `json:"id" pg:"id"`
	ProductID    int             `json:"product_id" pg:"product_id"`
	Stock        int             `json:"stock" pg:"stock"`
	Sold         int             `json:"sold" pg:"sold"`
	AddedDT      time.Time       `json:"added_dt" pg:"added_dt"`
	ModifiedOn   time.Time       `json:"modified_on" pg:"modified_on"`
	Status       string          `json:"status" pg:"status"`
	Cost         float64         `json:"cost" pg:"cost"`
	SellingPrice float64         `json:"selling_price" pg:"selling_price"`
	AddedBy      int             `json:"added_by" pg:"added_by"`
	Unit         string          `json:"unit" pg:"unit"`
	Properties   json.RawMessage `json:"properties" pg:"properties"`
}

func (product *Product) Insert() map[string]interface{} {
	fmt.Println("InsertSales")
	product.AddedDT = time.Now()
	errdb := DBM.Insert(product)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, errdb.Error())
	}

	response := u.Message(true, "New product added!")
	response["product"] = product
	return response

}

func (product *Product) Update() map[string]interface{} {
	fmt.Println("UpdateSales")
	product.ModifiedOn = time.Now()
	errdb := DBM.Update(product)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, errdb.Error())
	}

	response := u.Message(true, "product changes has been saved!")
	response["product"] = product
	return response

}
