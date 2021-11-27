package controllers

import (
	"bytes"
	"encoding/json"
	"game-currency/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"game-currency/config"

	"github.com/julienschmidt/httprouter"
)

type (
	// TaxCodeController represents the controller for operating on the Tax Code resource
	WebController struct{}
)

func NewWebController() *WebController {
	return &WebController{}
}

//custom var
type BaseUrl string

const (
	development BaseUrl = "http://localhost:1000/"
	staging     BaseUrl = "http://localhost:2000/"
	production  BaseUrl = "http://localhost:3000/"
)

func (web WebController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("api/game-currency/currencies")
	url := buffer.String()
	body := RequestGet(url)

	currency := CurrencyResp{}
	jsonErr := json.Unmarshal(body, &currency)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	data := currency.Data

	config.TPL.ExecuteTemplate(w, "create.gohtml", data)
}

func (web WebController) CreateCurrencyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get form values
	formCurrencyId := r.FormValue("currency_id")
	formCurrencyName := r.FormValue("currency_name")
	// string to int
	currencyId, err := strconv.Atoi(formCurrencyId)
	if err != nil {
		log.Fatal(err)
	}

	//store new currency object
	currency := models.Currency{}
	currency.ID = currencyId
	currency.Name = formCurrencyName

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(currency)

	//POST localhost:1000/api/game-currency/currencies
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("api/game-currency/currencies")
	url := buffer.String()

	//make post request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(uj))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	data := struct {
		RespStatus string
		RespBody   string
	}{
		resp.Status,
		string(body),
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", data)
}

func (web WebController) ViewListAllCurrency(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("api/game-currency/currencies")
	url := buffer.String()
	body := RequestGet(url)

	currency := CurrencyResp{}
	jsonErr := json.Unmarshal(body, &currency)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	data := currency.Data

	response := struct{ Data []models.Currency }{
		data,
	}

	config.TPL.ExecuteTemplate(w, "list.gohtml", response)
}

func RequestGet(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "api/game-currency")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
