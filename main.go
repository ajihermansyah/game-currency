package main

import (
	"log"
	"net/http"

	"github.com/ajihermansyah/game-currency/controllers"

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
	router.GET("/api/game-currency/conversion", conversion.GetListConversionRate)
	router.GET("/api/game-currency/convert", conversion.GetListConversionRate)

	router.POST("/api/game-currency/currencies", currency.CreateCurrency)
	router.POST("/api/game-currency/conversion", conversion.CreateConversion)
	router.POST("/api/game-currency/convert", conversion.Convert)

	/** Web Frontend **/
	router.GET("/", web.Index)
	router.GET("/web/currency/list", web.ViewListAllCurrency)
	router.GET("/web/currency/rate/list", web.ViewListAllCurrencyRate)
	router.GET("/web/currency/rate/read", web.RenderCurrencyRate)
	router.GET("/web/currency/convert/read", web.RenderConvertCurrency)

	router.POST("/web/currency/add", web.CreateCurrencyProcess)
	router.POST("/web/currency/rate/add", web.CreateCurrencyRateProcess)
	router.POST("/web/currency/convert/process", web.ConvertCurrencyProcess)

	log.Fatal(http.ListenAndServe(":1000", router))
}
