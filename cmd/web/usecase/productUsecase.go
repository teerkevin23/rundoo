package usecase

import (
	"fmt"
	"context"
	"github.com/teerkevin23/rundoo/cmd/web/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
}

func (usecase *productUsecase) Get(c context.Context, filter string) ([]domain.Product, error) {
	products, err := usecase.productRepository.Get(c, filter)
	fmt.Println("in get", products, err)
	return products, err
}

func (usecase *productUsecase) Create(c context.Context, product *domain.Product) error {
	return usecase.productRepository.Create(c, product)
}

func (usecase *productUsecase) FetchByCategory(c context.Context, categoryString string) ([]domain.Product, error) {
	return nil, nil
}

func (usecase *productUsecase) FetchByName(c context.Context, nameString string) ([]domain.Product, error) {
	return nil, nil
}

func (usecase *productUsecase) FetchBySKU(c context.Context, skuString string) ([]domain.Product, error) {
	return nil, nil
}

func NewProductUsecase(repository domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase {
		productRepository: repository,
	}
}