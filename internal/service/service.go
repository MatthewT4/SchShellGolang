package service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	User = iota
	Moderation
	Administration
)

type Service struct {
	CatalogSer SCatalogues
	ScreenSer  SScreens
	UserSer    SUsers
}

func NewService(db *mongo.Database) *Service {
	return &Service{
		CatalogSer: NewSCataloguesService(db),
		ScreenSer:  NewScreenService(db),
		UserSer:    NewUserService(db),
	}
}
