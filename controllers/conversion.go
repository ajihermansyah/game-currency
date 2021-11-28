package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ajihermansyah/game-currency/models"

	"github.com/julienschmidt/httprouter"
)

type (
	// ConversionController represents the controller for operating on the currency resource
	ConversionController struct{}
)

type ConversionResp struct {
	RespCode   string              `json:"response_code"`
	RespStatus string              `json:"response_status"`
	RespDesc   string              `json:"response_description"`
	Data       []models.Conversion `json:"data"`
}

type CreateConversionResp struct {
	RespCode   string      `json:"response_code"`
	RespStatus string      `json:"response_status"`
	RespDesc   string      `json:"response_description"`
	Data       interface{} `json:"data"`
}

func NewConversionController() *ConversionController {
	return &ConversionController{}
}

func (cr ConversionController) CreateConversion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		conversion = models.Conversion{}
		d          = CreateConversionResp{}
		respCode   int
	)

	// Populate the currency data
	json.NewDecoder(r.Body).Decode(&conversion)

	conversionData, err := models.GetConversionBy(conversion.CurrencyIDFrom, conversion.CurrencyIDTo)
	if err != nil {
		respCode = 400
		d = CreateConversionResp{strconv.Itoa(respCode), "Bad Request", err.Error(), nil}
		uj, _ := json.Marshal(d)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(respCode)
		fmt.Fprintf(w, "%s", uj)
		return
	}

	if conversionData != (models.Conversion{}) {
		respCode = 400
		message := "conversion is already exist"
		d = CreateConversionResp{strconv.Itoa(respCode), "Bad Request", message, nil}
	} else {
		respCode = 201

		//Query insert into model
		_, err := models.CreateConversion(conversion)
		if err != nil {
			fmt.Println(err.Error())
		}

		result := map[string]int{
			"conversion_from": conversion.CurrencyIDFrom,
			"conversion_to":   conversion.CurrencyIDTo,
			"rate":            conversion.Rate,
		}

		//define response
		d = CreateConversionResp{"201", "Success", "Created success", result}
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	fmt.Fprintf(w, "%s", uj)
}

func (cr ConversionController) Convert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		input    = models.ConvertCurrency{}
		response = models.ConvertCurrencyResp{}
		respCode int
	)

	// Populate the currency data
	json.NewDecoder(r.Body).Decode(&input)

	conversionData, err := models.GetConversionBy(input.CurrencyIDFrom, input.CurrencyIDTo)
	if err != nil {
		respCode = 400
		response = models.ConvertCurrencyResp{strconv.Itoa(respCode), err.Error(), nil}
		uj, _ := json.Marshal(response)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(respCode)
		fmt.Fprintf(w, "%s", uj)
		return
	}

	if conversionData.Rate == 0 {
		respCode = 400
		message := "cannot convert from currency " + strconv.Itoa(input.CurrencyIDFrom) + " to currency " + strconv.Itoa(input.CurrencyIDTo)
		response = models.ConvertCurrencyResp{strconv.Itoa(respCode), message, nil}
	} else {
		respCode = 201

		result := 0
		if conversionData.CurrencyIDFrom == input.CurrencyIDFrom {
			result = conversionData.Rate * input.Amount
		} else {
			result = input.Amount / conversionData.Rate
		}

		//define response
		response = models.ConvertCurrencyResp{"201", "success", map[string]int{
			"currency_from": input.CurrencyIDFrom,
			"currency_to":   input.CurrencyIDTo,
			"amount":        input.Amount,
			"result":        result,
		}}
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(response)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	fmt.Fprintf(w, "%s", uj)
}

func (cr ConversionController) GetListConversionRate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Fetch data from model
	conversions := models.GetAllConversionRate()

	//define response
	data := ConversionResp{"200", "Success", "Get list conversion rate succesfully", conversions}

	// Marshal provided interface into JSON structure
	jm, _ := json.Marshal(data)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jm)
}
