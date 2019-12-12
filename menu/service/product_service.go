package service

import (
	".."
	"../../entity"
)

// CategoryService implements menu.CategoryService interface
type ProductService struct {
	productRepo menu.ProductRepository
}

// NewCategoryService will create new CategoryService object
func NewProductService(ProRepo menu.ProductRepository) *ProductService {
	return &ProductService{productRepo: ProRepo}
}

// Categories returns list of categories
func (ps *ProductService) Products() ([]entity.Product, error) {

	products, err := ps.productRepo.Products()

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (ps *ProductService) CamProducts() ([]entity.Product, error) {

	products, err := ps.productRepo.CamProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (ps *ProductService) CompProducts() ([]entity.Product, error) {

	products, err := ps.productRepo.CompProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (ps *ProductService) LapProducts() ([]entity.Product, error) {

	products, err := ps.productRepo.LapProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (ps *ProductService) MobProducts() ([]entity.Product, error) {

	products, err := ps.productRepo.MobProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

// StoreCategory persists new category information
func (ps *ProductService) StoreProduct(product entity.Product) error {

	err := ps.productRepo.StoreProduct(product)

	if err != nil {
		return err
	}

	return nil
}

// Category returns a category object with a given id
func (ps *ProductService) Product(id int) (entity.Product, error) {

	p, err := ps.productRepo.Product(id)

	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps *ProductService) SearchProduct(index string) ([]entity.Product, error) {
	products, err := ps.productRepo.SearchProduct(index)

	if err != nil {
		return nil, err
	}

	return products, nil
}

// UpdateCategory updates a cateogory with new data
func (ps *ProductService) UpdateProduct(product entity.Product) error {

	err := ps.productRepo.UpdateProduct(product)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a category by its id
func (ps *ProductService) DeleteProduct(id int) error {

	err := ps.productRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) RateProduct(pro entity.Product) (entity.Product, error) {

	prowithrate, err := ps.productRepo.RateProduct(pro)
	if err != nil {
		return prowithrate, err
	}
	return prowithrate, nil
}



