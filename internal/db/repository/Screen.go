package repository

import (
	"context"
	"github.com/MatthewT4/SchShellGolang/internal/structions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Screens interface {
	GetImageScreen(ctx context.Context, screenId string) (string, error)
}

type DbScreen struct {
	ScreenId *string `bson:"screenId"`
	Name     *string `bson:"name"`
	Image    *string `bson:"image"`
	Position *string `bson:"position"`
}

type ScreenRepo struct {
	collection *mongo.Collection
}

func NewScreenRepo(db *mongo.Database) *ScreenRepo {
	return &ScreenRepo{collection: db.Collection(NameScreen)}
}

func (s *ScreenRepo) AddScreen(ctx context.Context, scr structions.Screen) (interface{}, error) {
	data, er := bson.Marshal(scr)
	if er != nil {
		return 0, er
	}
	res, err := s.collection.InsertOne(ctx, data)
	return res.InsertedID, err
}

/*
func (s *ShellRepo) DelScreen() {

}
*/
func (s *ScreenRepo) GetImageScreen(ctx context.Context, screenId string) (string, error) {
	filter := bson.M{"screenId": screenId}
	var ScreenLimitOnlyGetData struct {
		Image string `bson:"image"`
	}
	err := s.collection.FindOne(ctx, filter).Decode(&ScreenLimitOnlyGetData)
	return ScreenLimitOnlyGetData.Image, err
}
