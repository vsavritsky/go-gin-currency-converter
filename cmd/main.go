package main

import (
	"fmt"
)
import currencyConvertor "github.com/vsavritsky/go-gin-currency-converter/pkg/common/service/currencyConvertor"

func main() {
	//currencyLoader.CreateCurrencies()
	//currencyLoader.Load("cbr")

	value := currencyConvertor.ConvertByCode("USD", "RUB", 100, "cbr")

	fmt.Println(value)

	value1 := currencyConvertor.ConvertByCode("USD", "RUB", 100, "cbr")

	fmt.Println(value1)

	value2 := currencyConvertor.ConvertByCode("EUR", "RUB", 100, "cbr")

	fmt.Println(value2)

	value2 = currencyConvertor.ConvertByCode("EUR", "RUB", 100, "cbr")

	fmt.Println(value2)
}
