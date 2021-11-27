package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"game-currency/models"

	"github.com/julienschmidt/httprouter"
)

type (
	// CurrencyController represents the controller for operating on the currency resource
	CurrencyController struct{}
)

type CurrencyResp struct {
	RespCode string            `json:"response_code"`
	RespDesc string            `json:"response_description"`
	Data     []models.Currency `json:"data"`
}

type CreateCurrencyResp struct {
	RespCode string      `json:"response_code"`
	RespDesc string      `json:"response_description"`
	Data     interface{} `json:"data"`
}

func NewCurrencyController() *CurrencyController {
	return &CurrencyController{}
}

// GetListCurrency retrieves list of all data currency
func (cr CurrencyController) GetListCurrency(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Fetch data from model
	currencies := models.GetAllCurrency()

	//define response
	data := CurrencyResp{"200", "success", currencies}

	// Marshal provided interface into JSON structure
	jm, _ := json.Marshal(data)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jm)
}

func (cr CurrencyController) CreateCurrency(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		currency = models.Currency{}
		d        = CreateCurrencyResp{}
		respCode int
	)

	// Populate the currency data
	json.NewDecoder(r.Body).Decode(&currency)

	isExist := models.CheckCurrencyIsExist(currency.ID)
	if isExist {
		respCode = 400
		message := "currency with ID " + strconv.Itoa(currency.ID) + " is already exist"
		d = CreateCurrencyResp{strconv.Itoa(respCode), message, nil}
	} else {
		respCode = 201

		//Query insert into model
		_, err := models.CreateCurrency(currency)
		if err != nil {
			fmt.Println(err.Error())
		}

		//define response
		d = CreateCurrencyResp{"201", "success", &currency}
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	fmt.Fprintf(w, "%s", uj)
}
