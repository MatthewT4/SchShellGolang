package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type TypeCatalog int
const (
	Image = iota
	Audio
	Video
)

type Catalog struct {
	Holder string 					`bson:"holder,omitempty"`
	Data []string					`bson:"data,omitempty"`
	Name string 					`bson:"name,omitempty"`
	Type TypeCatalog				`bson:"type,omitempty"`
}

type CataloguesRepo struct {
	collection *mongo.Collection
}
func NewCataloguesRepo(db *mongo.Database) *CataloguesRepo {
	return &CataloguesRepo{collection: db.Collection(NameCatalogues)}
}

func (c *CataloguesRepo) AddCatalog(ctx context.Context, cat Catalog) (result *mongo.InsertOneResult, err error) {
	bs, er := bson.Marshal(cat)
	if er != nil {
		logrus.Debug(er)
		return result, er
	}
	result, err = c.collection.InsertOne(ctx, bs)
	return
}

//DelCatalog Delete from ID catalogues and Login Holder (in User)
func (c *CataloguesRepo) DelCatalog(ctx context.Context, cat Catalog, us User) (int64, error) {
	res, err := c.collection.DeleteOne(ctx, bson.M{"name":cat.Name, "holder": us.Login})
	return res.DeletedCount, err
}
func (c *CataloguesRepo) AddDataInCatalog(ctx context.Context, cat Catalog, us User, DataName string) error {
	filter := bson.M{"holder": us.Login, "name": cat.Name}
	update := bson.D{
		{"$push", bson.D{
			{"data", DataName},
		}},
	}
	res, err := c.collection.UpdateOne(ctx, filter, update)
	fmt.Println(res)
	return err
}

func (c *CataloguesRepo) GetDataInCatalog(ctx context.Context, cat Catalog, us User) ([]string, error) {
	filter := bson.M{"holder": us.Login, "name": cat.Name}
	var CataloguesLimitOnlyGetData struct {
		Data []string		`bson:"data"`
	}
	err := c.collection.FindOne(ctx, filter).Decode(&CataloguesLimitOnlyGetData)
	return CataloguesLimitOnlyGetData.Data, err
}