package usecase_test

import (
	"context"
	"testing"

	"github.com/teerkevin23/rundoo/cmd/web/domain"
	"github.com/teerkevin23/rundoo/cmd/web/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFetchByCategory(t *testing.T) {
	mockProductRepository := new(domain.ProductRepository)

	t.Run("success", func(t *testing.T) {
		mockProduct := domain.Product{
			ID: primitive.NewObjectID(),
			Category: domain.ProductPaint,
			Name: "eggshell",
			SKU: "XYZ12345",
		}

		mockListProduct := make([]domain.Product, 0)
		mockListProduct = append(mockListProduct, mockProduct)

		mockProductRepository.On("FetchByCategory", mock.Anything, "paint").Return(mockListProduct, nil).Once()

		p := usecase.NewProductUsecase(mockProductRepository)

		list, err := p.FetchByCategory(context.Background(), "pain")

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockListProduct))

		mockProductRepository.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {

	})
}
func TestFetchByName(t *testing.T) {
	t.Run("success", func(t *testing.T) {

	})
	t.Run("error", func(t *testing.T) {

	})
}
func TestFetchBySKU(t *testing.T) {
	t.Run("success", func(t *testing.T) {

	})
	t.Run("error", func(t *testing.T) {

	})
}
