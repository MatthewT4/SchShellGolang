package repository

import (
	"context"
	"fmt"
	"github.com/MatthewT4/SchShellGolang/internal/structions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Screens interface {
	GetImageScreen(ctx context.Context, screenId string) (string, error)
	AddScreen(ctx context.Context, scr structions.Screen) (interface{}, error)
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
		fmt.Println(er.Error())
		return 0, er
	}
	_, err := s.collection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return 0, err
}

/*
func (s *ShellRepo) DelScreen() {

}
*/
func (s *ScreenRepo) GetImageScreen(ctx context.Context, screenId string) (string, error) {
	filter := bson.M{"screenId": screenId}
	var ScreenLimitOnlyGetData struct {
		Image string `bson:"data"`
	}
	err := s.collection.FindOne(ctx, filter).Decode(&ScreenLimitOnlyGetData)
	return ScreenLimitOnlyGetData.Image, err
}

func (s *ScreenRepo) SetDataInScreen(ctx context.Context, screenId, data string) error {
	filter := bson.M{"screen_id": screenId}
	update := bson.M{"data": data}
	res, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Println(res.ModifiedCount, res.UpsertedCount, res.MatchedCount) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return nil
}
