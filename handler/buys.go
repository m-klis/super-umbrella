package handler

import (
	"gochicoba/service"
	"net/http"
)

type BuyHandler struct {
	buyService service.BuyService
}

func NewBuyHandler(buyService service.BuyService) BuyHandler {
	return BuyHandler{buyService: buyService}
}

func (ih *BuyHandler) GetAllBuys(w http.ResponseWriter, r *http.Request) {
	return
}

func (ih *BuyHandler) CreateBuy(w http.ResponseWriter, r *http.Request) {
	return
}
