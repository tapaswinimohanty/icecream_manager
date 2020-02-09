package database

import (
	"encoding/json"
	"bitbucket.com/libertywireless/icecream_manager/lib"
	"bitbucket.com/libertywireless/icecream_manager/models"
	"io/ioutil"
	"log"
)

func Migrate() {
	tables := []interface{}{
		models.ProductSourcingValue{},
		models.ProductIngredient{},
		models.SourcingValue{},
		models.Product{},
		models.Ingredient{},
		models.User{},
	}

	isSeed := lib.Config.Seed
	if isSeed {
		err := lib.DB.DropTableIfExists(tables...).Error
		if err != nil {
			log.Fatal(err)
		}

		err = lib.DB.AutoMigrate(tables...).Error
		if err != nil {
			log.Fatal(err)
		}

		createTable()
		seed()
	} else {
		lib.DB.AutoMigrate(tables...)
	}
}

func createTable() {
	var cascade = "CASCADE"
	if err := lib.DB.Table(models.ProductIngredientTable).AddForeignKey("product_id", models.ProductTable+"(id)", cascade, cascade).Error; err != nil {
		panic(err)
	}
	if err := lib.DB.Table(models.ProductIngredientTable).AddForeignKey("ingredient_id", models.IngredientTable+"(id)", cascade, cascade).Error; err != nil {
		panic(err)
	}
	if err := lib.DB.Table(models.ProductSourcingValueTable).AddForeignKey("product_id", models.ProductTable+"(id)", cascade, cascade).Error; err != nil {
		panic(err)
	}
	if err := lib.DB.Table(models.ProductSourcingValueTable).AddForeignKey("sourcing_value_id", models.SourcingValueTable+"(id)", cascade, cascade).Error; err != nil {
		panic(err)
	}
}

func seed() {
	data, err := ioutil.ReadFile(`database/sample/icecream.json`)
	if err != nil {
		panic(err)
	}

	var mapped []map[string]interface{}
	err = json.Unmarshal(data, &mapped)
	if err != nil {
		panic(err)
	}

	tx := lib.DB.Begin()
	for _, item := range mapped {
		product := models.Product{
			ID:                    item["productId"].(string),
			Name:                  item["name"].(string),
			ImageOpen:             item["image_open"].(string),
			ImageClosed:           item["image_open"].(string),
			Description:           item["description"].(string),
			AllergyInfo:           item["allergy_info"].(string),
			DietaryCertifications: item["dietary_certifications"].(string),
			Story:                 item["story"].(string),
		}

		tx.Create(&product)

		for _, item := range item["ingredients"].([]interface{}) {
			name := item.(string)
			if len(name) == 0 {
				continue
			}

			ingredient := &models.Ingredient{}
			tx.First(&ingredient, "name = ?", name)
			if ingredient.ID == 0 && len(name) > 0 {
				ingredient.Name = name
				tx.Create(&ingredient)
			}

			relateObject := &models.ProductIngredient{}
			tx.First(&relateObject, "product_id = ? and ingredient_id = ?", product.ID, ingredient.ID)

			if relateObject == nil || len(relateObject.ProductID) == 0 {
				relateObject.ProductID = product.ID
				relateObject.IngredientID = ingredient.ID
				tx.Create(&relateObject)
			}
		}

		for _, item := range item["sourcing_values"].([]interface{}) {
			name := item.(string)
			if len(name) == 0 {
				continue
			}

			sourcingValue := &models.SourcingValue{}
			tx.First(&sourcingValue, "name = ?", name)
			if sourcingValue.ID == 0 && len(name) > 0 {
				sourcingValue.Name = name
				tx.Create(&sourcingValue)
			}

			tx.Create(&models.ProductSourcingValue{
				ProductID:       product.ID,
				SourcingValueID: sourcingValue.ID,
			})
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatal(err)
	}

	seedUser()

}

func seedUser() {
	username := "testUser"
	password := "testPassword"

	if err := lib.DB.Create(&models.User{
		Username:       username,
		DisplayName:    "Zalora",
		HashedPassword: lib.HashPassword(password),
	}).Error; err != nil {
		log.Fatal(err)
	}
}
