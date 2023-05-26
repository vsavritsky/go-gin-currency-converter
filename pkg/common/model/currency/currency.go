package currency

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Currency struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Title     string         `json:"title"`
	Code      string         `gorm:"index:code,unique" json:"code"`
	Sing      string         `json:"sing"`
}

func (currency Currency) ToString() string {
	return fmt.Sprintf("id: %d\ncode: %s\ntitle: %0.1f", currency.ID, currency.Code, currency.Title)
}
