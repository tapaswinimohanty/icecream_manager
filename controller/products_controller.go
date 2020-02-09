package controller

import (
	"github.com/labstack/echo"
	"bitbucket.com/libertywireless/icecream_manager/lib"
	"bitbucket.com/libertywireless/icecream_manager/models"
	"bitbucket.com/libertywireless/icecream_manager/services"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func ProductListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, services.FindProductModel())
}

func ProductGetByIDHandler(c echo.Context) error {
	cc := c.(*lib.CustomContext)
	id := c.Param("id")
	if len(id) < 1 {
		return cc.BadRequest("Id is required")
	}

	product := services.FindOneProductModelById(id)

	if product == nil {
		return cc.NotFound("Product not found")
	}

	return cc.OK(product)
}

func ProductAddHandler(c echo.Context) error {
	cc := c.(*lib.CustomContext)

	name := c.FormValue("name")
	description := c.FormValue("description")
	story := c.FormValue("story")
	allergyInfo := c.FormValue("allergy_info")
	dietaryCertifications := c.FormValue("dietary_certifications")

	sourcing_value_ids := strings.TrimSpace(c.FormValue("sourcing_value_ids"))
	ingredient_ids := strings.TrimSpace(c.FormValue("ingredient_ids"))

	image_open, err := c.FormFile("image_open")
	if err != nil {
		return cc.BadRequest("Open image is not valid")
	}

	imageClosed, err := c.FormFile("image_closed")
	if err != nil {
		return cc.BadRequest("Closed image is not valid")
	}

	if len(name) == 0 {
		return cc.BadRequest("Product Name is required")
	}

	if len(description) == 0 {
		return cc.BadRequest("Description is required")
	}

	if len(story) == 0 {
		return cc.BadRequest("Story is required")
	}

	if image_open == nil {
		return cc.BadRequest("Open image is required")
	}

	if imageClosed == nil {
		return cc.BadRequest("Closed image is required")
	}

	if len(sourcing_value_ids) == 0 {
		return cc.BadRequest("Sourcing values are required")
	}

	if len(ingredient_ids) == 0 {
		return cc.BadRequest("Ingredients are required")
	}

	sourcingValueIds := strings.Split(sourcing_value_ids, ",")
	if len(sourcingValueIds) == 0 {
		return cc.BadRequest("Sourcing values are required")
	}

	count := 0
	lib.DB.Model(&models.SourcingValue{}).Where("id in (?)", sourcingValueIds).Count(&count)
	if count == 0 {
		return cc.NotFound("Sourcing value is not exist")
	}

	ingredientIds := strings.Split(ingredient_ids, ",")
	if len(ingredientIds) == 0 {
		return cc.BadRequest("Ingredients are required")
	}

	lib.DB.Model(&models.Ingredient{}).Where("id in (?)", ingredientIds).Count(&count)
	if count == 0 {
		return cc.NotFound("Ingredient is not exist")
	}

	if services.ProductExistByName(name) {
		return cc.Conflict("Product Name has already exist")
	}

	product := models.Product{
		ID:                    strconv.Itoa(rand.Intn(999999)),
		Name:                  name,
		Description:           description,
		DietaryCertifications: dietaryCertifications,
		AllergyInfo:           allergyInfo,
		Story:                 story,
	}

	url, err := lib.Upload(image_open)
	if err != nil {
		return cc.Internal("Can not upload open image file")
	}
	product.ImageOpen = *url

	url, err = lib.Upload(imageClosed)
	if err != nil {
		return cc.Internal("Can not upload closed image file")
	}
	product.ImageClosed = *url

	tx := lib.DB.Begin()
	if err = tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return cc.Internal("Can not create product")
	}

	for _, ingredientIdStr := range ingredientIds {
		ingredientId, err := strconv.Atoi(ingredientIdStr)
		if err != nil {
			return cc.BadRequest("Ingredient ID should be integer")
		}

		//lib.DB.
		//	Model(&models.ProductIngredient{}).
		//	Where("product_id = ? and ingredient_id = ?", product.ID, ingredientId).
		//	Count(&count)
		//
		//if count > 0 {
		//	tx.Rollback()
		//	return cc.Conflict("Product Ingredient has already exist")
		//}

		if err := tx.Create(&models.ProductIngredient{
			ProductID:    product.ID,
			IngredientID: ingredientId,
		}).Error; err != nil {
			tx.Rollback()
			return cc.Internal("Can not create product's ingredient")
		}
	}

	for _, sourcingValueIdStr := range sourcingValueIds {
		sourcingValueId, err := strconv.Atoi(sourcingValueIdStr)
		if err != nil {
			return cc.BadRequest("Sourcing Value ID should be integer")
		}

		//lib.DB.
		//	Model(&models.ProductSourcingValue{}).
		//	Where("product_id = ? and sourcing_value_id = ?", product.ID, sourcingValueId).
		//	Count(&count)
		//
		//if count > 0 {
		//	tx.Rollback()
		//	return cc.Conflict("Product Sourcing Value has already exist")
		//}

		if err := tx.Create(&models.ProductSourcingValue{
			ProductID:       product.ID,
			SourcingValueID: sourcingValueId,
		}).Error; err != nil {
			tx.Rollback()
			return cc.Internal("Can not create product's sourcing value")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return cc.Internal("Can not create product")
	}

	return c.JSON(http.StatusCreated, services.FindOneProductModelById(product.ID))
}

func ProductUpdateHandler(c echo.Context) error {
	cc := c.(*lib.CustomContext)
	id := c.Param("id")
	if len(id) < 1 {
		return cc.BadRequest("Product ID is required")
	}

	name := c.FormValue("name")
	description := c.FormValue("description")
	story := c.FormValue("story")
	allergyInfo := c.FormValue("allergy_info")
	dietaryCertifications := c.FormValue("dietary_certifications")

	sourcing_value_ids := strings.TrimSpace(c.FormValue("sourcing_value_ids"))
	ingredient_ids := strings.TrimSpace(c.FormValue("ingredient_ids"))

	product := services.FindOneProductById(id)
	if product == nil {
		return cc.NotFound("Product not found")
	}

	if len(name) > 0 {
		product.Name = name
	}
	if len(description) > 0 {
		product.Description = description
	}
	if len(story) > 0 {
		product.Story = story
	}
	if len(allergyInfo) > 0 {
		product.AllergyInfo = allergyInfo
	}
	if len(dietaryCertifications) > 0 {
		product.DietaryCertifications = dietaryCertifications
	}

	imageClosed, _ := c.FormFile("image_closed")
	if imageClosed != nil {
		url, err := lib.Upload(imageClosed)
		if err != nil {
			return cc.Internal("Can not upload closed image file")
		}
		product.ImageClosed = *url
	}

	imageOpen, _ := c.FormFile("image_open")
	if imageOpen != nil {
		url, err := lib.Upload(imageOpen)
		if err != nil {
			return cc.Internal("Can not upload open image file")
		}
		product.ImageOpen = *url
	}

	tx := lib.DB.Begin()

	if err := tx.Save(&product).Error; err != nil {
		return cc.Internal("Can not update product")
	}

	if len(sourcing_value_ids) > 0 {
		sourcingValueIds := strings.Split(sourcing_value_ids, ",")

		if len(sourcingValueIds) > 0 {
			for _, sourcingValueIdStr := range sourcingValueIds {
				sourcingValueId, err := strconv.Atoi(sourcingValueIdStr)
				if err != nil {
					tx.Rollback()
					return cc.BadRequest("Sourcing Value ID should be integer")
				}

				// if product does not have any sourcing value -> create one
				if !services.ProductSourcingExist(product.ID, sourcingValueId) {
					if err := tx.Create(&models.ProductSourcingValue{
						ProductID:       product.ID,
						SourcingValueID: sourcingValueId,
					}).Error; err != nil {
						tx.Rollback()
						return cc.Internal("Can not create product's sourcing value")
					}
				}

				for _, sourcing := range product.SourcingValues {
					if sourcing.ID != sourcingValueId {
						if err := tx.Unscoped().Delete(&models.ProductSourcingValue{
							ProductID:       product.ID,
							SourcingValueID: sourcing.ID,
						}).Error; err != nil {
							tx.Rollback()
							return cc.Internal("Can not delete product's sourcing")
						}
					}
				}

			}
		}
	}

	if len(ingredient_ids) > 0 {
		ingredientIds := strings.Split(ingredient_ids, ",")

		if len(ingredientIds) > 0 {
			for _, idString := range ingredientIds {
				ingredientId, err := strconv.Atoi(idString)
				if err != nil {
					tx.Rollback()
					return cc.BadRequest("Sourcing Value ID should be integer")
				}

				// if product does not have any sourcing value -> create one
				if !services.ProductIngredientExist(product.ID, ingredientId) {
					if err := tx.Create(&models.ProductIngredient{
						ProductID:    product.ID,
						IngredientID: ingredientId,
					}).Error; err != nil {
						tx.Rollback()
						return cc.Internal("Can not create product's ingredient")
					}
				}

				for _, ingr := range product.Ingredients {
					if ingr.ID != ingredientId {
						if err := tx.Unscoped().Delete(&models.ProductIngredient{
							ProductID:    product.ID,
							IngredientID: ingr.ID,
						}).Error; err != nil {
							tx.Rollback()
							return cc.Internal("Can not delete product's ingredient")
						}
					}
				}
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return cc.Internal("Can not update product")
	}

	return cc.OK(services.FindOneProductModelById(product.ID))
}

func ProductDeleteHandler(c echo.Context) error {
	cc := c.(*lib.CustomContext)
	id := c.Param("id")
	if len(id) < 1 {
		return cc.BadRequest("Product ID is required")
	}

	product := services.FindOneProductById(id)
	if product == nil || len(product.ID) == 0 {
		return cc.NotFound("Product not found")
	}

	if err := lib.DB.Unscoped().Delete(&product).Error; err != nil {
		return cc.Internal("Can not delete product")
	}

	return cc.OK(echo.Map{
		"product_id": product.ID,
	})
}
