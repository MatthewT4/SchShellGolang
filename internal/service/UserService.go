package service

import (
	"context"
	"errors"
	"github.com/MatthewT4/SchShellGolang/internal/auth"
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const SignKey = "djskffdlsvshjsjdcfnv"

type RoleType int

const (
	UserRole = iota
	Moderation
	Administration
)
const timerJWT = 15 * time.Minute

type User struct {
	Login      string   `bson:"login,omitempty"`
	Password   string   `bson:"password,omitempty"`
	Role       RoleType `bson:"group,omitempty"`
	Email      string   `bson:"email,omitempty"`
	Catalogues []string `bson:"catalogues,omitempty"`
}

func NewUser(login string, password string, role RoleType, email string, catalogues []string) *User {
	return &User{
		Login:      login,
		Password:   password,
		Role:       role,
		Email:      email,
		Catalogues: catalogues,
	}
}

type SUsers interface {
	AddUser(user User) (int, error)
	Authorization(login string, password string) (string, time.Time, error)
	Authentication(token string) (int, string, error)
}

type UserService struct {
	Us       repository.Users
	TokenSer auth.TokenManager
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		Us: repository.NewUserRepo(db),
		TokenSer: func(SignKey string) *auth.Manager {
			res, err := auth.NewManager(SignKey)
			if err != nil {
				log.Fatal(err)
			}
			return res
		}(SignKey),
	}
}

func (u *UserService) AddUser(user User) (int, error) {
	_, err := u.Us.AddUser(context.TODO(), repository.NewDbUser(&user.Login, &user.Password, int(user.Role), &user.Email, &user.Catalogues))
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return 409, errors.New("Duplicate")
		}
		return 500, errors.New("Internal Server Error")
	}
	return 200, nil
}
func (u *UserService) Authorization(login string, password string) (string, time.Time, error) {
	group, err := u.Us.CheckUser(context.TODO(), login, password)
	if err != nil {
		return "404", time.Time{}, errors.New("user not Found")
	}
	times := time.Now().Add(timerJWT)
	res, er := u.TokenSer.NewJWT(login, group, timerJWT)
	if er != nil {
		return "500", time.Time{}, er
	}
	return res, times, nil
}

func (u *UserService) Authentication(token string) (int, string, error) {
	return u.TokenSer.Parse(token)
}
