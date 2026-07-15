package data

import (
	"encoding/json"
	"time"
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
