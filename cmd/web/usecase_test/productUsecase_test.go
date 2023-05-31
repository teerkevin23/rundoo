package usecase_test

import (
	"context"
	"testing"
	"github.com/teerkevin23/rundoo/cmd/web/domain/mocks"
	"github.com/teerkevin23/rundoo/cmd/web/domain"
	"github.com/teerkevin23/rundoo/cmd/web/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
func TestGet(t *testing.T) {
	mockProductRepository := new(mocks.ProductRepository)

	t.Run("success", func(t *testing.T) {
		mockProduct := domain.Product{
			ID: primitive.NewObjectID(),
			Category: domain.ProductPaint,
			Name: "eggshell",
			SKU: "123456",
		}

		mockListProduct := make([]domain.Product, 0)
		mockListProduct = append(mockListProduct, mockProduct)

		mockProductRepository.On("Get", mock.Anything, "paint").Return(mockListProduct, nil).Once()

		p := usecase.NewProductUsecase(mockProductRepository)

		list, err := p.Get(context.Background(), "paint")

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockListProduct))

		mockProductRepository.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockProduct := domain.Product{
			ID: primitive.NewObjectID(),
			Category: domain.ProductPaint,
			Name: "eggshell",
			SKU: "123456",
		}

		mockListProduct := make([]domain.Product, 0)
		mockListProduct = append(mockListProduct, mockProduct)

		mockProductRepository.On("Get", mock.Anything, "hardware").Return(mockListProduct, nil).Once()

		p := usecase.NewProductUsecase(mockProductRepository)

		list, err := p.Get(context.Background(), "hardware")

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, 1)

		mockProductRepository.AssertExpectations(t)
	})
}


/*
Originally wrote these tests but then decided that this was overengineering based on the specs and did consolidated GET instead
 */
//func TestFetchByCategory(t *testing.T) {
//	mockProductRepository := new(mocks.ProductRepository)
//
//	t.Run("success", func(t *testing.T) {
//		mockProduct := domain.Product{
//			ID: primitive.NewObjectID(),
//			Category: domain.ProductPaint,
//			Name: "eggshell",
//			SKU: "123456",
//		}
//
//		mockListProduct := make([]domain.Product, 0)
//		mockListProduct = append(mockListProduct, mockProduct)
//
//		mockProductRepository.On("FetchByCategory", mock.Anything, "paint").Return(mockListProduct, nil).Once()
//
//		p := usecase.NewProductUsecase(mockProductRepository)
//
//		list, err := p.FetchByCategory(context.Background(), "pain")
//
//		assert.NoError(t, err)
//		assert.NotNil(t, list)
//		assert.Len(t, list, len(mockListProduct))
//
//		mockProductRepository.AssertExpectations(t)
//	})
//	t.Run("error", func(t *testing.T) {
//		mockProduct := domain.Product{
//			ID: primitive.NewObjectID(),
//			Category: domain.ProductPaint,
//			Name: "eggshell",
//			SKU: "123456",
//		}
//
//		mockListProduct := make([]domain.Product, 0)
//		mockListProduct = append(mockListProduct, mockProduct)
//
//		mockProductRepository.On("FetchByCategory", mock.Anything, "hardware").Return(mockListProduct, nil).Once()
//
//		p := usecase.NewProductUsecase(mockProductRepository)
//
//		list, err := p.FetchByCategory(context.Background(), "hardware")
//
//		assert.NoError(t, err)
//		assert.Nil(t, list)
//		assert.Len(t, list, 0)
//
//		mockProductRepository.AssertExpectations(t)
//	})
//}
//func TestFetchByName(t *testing.T) {
//	mockProductRepository := new(mocks.ProductRepository)
//
//	t.Run("success", func(t *testing.T) {
//		mockProduct := domain.Product{
//			ID: primitive.NewObjectID(),
//			Category: domain.ProductPaint,
//			Name: "eggshell",
//			SKU: "123456",
//		}
//
//		mockListProduct := make([]domain.Product, 0)
//		mockListProduct = append(mockListProduct, mockProduct)
//
//		mockProductRepository.On("FetchByName", mock.Anything, "paint").Return(mockListProduct, nil).Once()
//
//		p := usecase.NewProductUsecase(mockProductRepository)
//
//		list, err := p.FetchByName(context.Background(), "pain")
//
//		assert.NoError(t, err)
//		assert.NotNil(t, list)
//		assert.Len(t, list, len(mockListProduct))
//
//		mockProductRepository.AssertExpectations(t)
//	})
//	t.Run("error", func(t *testing.T) {
//		mockProduct := domain.Product{
//			ID: primitive.NewObjectID(),
//			Category: domain.ProductPaint,
//			Name: "eggshell",
//			SKU: "123456",
//		}
//
//		mockListProduct := make([]domain.Product, 0)
//		mockListProduct = append(mockListProduct, mockProduct)
//
//		mockProductRepository.On("FetchByName", mock.Anything, "hardware").Return(mockListProduct, nil).Once()
//
//		p := usecase.NewProductUsecase(mockProductRepository)
//
//		list, err := p.FetchByName(context.Background(), "hardware")
//
//		assert.NoError(t, err)
//		assert.Nil(t, list)
//		assert.Len(t, list, 0)
//
//		mockProductRepository.AssertExpectations(t)
//	})
//}
//func TestFetchBySKU(t *testing.T) {
//	mockProductRepository := new(mocks.ProductRepository)
//
//	t.Run("success", func(t *testing.T) {
//		mockProduct := domain.Product{
//			ID: primitive.NewObjectID(),
//			Category: domain.ProductPaint,
//			Name: "eggshell",
//			SKU: "123456",
//		}
//
//		mockListProduct := make([]domain.Product, 0)
//		mockListProduct = append(mockListProduct, mockProduct)
//
//		mockProductRepository.On("FetchBySKU", mock.Anything, "123").Return(mockListProduct, nil).Once()
//
//		p := usecase.NewProductUsecase(mockProductRepository)
//
//		list, err := p.FetchBySKU(context.Background(), "123")
//
//		assert.NoError(t, err)
//		assert.NotNil(t, list)
//		assert.Len(t, list, len(mockListProduct))
//
//		mockProductRepository.AssertExpectations(t)
//	})
//	t.Run("error", func(t *testing.T) {
//		mockProduct := domain.Product{
//			ID: primitive.NewObjectID(),
//			Category: domain.ProductPaint,
//			Name: "eggshell",
//			SKU: "123456",
//		}
//
//		mockListProduct := make([]domain.Product, 0)
//		mockListProduct = append(mockListProduct, mockProduct)
//
//		mockProductRepository.On("FetchBySKU", mock.Anything, "123").Return(mockListProduct, nil).Once()
//
//		p := usecase.NewProductUsecase(mockProductRepository)
//
//		list, err := p.FetchBySKU(context.Background(), "123")
//
//		assert.NoError(t, err)
//		assert.Nil(t, list)
//		assert.Len(t, list, 0)
//
//		mockProductRepository.AssertExpectations(t)
//	})
//}
