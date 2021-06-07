package repository

import "go.mongodb.org/mongo-driver/mongo"

type Screen struct {
	Id string 		`bson:"id"`
	Name string		`bson:"name"`
	Image string	`bson:"image"`

}

type ScreenRepo struct {
	collection *mongo.Collection
}

func NewShellRepo(db *mongo.Database) *ScreenRepo {
	return &ScreenRepo{collection: db.Collection(NameScreen)}
}
/*
func (s *ShellRepo) AddScreen() {

}

func (s *ShellRepo) DelScreen() {

}

func (s *ShellRepo) SetScreen() {

*/