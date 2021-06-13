package repository

import (
	"context"
	"fmt"
	"github.com/MatthewT4/SchShellGolang/internal/service"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Catalogues interface {
	AddCatalog(ctx context.Context, cat service.Catalog) (result interface{}, err error)
	DelCatalog(ctx context.Context, cat service.Catalog, us User) (int64, error)
	AddDataInCatalog(ctx context.Context, cat service.Catalog, us User, dataName string) (int64, error)
	GetDataInCatalog(ctx context.Context, cat service.Catalog, us User) ([]string, error)
	DelDataInCatalog(ctx context.Context, cat service.Catalog, us User, nameDelData string) (int64, error)
}

type CataloguesRepo struct {
	collection *mongo.Collection
}

func NewCataloguesRepo(db *mongo.Database) *CataloguesRepo {
	return &CataloguesRepo{collection: db.Collection(NameCatalogues)}
}

func (c *CataloguesRepo) AddCatalog(ctx context.Context, cat service.Catalog) (ID interface{}, err error) {
	bs, er := bson.Marshal(cat)
	if er != nil {
		logrus.Debug(er)
		return 0, er
	}
	result, err := c.collection.InsertOne(ctx, bs)
	return result.InsertedID, err
}

//DelCatalog Delete from ID catalogues and Login Holder (in User)
func (c *CataloguesRepo) DelCatalog(ctx context.Context, cat service.Catalog, us User) (int64, error) {
	res, err := c.collection.DeleteOne(ctx, bson.M{"name": cat.Name, "holder": us.Login})
	return res.DeletedCount, err
}

func (c *CataloguesRepo) AddDataInCatalog(ctx context.Context, cat service.Catalog, us User, dataName string) (int64, error) {
	filter := bson.M{"holder": us.Login, "name": cat.Name}
	update := bson.D{
		{"$push", bson.D{
			{"data", dataName},
		}},
	}
	res, err := c.collection.UpdateOne(ctx, filter, update)
	fmt.Println(res) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return res.ModifiedCount, err
}

func (c *CataloguesRepo) GetDataInCatalog(ctx context.Context, cat service.Catalog, us User) ([]string, error) {
	filter := bson.M{"holder": us.Login, "name": cat.Name}
	var CataloguesLimitOnlyGetData struct {
		Data []string `bson:"data"`
	}
	err := c.collection.FindOne(ctx, filter).Decode(&CataloguesLimitOnlyGetData)
	return CataloguesLimitOnlyGetData.Data, err
}

func (c *CataloguesRepo) DelDataInCatalog(ctx context.Context, cat service.Catalog, us User, nameDelData string) (int64, error) {
	filter := bson.M{"holder": us.Login, "name": cat.Name}
	update := bson.D{
		{"$pull", bson.D{
			{"data", nameDelData},
		}},
	}
	res, err := c.collection.UpdateOne(ctx, filter, update)
	fmt.Println(res)
	return res.ModifiedCount, err
}
