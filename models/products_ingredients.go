package models

const ProductIngredientTable = "products_ingredients"

type ProductIngredient struct {
	ProductID    string `gorm:"primary_key" json:"product_id"`
	IngredientID int    `gorm:"primary_key" json:"ingredient_id"`
}

func (ProductIngredient) TableName() string {
	return ProductIngredientTable
}
