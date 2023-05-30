package repository_test

import (
	"context"
	"github.com/teerkevin23/rundoo/cmd/web/domain"
	"github.com/teerkevin23/rundoo/cmd/web/repository"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
const (
	CollectionPaint = "PAINT"
)

func TestCreate(t *testing.T) {
	mockProduct := &domain.Product{
		ID: primitive.NewObjectID(),
		Category:  CollectionPaint,
		Name: "TestProduct",
		SKU: "123456",
	}
	mockProduct2 := &domain.Product{
		ID: primitive.NewObjectID(),
		Category:  CollectionPaint,
		Name: "TestProduct2",
		SKU: "123456",
	}
	mockEmptyProduct := &domain.Product{}
	mockProductID := primitive.NewObjectID()

	fmt.Println("...testing", mockEmptyProduct, mockProductID)

	t.Run("success", func(t *testing.T) {
		pr := repository.NewProductRepository()
		err := pr.Create(context.Background(), mockProduct)
		err2 := pr.Create(context.Background(), mockProduct2)

		assert.NoError(t, err, err2)
	})
	t.Run("error", func(t *testing.T) {
		pr := repository.NewProductRepository()
		err := pr.Create(context.Background(), mockProduct)

		assert.Error(t, err)
	})
}
