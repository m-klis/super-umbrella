package models

import (
	"time"
)

// respon get /buy/
type BuyResponse struct {
	ID          int     `json:"id"`
	IdUser      int     `json:"id_user"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
	CreatedAt   string  `json:"created_at"`
}

type TransactionResponse struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	ItemId    int       `json:"item_id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Uuid      string    `json:"uuid"`
}

type BuyOrder struct {
	ID        int         `json:"id"`
	IdUser    int         `json:"id_user" validation:"required"`
	ListOrder []ListOrder `json:"list_order"`
}

type ListOrder struct {
	IdItem      int     `json:"id_item"`
	ItemName    string  `json:"item_name"`
	ItemPrice   float64 `json:"item_price"`
	ItemAmount  int     `json:"item_amount"`
	PriceAmount float64 `json:"price_amount"`
}
