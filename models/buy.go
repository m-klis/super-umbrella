package models

import (
	"time"
)

type ResponseGet struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       []ResponseGetData `json:"data"`
}

type ResponseGetSingle struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       ResponseGetData `json:"data"`
}

type ResponseGetData struct {
	ID          int     `json:"id"`
	IdUser      int     `json:"id_user"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
	CreatedAt   string  `json:"create_at"`
}

type BuyResponse struct {
	ID          int     `json:"id"`
	IdUser      int     `json:"id_user"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
	CreatedAt   string  `json:"created_at"`
}

type Transaction struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	ItemId    int       `json:"item_id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Uuid      string    `json:"uuid"`
}

type DataBuy struct {
	IdUser      int     `json:"id_user"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
}

type BuyDetail struct {
	IdItem      int     `json:"id_item"`
	ItemPrice   float64 `json:"item_price"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
}

type CreateBuy struct {
	DataBuy  DataBuy     `json:"data_buy"`
	DataItem []BuyDetail `json:"data_item"`
}
