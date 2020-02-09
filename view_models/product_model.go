package view_models

type ProductModel struct {
	ProductId   string `json:"product_id"`
	Name        string `json:"name"`
	ImageOpen   string `json:"image_open"`
	ImageClosed string `json:"image_closed"`
	Description string `json:"description"`

	Story       string `json:"story"`
	AllergyInfo string `json:"allergy_info"`

	DietaryCertifications string `json:"dietary_certifications"`

	SourcingValues []string `json:"sourcing_values"`
	Ingredients    []string `json:"ingredients"`
}

type ProductAddModel struct {
	Name        string `json:"name"`
	ImageOpen   string `json:"image_open"`
	ImageClosed string `json:"image_closed"`
	Description string `json:"description"`

	Story       string `json:"story"`
	AllergyInfo string `json:"allergy_info"`

	DietaryCertifications string `json:"dietary_certifications"`

	SourcingValues []int `json:"sourcing_values"`
	Ingredients    []int `json:"ingredients"`
}

type ProductUpdateModel struct {
	ProductId   string `json:"product_id"`
	Name        string `json:"name"`
	ImageOpen   string `json:"image_open"`
	ImageClosed string `json:"image_closed"`
	Description string `json:"description"`

	Story       string `json:"story"`
	AllergyInfo string `json:"allergy_info"`

	DietaryCertifications string `json:"dietary_certifications"`

	SourcingValues []int `json:"sourcing_values"`
	Ingredients    []int `json:"ingredients"`
}
