package service

import "go.mongodb.org/mongo-driver/mongo"

type Service struct {
	CatalogSer Catalogues
}

func NewService(db *mongo.Database) *Service {
	return &Service{NewCataloguesService(db)}
}
