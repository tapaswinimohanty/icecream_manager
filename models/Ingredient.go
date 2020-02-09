package models

const IngredientTable = "ingredients"

type Ingredient struct {
	ID   int    `gorm:"primary_key" json:"ingredient_id"`
	Name string `gorm:"not null" json:"name"`
}

func (Ingredient) TableName() string {
	return IngredientTable
}
