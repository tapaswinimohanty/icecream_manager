package models

import "bitbucket.com/libertywireless/icecream_manager/view_models"

const ProductTable = "products"

type Product struct {
	ID          string `gorm:"primary_key" json:"product_id"`
	Name        string `gorm:"not null" json:"name"`
	ImageOpen   string `json:"image_open"`
	ImageClosed string `json:"image_closed"`
	Description string `json:"description"`
	Story       string `gorm:"not null" json:"story"`

	AllergyInfo           string `json:"allergy_info"`
	DietaryCertifications string `json:"dietary_certifications"`

	SourcingValues []SourcingValue `gorm:"many2many:products_sourcing_values" json:"sourcing_values"`
	Ingredients    []Ingredient    `gorm:"many2many:products_ingredients" json:"ingredients"`
}

func (Product) TableName() string {
	return ProductTable
}

func (p Product) ToModel() view_models.ProductModel {
	result := view_models.ProductModel{
		ProductId:             p.ID,
		Name:                  p.Name,
		ImageOpen:             p.ImageOpen,
		ImageClosed:           p.ImageClosed,
		Description:           p.Description,
		Story:                 p.Story,
		AllergyInfo:           p.AllergyInfo,
		DietaryCertifications: p.DietaryCertifications,
	}

	ingredients := make([]string, 0)
	for _, ingredient := range p.Ingredients {
		ingredients = append(ingredients, ingredient.Name)
	}
	result.Ingredients = ingredients

	sourcingValues := make([]string, 0)
	for _, sourcingValue := range p.SourcingValues {
		sourcingValues = append(sourcingValues, sourcingValue.Name)
	}
	result.SourcingValues = sourcingValues

	return result
}
