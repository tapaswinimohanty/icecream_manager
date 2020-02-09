package models

const SourcingValueTable = "sourcing_values"

type SourcingValue struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (SourcingValue) TableName() string {
	return SourcingValueTable
}
