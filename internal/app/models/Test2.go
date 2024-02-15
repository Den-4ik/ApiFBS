package models

type Test2Stoks struct {
	Stocks []Test2SKU `json:"stocks"`
}

type Test2SKU struct {
	Sku    string `json:"sku"`
	Amount int    `json:"amount"`
}
