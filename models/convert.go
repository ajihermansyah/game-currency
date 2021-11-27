package models

type ConvertCurrency struct {
	CurrencyIDFrom int `json:"currency_id_from"`
	CurrencyIDTo   int `json:"currency_id_to"`
	Amount         int `json:"amount"`
}

type ConvertCurrencyResp struct {
	RespCode string      `json:"response_code"`
	RespDesc string      `json:"response_description"`
	Data     interface{} `json:"data"`
}
