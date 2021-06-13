package service

import (
	"context"
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
	Holder string      `bson:"holder,omitempty"`
	Data   []string    `bson:"data,omitempty"`
	Name   string      `bson:"name,omitempty"`
	Type   TypeCatalog `bson:"type,omitempty"`
}

type Catalogues interface {
	SAddCatalog(c Catalog) error
}

type CataloguesService struct {
	Cat repository.Catalogues
}

func NewCataloguesService(db *mongo.Database) *CataloguesService {
	return &CataloguesService{Cat: repository.NewCataloguesRepo(db)}
}

func (cat *CataloguesService) SAddCatalog(c Catalog) error {
	_, err := cat.Cat.AddCatalog(context.TODO(), c)
	if err != nil {
		return err
	}
	return nil
}
