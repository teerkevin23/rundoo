package repository

import (
	"context"
	"github.com/rundoo/cmdmain"
)

type CategoryDatabase struct {
	database map[string][]domain.Product
}
type NameDatabase struct {
	database map[string][]domain.Product
}
type SKUDatabase struct {
	database map[string][]domain.Product
}
type productRepository struct {
	database []interface{}
}

func newProductRepository() domain.ProductRepository {
	categoryDB := CategoryDatabase{
		database: map[string][]domain.Product{},
	}
	nameDB := NameDatabase{
		database: map[string][]domain.Product{},
	}
	skuDB := SKUDatabase{
		database: map[string][]domain.Product{},
	}

	productDB := append(categoryDB, nameDB, skuDB)

	return &productRepository{
		database: productDB,
	}
}

func (pr *productRepository) Create(c context.Context, p *domain.Product) error {
	return nil
}
func (pr *productRepository) FetchByCategory(c context.Context, categoryString string) ([]domain.Product, error) {
	return []domain.Product{}, nil
}
func (pr *productRepository) FetchByName(c context.Context, nameString string) ([]domain.Product, error) {
	return []domain.Product{}, nil
}
func (pr *productRepository) FetchBySKU(c context.Context, skuString string) ([]domain.Product, error) {
	return []domain.Product{}, nil
}