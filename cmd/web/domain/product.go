package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProduct = "products"
	ProductPaint = "paint"
	ProductWood = "wood"
)

type Product struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Category string             `json:"category"`
	SKU      string             `json:"sku"`
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	FetchByCategory(c context.Context, categoryString string) ([]Product, error)
	FetchByName(c context.Context, nameString string) ([]Product, error)
	FetchBySKU(c context.Context, skuString string) ([]Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product *Product) error
	FetchByCategory(c context.Context, categoryString string) ([]Product, error)
	FetchByName(c context.Context, nameString string) ([]Product, error)
	FetchBySKU(c context.Context, skuString string) ([]Product, error)
}
