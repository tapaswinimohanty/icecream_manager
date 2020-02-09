package services

import (
	"bitbucket.com/libertywireless/icecream_manager/lib"
	"bitbucket.com/libertywireless/icecream_manager/models"
	"bitbucket.com/libertywireless/icecream_manager/view_models"
)

func FindOneProductById(id string) *models.Product {
	result := &models.Product{}

	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		First(&result, id)

	return result
}

func FindOneProductModelById(id string) *view_models.ProductModel {
	var product models.Product
	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		First(&product, id)

	if len(product.ID) < 1 {
		return nil
	}

	result := product.ToModel()

	return &result
}

func ProductExistByName(name string) bool {
	count := 0
	lib.DB.Model(&models.Product{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func ProductSourcingExist(productId string, sourcingId int) bool {
	count := 0
	lib.DB.
		Model(&models.ProductSourcingValue{}).
		Where("product_id = ? and sourcing_value_id = ?", productId, sourcingId).
		Count(&count)
	return count > 0
}

func ProductIngredientExist(productId string, ingredientId int) bool {
	count := 0
	lib.DB.
		Model(&models.ProductIngredient{}).
		Where("product_id = ? and ingredient_id = ?", productId, ingredientId).
		Count(&count)
	return count > 0
}

func FindProductModel() []view_models.ProductModel {
	result := make([]view_models.ProductModel, 0)
	var products []models.Product

	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		Find(&products)

	for _, p := range products {
		result = append(result, p.ToModel())
	}

	return result
}
