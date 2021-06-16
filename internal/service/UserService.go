package service

import (
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Login      string   `bson:"login,omitempty"`
	Password   string   `bson:"password,omitempty"`
	Role       int      `bson:"group,omitempty"`
	Email      string   `bson:"email,omitempty"`
	Catalogues []string `bson:"catalogues,omitempty"`
}

func NewUser(login string, password string, role int, email string, catalogues []string) *User {
	return &User{
		Login:      login,
		Password:   password,
		Role:       role,
		Email:      email,
		Catalogues: catalogues,
	}
}

type UserService struct {
	Us repository.Users
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{repository.NewUserRepo(db)}
}
