package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/teerkevin23/rundoo/cmd/web/domain"
)
const (
	CAT = "CATEGORY"
	NAME = "NAME"
	SKU = "SKU"
)
type Database struct {
	typeOfDb string
	database map[string][]domain.Product
}
//type CategoryDatabase struct {
//	TypeOfDb string
//	database map[string][]domain.Product
//}
//type NameDatabase struct {
//	TypeOfDb string
//	database map[string][]domain.Product
//}
//type SKUDatabase struct {
//	TypeOfDb string
//	database map[string][]domain.Product
//}
type productRepository struct {
	database []Database
}

func NewProductRepository() domain.ProductRepository {
	categoryDB := Database{
		typeOfDb: CAT,
		database: map[string][]domain.Product{},
	}
	nameDB := Database{
		typeOfDb: NAME,
		database: map[string][]domain.Product{},
	}
	skuDB := Database{
		typeOfDb: SKU,
		database: map[string][]domain.Product{},
	}
	productDB := make([]Database, 0)
	db := append(productDB, categoryDB, nameDB, skuDB)

	return &productRepository{
		database: db,
	}
}
func (pr *productRepository) Get(c context.Context, filter string) ([]domain.Product, error) {
	pDB := pr.database
	var listOfProducts []domain.Product
	var mapOfProducts = make(map[string]domain.Product)
	if filter == "" {
		// TODO should optimize this...
		for _, db := range pDB {
			for _, subDB := range db.database {
				for _, product := range subDB {
					mapOfProducts[product.SKU] = product
				}
			}
		}

		for _, element := range mapOfProducts {
			listOfProducts = append(listOfProducts, element)
		}

		return listOfProducts, nil
	}

	// all other filters
	for _, db := range pDB {
		m1 := db.database
		val, ok := m1[filter]
		if ok {
			for _, product := range val {
				mapOfProducts[product.SKU] = product
			}
		}
	}

	for _, element := range mapOfProducts {
		listOfProducts = append(listOfProducts, element)
	}

	return listOfProducts, nil
}

func (pr *productRepository) Create(c context.Context, p *domain.Product) error {
	productId :=  primitive.NewObjectID()
	p.ID = productId
	// fake DB connection
	pDB := pr.database

	for _, db := range pDB {
		err := fakeInsertOne(db, *p)
		if err != nil {
			return err
		}
	}

	return nil
}
// decided not to implement these as a result of complete (exact) string matching on fetch and to fetch
// all at once
//func (pr *productRepository) FetchByCategory(c context.Context, categoryString string) ([]domain.Product, error) {
//	var products []domain.Product
//	return products, nil
//}
//func (pr *productRepository) FetchByName(c context.Context, nameString string) ([]domain.Product, error) {
//	var products []domain.Product
//	return products, nil
//}
//func (pr *productRepository) FetchBySKU(c context.Context, skuString string) ([]domain.Product, error) {
//	var products []domain.Product
//	return products, nil
//}

// fake implementation of an InsertOne onto three separate 'databases' as actual slice objects
func fakeInsertOne(db Database, value domain.Product) error {
	// cannot have duplicate SKUs
	if db.typeOfDb == SKU {
		if _, ok := db.database[value.SKU]; ok {
			error := errors.New("Duplicate SKU error: cannot have duplicate SKU")
			SKU_error := domain.NewErrorWrapper(409, error, "bad data")
			return SKU_error
		} else {
			db.database[value.SKU] = append(db.database[value.Name], value)
			return nil
		}
	}
	if db.typeOfDb == CAT {
		db.database[value.Category] = append(db.database[value.Category], value)
		return nil
	}
	if db.typeOfDb == NAME {
		db.database[value.Name] = append(db.database[value.Name], value)
		return nil
	}

	return nil
}