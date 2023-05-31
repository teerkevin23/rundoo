package repository

import (
	"context"
	"errors"
	"fmt"
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
	if filter == "" {
		fmt.Println("no filter")
		for _, db := range pDB {
			fmt.Println(db)
			for _, subDB := range db.database {
				fmt.Println(subDB)
				for _, product := range subDB {
					fmt.Println(product)
					listOfProducts = append(listOfProducts, product)
				}
			}
		}

		return listOfProducts, nil
	}
	// do more

	return nil, nil
}

func (pr *productRepository) Create(c context.Context, p *domain.Product) error {
	productId :=  primitive.NewObjectID()
	p.ID = productId
	// fake DB connection
	pDB := pr.database
	fmt.Println("in create!!!", p, pDB)

	for index, db := range pDB {
		fmt.Println(index, db.typeOfDb, db.database)
		err := fakeInsertOne(db, *p)
		if err != nil {
			return err
		}
	}

	for index, db := range pDB {
		fmt.Println("***********")
		fmt.Println(index, db.typeOfDb, db.database)
	}


	return nil
}
func (pr *productRepository) FetchByCategory(c context.Context, categoryString string) ([]domain.Product, error) {
	var products []domain.Product
	return products, nil
}
func (pr *productRepository) FetchByName(c context.Context, nameString string) ([]domain.Product, error) {
	var products []domain.Product
	return products, nil
}
func (pr *productRepository) FetchBySKU(c context.Context, skuString string) ([]domain.Product, error) {
	var products []domain.Product
	return products, nil
}

func fakeInsertOne(db Database, value domain.Product) error {
	fmt.Println(db.typeOfDb, db.database, value)
	if db.typeOfDb == CAT {
		fmt.Println("found cat", value.Category)
		if _, ok := db.database[value.Category]; ok {
			fmt.Println("!!! KEY FOUND", db.database[value.Category])
			db.database[value.Category] = append(db.database[value.Category], value)
		} else {
			fmt.Println("key DNE")
			db.database[value.Category] = append(db.database[value.Category], value)
		}
		fmt.Println("end", db.database)
	}
	if db.typeOfDb == NAME {
		fmt.Println("found name", value.Name)
		if _, ok := db.database[value.Name]; ok {
			db.database[value.Name] = append(db.database[value.Name], value)
		} else {
			fmt.Println("key DNE")
			db.database[value.Name] = append(db.database[value.Name], value)
		}
	}
	if db.typeOfDb == SKU {
		fmt.Println("found sku", value.SKU)
		if _, ok := db.database[value.SKU]; ok {
			SKU_error := errors.New("Duplicate SKU")
			fmt.Println(SKU_error)
			return SKU_error
		} else {
			fmt.Println("key DNE")
			db.database[value.SKU] = append(db.database[value.Name], value)
		}
	}

	return nil
}