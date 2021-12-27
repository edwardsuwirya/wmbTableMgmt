package entity

import "gorm.io/gorm"

type CustomerTable struct {
	ID string `gorm:"column:id;size:3;primaryKey"`
	CustomerTableTransactions []CustomerTableTransaction
	gorm.Model

}
func (c *CustomerTable) TableName() string {
	return "customer_table"
}
