package usecase

import "github.com/rundoo/cmd/web/domain"

type productUsecase struct {
	productRepository domain.ProductRepository
}
/*
func (productUsecase) Create(c interface{}, product *domain.Product) interface{} {
	panic("implement me")
}
 */
func (productUsecase) Create(c interface{}, product *domain.Product) error {
	panic("implement me")
}

func (productUsecase) FetchByCategory(c context.Context, categoryString string) ([]domain.Product, interface{}) {
	panic("implement me")
}

func (productUsecase) FetchByName(c context.Context, nameString string) ([]domain.Product, interface{}) {
	panic("implement me")
}

func (productUsecase) FetchBySKU(c context.Context, skuString string) ([]domain.Product, interface{}) {
	panic("implement me")
}
/*
func (productUsecase) FetchBySKU(c interface{}, skuString interface{}) ([]domain.Product, interface{}) {
	panic("implement me")
}
 */

func NewProductUsecase(repository domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase {
		productRepository: repository,
	}
}