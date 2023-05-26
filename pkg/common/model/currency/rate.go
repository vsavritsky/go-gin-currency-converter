package currency

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Rate struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	CurrencyID int            `json:"-"`
	Currency   Currency       `json:"currency"`
	Value      float64        `json:"value"`
	Provider   string         `json:"provider"`
}

func (rate *Rate) TableName() string {
	return "currency_rates"
}

func (rate Rate) ToString() string {
	return fmt.Sprintf("id: %d\nprovider: %s\ncurrency: %0.1f\nvalue: %d", rate.ID, rate.Provider, rate.Currency.Code, rate.Value)
}
