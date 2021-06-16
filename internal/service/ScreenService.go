package service

import (
	"context"
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Screen struct {
	Id       string `bson:"id"`
	Name     string `bson:"name"`
	Image    string `bson:"image"`
	Position string `bson:"position"`
}

type SScreens interface {
	SGetImage(id string) (int, string)
}

type ScreenService struct {
	Scr repository.Screens
}

func NewScreenService(db *mongo.Database) *ScreenService {
	return &ScreenService{Scr: repository.NewScreenRepo(db)}
}

//SGetImage return Result Code and result (or what error if error there is result)
func (s *ScreenService) SGetImage(id string) (int, string) {
	res, err := s.Scr.GetImageScreen(context.TODO(), id)
	if err != nil {
		return 400, err.Error()
	}
	return 200, res
}
