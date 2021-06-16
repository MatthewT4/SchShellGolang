package service

import (
	"context"
	"errors"
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"go.mongodb.org/mongo-driver/mongo"
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

	_, err := cat.Cat.AddCatalog(context.TODO(), repository.NewDbCatalog(&c.Holder, &c.Data, &c.Name, &c.Type))

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return 409, errors.New("Duplicate")
		}
		return 500, errors.New("Internal Server Error")
	}
	return 200, nil
}
