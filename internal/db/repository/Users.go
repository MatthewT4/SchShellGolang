package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	AddUser(ctx context.Context, user *DbUser) error
	CheckUser(ctx context.Context, login *string, password *string) error
	ReplayUserPassword(ctx context.Context, login *string, password *string, newPassword *string) error
}

type DbUser struct {
	Login      *string   `bson:"login,omitempty"`
	Password   *string   `bson:"password,omitempty"`
	Role       *int      `bson:"group,omitempty"`
	Email      *string   `bson:"email,omitempty"`
	Catalogues *[]string `bson:"catalogues,omitempty"`
}

type UserRepo struct {
	collection *mongo.Collection
}

func NewDbUser(login *string, password *string, role *int, email *string, catalogues *[]string) *DbUser {
	return &DbUser{
		Login:      login,
		Password:   password,
		Role:       role,
		Email:      email,
		Catalogues: catalogues,
	}
}
func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{collection: db.Collection(NameUserCollection)}
}

func (u *UserRepo) AddUser(ctx context.Context, user *DbUser) error {
	//_, err := u.collection.InsertOne(ctx, bson.M{"_id": user.Login, "password": user.password, "group": user.role})
	bs, er := bson.Marshal(user)
	if er != nil {
		logrus.Debug(er)
		return er
	}
	//bs := bson.M{"login": login, "password"}
	_, err := u.collection.InsertOne(ctx, bs)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (u *UserRepo) CheckUser(ctx context.Context, login *string, password *string) error {
	err := u.collection.FindOne(ctx, bson.M{"login": login, "password": password}).Err()
	return err
}

func (u *UserRepo) ReplayUserPassword(ctx context.Context, login *string, password *string, newPassword *string) error {
	//filter := bson.D{{"name", user.Login}, {"password", user.Password}}
	filter := bson.M{"login": login, "password": password}
	update := bson.D{
		{"$set", bson.D{
			{"password", newPassword},
		}},
	}
	updResult, err := u.collection.UpdateOne(ctx, filter, update)
	fmt.Println(updResult)
	return err
}

//RemoveUser Delete user by login
func (u *UserRepo) RemoveUser(ctx context.Context, login *string) error {
	deleteResult, err := u.collection.DeleteOne(ctx, bson.M{"login": login})
	fmt.Println(deleteResult)
	return err
}

func (u *UserRepo) SetRole(ctx context.Context, login *string, password *string, newRole int) error {
	filter := bson.M{"login": login, "password": password}
	update := bson.D{
		{"$set", bson.D{
			{"group", newRole},
		}},
	}
	updResult, err := u.collection.UpdateOne(ctx, filter, update)
	fmt.Println(updResult)
	return err
}

/*
func (m *Mongo) OneInsert(Data interface{}, settings Settings) (string, error) {
	collection := m.client.Database(settings.DataBaseName).Collection(settings.CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, Data)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	id := res.InsertedID
}*/
