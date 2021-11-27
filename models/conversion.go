package models

import (
	"database/sql"
	"fmt"
	"game-currency/config"
	"time"
)

type Conversion struct {
	ID             int        `json:"id"`
	CurrencyIDFrom int        `json:"currency_id_from"`
	CurrencyIDTo   int        `json:"currency_id_to"`
	Rate           int        `json:"rate"`
	CreateAt       *time.Time `json:crated_at`
	UpdatedAt      *time.Time `json:updated_at`
}

func CreateConversion(rate Conversion) (Conversion, error) {
	// insert values
	sqlStr := "INSERT INTO conversions(id, currency_id_from, currency_id_to, rate, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	//prepare the statement
	stmt, err := config.DB.Prepare(sqlStr)
	checkErr(err)
	//format all vals at once
	_, err = stmt.Exec(0, rate.CurrencyIDFrom, rate.CurrencyIDTo, rate.Rate, rate.CreateAt, rate.UpdatedAt)
	checkErr(err)

	return rate, nil
}

func GetConversionBy(from int, to int) (Conversion, error) {
	conversion := Conversion{}

	sqlStr := "SELECT * FROM conversions WHERE (currency_id_from = ? AND currency_id_to = ?) OR (currency_id_to = ? AND currency_id_from = ?)"
	row := config.DB.QueryRow(sqlStr, from, to, from, to)

	err := row.Scan(&conversion.ID, &conversion.CurrencyIDFrom, &conversion.CurrencyIDTo, &conversion.Rate, &conversion.CreateAt, &conversion.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Data ")
			return conversion, nil
		} else {
			return conversion, err
		}
	}

	return conversion, nil
}

func GetAllConversionRate() []Conversion {
	rows, err := config.DB.Query("SELECT * FROM conversions")
	checkErr(err)

	defer rows.Close()

	conversions := make([]Conversion, 0)
	for rows.Next() {
		conversion := Conversion{}
		err := rows.Scan(&conversion.ID, &conversion.CurrencyIDFrom, &conversion.CurrencyIDTo, &conversion.Rate, &conversion.CreateAt, &conversion.UpdatedAt)
		checkErr(err)

		conversions = append(conversions, conversion)
	}
	return conversions
}
