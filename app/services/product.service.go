package services

	import (
		"rest_go/app/models"
		"rest_go/app/types"
	)
	
	func ShowAllProduct() ([]models.Product, error, types.MetaData) {
		products, err := models.ProductQuery.FindAll()
		if err != nil {
			return nil, err, types.MetaData{}
		}
		totalItems, err := models.ProductQuery.Count()
		if err != nil {
			return nil, err, types.MetaData{}
		}
		metaData := types.MetaData{
			TotalItems: totalItems,
		}
	
		return products, nil, metaData
	}
	
	func CreateProduct(input types.ProductRequest) (models.Product, error) {
		product := models.Product{
			// Add your fields here
		}
	
		createdProduct, err := models.ProductQuery.Create(product)
		if err != nil {
			return models.Product{}, err
		}
	
		return *createdProduct, nil
	}
	
	func GetProductByID(ID uint) (models.Product, error) {
		product, err := models.ProductQuery.FindByID(ID)
		if err != nil {
			return models.Product{}, err
		}
	
		return *product, nil
	}
	
	func UpdateProduct(ID uint, input types.ProductRequest) (models.Product, error) {
		product, err := models.ProductQuery.FindByID(ID)
		if err != nil {
			return models.Product{}, err
		}
	
		// Add your fields here
	
		updatedProduct, err := models.ProductQuery.Update(*product)
		if err != nil {
			return models.Product{}, err
		}
	
		return *updatedProduct, nil
	}
	
	func DeleteProduct(ID uint) error {
		err := models.ProductQuery.DeleteByID(ID)
		if err != nil {
			return err
		}
	
		return nil
	}