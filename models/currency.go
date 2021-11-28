package models

import (
	"github.com/ajihermansyah/game-currency/config"
)

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllCurrency() []Currency {
	rows, err := config.DB.Query("SELECT * FROM currency")
	checkErr(err)

	defer rows.Close()

	currencies := make([]Currency, 0)
	for rows.Next() {
		currency := Currency{}
		err := rows.Scan(&currency.ID, &currency.Name)
		checkErr(err)

		currencies = append(currencies, currency)
	}
	return currencies
}

func CheckCurrencyIsExist(id int) bool {
	var isExist bool
	rows, err := config.DB.Query("SELECT * FROM currency where id = ?", id)
	checkErr(err)

	defer rows.Close()

	currencies := make([]Currency, 0)
	for rows.Next() {
		currency := Currency{}
		err := rows.Scan(&currency.ID, &currency.Name)
		checkErr(err)

		currencies = append(currencies, currency)
	}

	if len(currencies) == 0 {
		isExist = false
	} else {
		isExist = true
	}

	return isExist
}

func CreateCurrency(input Currency) (Currency, error) {
	// insert values
	sqlStr := "INSERT INTO currency(id, name) VALUES (?, ?)"
	//prepare the statement
	stmt, err := config.DB.Prepare(sqlStr)
	checkErr(err)
	//format all vals at once
	_, err = stmt.Exec(input.ID, input.Name)
	checkErr(err)

	return input, nil
}

// function for handle checking error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
