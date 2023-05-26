package currencyConvertor

import (
	"github.com/vsavritsky/go-currency-rate/pkg/common/service/cbr"
	"github.com/vsavritsky/go-gin-currency-converter/pkg/common/db"
	cbrCurrency "github.com/vsavritsky/go-gin-currency-converter/pkg/common/model/currency"
	modelCurrency "github.com/vsavritsky/go-gin-currency-converter/pkg/common/model/currency"
)

func CreateCurrencies() {
	query := db.GetDb()
	currencyItem := modelCurrency.Currency{Title: "Доллары", Code: "USD", Sing: "$"}
	query.Create(&currencyItem)
	currencyItem = modelCurrency.Currency{Title: "Рубли", Code: "RUB", Sing: "₽"}
	query.Create(&currencyItem)
	currencyItem = modelCurrency.Currency{Title: "Евро", Code: "EUR", Sing: "€"}
	query.Create(&currencyItem)
}

func Load(provider string) {
	query := db.GetDb()
	if provider == "cbr" {
		rates := cbr.GetCurrencyRates()

		for _, el := range rates {
			var currencyItem modelCurrency.Currency
			query.Where("code = ?", el.CurrencyCode).First(&currencyItem)

			if currencyItem.ID > 0 {
				rate := cbrCurrency.Rate{Value: el.Value, CurrencyID: int(currencyItem.ID), Provider: provider}
				query.Create(&rate)
			}
		}

		var currencyItem modelCurrency.Currency
		query.Where("code = ?", "RUB").First(&currencyItem)
		rate := cbrCurrency.Rate{Value: 1, CurrencyID: int(currencyItem.ID), Provider: provider}
		query.Create(&rate)
	}
}
