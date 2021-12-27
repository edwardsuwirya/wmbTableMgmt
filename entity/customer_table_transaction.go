package entity

import (
	"gorm.io/gorm"
)

type CustomerTableTransaction struct {
	BillNo          string `gorm:"size:36;unique"`
	CustomerTableID string `gorm:"size:3"`
	gorm.Model
}

func (c *CustomerTableTransaction) TableName() string {
	return "customer_table_transaction"
}
