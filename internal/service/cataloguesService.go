package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TypeCatalog int

const (
	Image = iota
	Audio
	Video
)

type Catalog struct {
	Holder string   `bson:"holder,omitempty"`
	Data   []string `bson:"data,omitempty"`
	Name   string   `bson:"name,omitempty"`
	Type   int      `bson:"type,omitempty"`
}

type SCatalogues interface {
	SAddCatalog(c Catalog) (int, error)
	SInsertDataInCatalog(Holder string, NameCatalog string, data string) (int, string)
	SGetCatalogs(holder string) ([]string, error)
	SGetDataInCatalog(holder string, catalogName string) (int, []byte)
}

type CataloguesService struct {
	Cat repository.Catalogues
}

func NewSCataloguesService(db *mongo.Database) *CataloguesService {
	return &CataloguesService{Cat: repository.NewCataloguesRepo(db)}
}

func (cat *CataloguesService) SAddCatalog(c Catalog) (int, error) {
	if c.Holder == "" {
		return 400, errors.New("Holder is null")
	}
	if c.Name == "" {
		return 400, errors.New("Name is null")
	}
	c.Data = append(c.Data, "logo.jpg", "OK.jpg")
	_, err := cat.Cat.AddCatalog(context.TODO(), repository.NewDbCatalog(&c.Holder, &c.Data, &c.Name, &c.Type))

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return 409, errors.New("Duplicate")
		}
		return 500, errors.New("Internal Server Error")
	}
	return 200, nil
}

func (cat *CataloguesService) SGetCatalogs(holder string) ([]string, error) {
	return cat.Cat.GetCatalogs(context.TODO(), holder)
}

func (cat *CataloguesService) SGetDataInCatalog(holder string, catalogName string) (int, []byte) {
	res, err := cat.Cat.GetDataInCatalog(context.TODO(), holder, catalogName)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 404, nil
		}
		fmt.Println(err)
		return 418, nil
	}
	data, er := json.Marshal(res)
	if er != nil {
		fmt.Println(er)
		return 500, nil
	}
	return 200, data
}

func (cat *CataloguesService) SInsertDataInCatalog(Holder string, NameCatalog string, data string) (int, string) {
	_, err := cat.Cat.AddDataInCatalog(context.TODO(), Holder, NameCatalog, data)
	if mongo.IsDuplicateKeyError(err) {
		return 208, "Already Reported"
	}
	if err != nil {
		log.Println(err.Error())
		return 500, "Server Error"
	}
	return 200, "OK"
}
