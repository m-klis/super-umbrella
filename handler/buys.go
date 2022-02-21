package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gochicoba/helpers"
	"gochicoba/models"
	"gochicoba/service"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

const link string = "http://localhost:8080/buys/"

type BuyHandler struct {
	buyService service.BuyService
}

func NewBuyHandler(buyService service.BuyService) BuyHandler {
	return BuyHandler{buyService: buyService}
}

func (ih *BuyHandler) GetAllBuys(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(link)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	var data models.ResponseGet
	json.Unmarshal([]byte(body), &data)
	// fmt.Println(data.Data)
	helpers.CustomResponse(w, r, http.StatusOK, "success", data.Data)
	return
}

func (ih *BuyHandler) CreateBuy(w http.ResponseWriter, r *http.Request) {
	db := models.DataBuy{
		IdUser:      1,
		ItemAmount:  3,
		PriceAmount: 9000,
	}

	bd := make([]models.BuyDetail, 0)
	bd = append(bd, models.BuyDetail{
		IdItem:      10,
		ItemPrice:   3000,
		ItemAmount:  3,
		PriceAmount: 9000,
	})

	reqBody, err := json.Marshal(models.CreateBuy{
		DataBuy:  db,
		DataItem: bd,
	})
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
	}

	res, err := http.Post(link, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
	}

	var data models.ResponseGetSingle
	json.Unmarshal([]byte(body), &data)

	ud := uuid.New().String()
	var transaction = models.Transaction{
		UserId: data.Data.IdUser,
		ItemId: data.Data.ID,
		Price:  data.Data.PriceAmount,
		Uuid:   ud,
	}

	ide := uuid.New().String()
	fmt.Println(ide)

	ress, err := ih.buyService.CreateBuy(&transaction)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", ress)
	return
}
