package main

import (
	"log"
	"net/http"

	"game-currency/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Running ...")
	router := httprouter.New()

	// Get Controller instance
	currency := controllers.NewCurrencyController()
	conversion := controllers.NewConversionController()
	web := controllers.NewWebController()

	/** REST API **/
	// Get list all of currency
	router.GET("/api/game-currency/currencies", currency.GetListCurrency)

	router.POST("/api/game-currency/currencies", currency.CreateCurrency)
	router.POST("/api/game-currency/conversion", conversion.CreateConversion)
	router.POST("/api/game-currency/convert", conversion.Convert)

	/** Web Frontend **/
	router.GET("/", web.Index)
	router.GET("/web/currency/list", web.ViewListAllCurrency)

	router.POST("/web/currency/create", web.CreateCurrencyProcess)
	log.Fatal(http.ListenAndServe(":1000", router))
}
