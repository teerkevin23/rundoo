package repository_test

import (
	"github.com/rundoo/cmd/web/domain"
	"github.com/rundoo/cmd/web/repository"

	"testing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
const (
	CollectionTest = "test"
)

func TestCreate(t *testing.T) {
	mockProduct := &domain.Product{
		ID: primitive.NewObjectID(),
		Category:  CollectionTest,
		Name: "TestProduct",
		SKU: "123456",
	}
	mockEmptyProduct := &domain.Product{}
	mockProductID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		pr := repository.newProductRepository()
		err := pr.Create(context.Background(), mockProduct)

		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {

	})
}
