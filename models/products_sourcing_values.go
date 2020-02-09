package models

const ProductSourcingValueTable = "products_sourcing_values"

type ProductSourcingValue struct {
	ProductID       string `gorm:"primary_key" json:"product_id"`
	SourcingValueID int    `gorm:"primary_key" json:"sourcing_value_id"`
}

func (ProductSourcingValue) TableName() string {
	return ProductSourcingValueTable
}
