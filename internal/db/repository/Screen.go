package repository

import (
	"context"
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

func NewDbScreen(screenId *string, name *string, image *string, position *string) *DbScreen {
	return &DbScreen{
		ScreenId: screenId,
		Name:     name,
		Image:    image,
		Position: position,
	}
}

type ScreenRepo struct {
	collection *mongo.Collection
}

func NewScreenRepo(db *mongo.Database) *ScreenRepo {
	return &ScreenRepo{collection: db.Collection(NameScreen)}
}

func (s *ScreenRepo) AddScreen(ctx context.Context, scr *DbScreen) (interface{}, error) {
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
