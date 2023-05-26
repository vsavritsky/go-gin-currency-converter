package db

import (
	"github.com/spf13/viper"
	"github.com/vsavritsky/go-gin-currency-converter/pkg/common/model/currency"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var h *gorm.DB = nil

func GetDb() *gorm.DB {
	if h != nil {
		return h
	}

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)
	h = Init(dbUrl)

	return h
}

func Init(url string) *gorm.DB {
	log.Printf(url)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&currency.Currency{})
	db.AutoMigrate(&currency.Rate{})

	return db
}
