package currencyConvertor

import (
	"github.com/vsavritsky/go-gin-currency-converter/pkg/common/db"
	"github.com/vsavritsky/go-gin-currency-converter/pkg/common/model/currency"
)

var currencies = make(map[string]currency.Currency)
var rates = make(map[string]float64)

func Convert(currencyFrom currency.Currency, currencyTo currency.Currency, value float64, provider string) float64 {
	query := db.GetDb()

	var code = currencyFrom.Code + "-" + currencyTo.Code

	currencyRateValue := rates[code]
	if currencyRateValue > 0 {
		var result float64 = value * currencyRateValue
		return result
	}

	var currencyRateFrom currency.Rate
	query.Where("currency_id = ?", currencyFrom.ID).
		Where("provider = ?", provider).
		Last(&currencyRateFrom)

	var currencyRateTo currency.Rate
	query.Where("currency_id = ?", currencyTo.ID).
		Where("provider = ?", provider).
		Last(&currencyRateTo)

	rates[code] = currencyRateFrom.Value / currencyRateTo.Value
	var result float64 = (value * currencyRateFrom.Value) / currencyRateTo.Value

	return result
}

func ConvertByCode(currencyFromCode string, currencyToCode string, value float64, provider string) float64 {
	query := db.GetDb()

	currencyFrom := currencies[currencyFromCode]
	if currencyFrom.ID == 0 {
		query.Where("code = ?", currencyFromCode).Last(&currencyFrom)
		currencies[currencyFrom.Code] = currencyFrom
	}

	currencyTo := currencies[currencyToCode]
	if currencyTo.ID == 0 {
		query.Where("code = ?", currencyToCode).Last(&currencyTo)
		currencies[currencyTo.Code] = currencyTo
	}

	if currencyFrom.ID == 0 || currencyTo.ID == 0 {
		return 0
	}

	return Convert(currencyFrom, currencyTo, value, provider)
}
