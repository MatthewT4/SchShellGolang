package repository

import "go.mongodb.org/mongo-driver/mongo"

type ImagesRepo struct {
	collection *mongo.Collection
}

func NewCollectionImages(db *mongo.Database) *ImagesRepo {
	return &ImagesRepo{collection: db.Collection(NameDataImages)}
}
