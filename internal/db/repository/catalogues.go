package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DbCatalog struct {
	Holder *string   `bson:"holder,omitempty"`
	Data   *[]string `bson:"data"`
	Name   *string   `bson:"name,omitempty"`
	Type   *int      `bson:"type,omitempty"`
}

func NewDbCatalog(holder *string, data *[]string, name *string, dtype *int) *DbCatalog {
	if data == nil {
		var dd []string
		dd = append(dd, "logo.jpg")
		data = &dd
	}
	return &DbCatalog{
		Holder: holder,
		Data:   data,
		Name:   name,
		Type:   dtype,
	}
}

type Catalogues interface {
	AddCatalog(ctx context.Context, cat *DbCatalog) (ID interface{}, err error)
	DelCatalog(ctx context.Context, catName string, holder string) (int64, error)
	GetCatalogs(ctx context.Context, holder string) ([]string, error)
	AddDataInCatalog(ctx context.Context, holder string, CatalogName string, dataName string) (int64, error)
	GetDataInCatalog(ctx context.Context, holder string, catalogName string) ([]string, error)
	DelDataInCatalog(ctx context.Context, holder string, catalogName string, nameDelData string) (int64, error)
}

type CataloguesRepo struct {
	collection *mongo.Collection
}

func NewCataloguesRepo(db *mongo.Database) *CataloguesRepo {
	return &CataloguesRepo{collection: db.Collection(NameCatalogues)}
}

func (c *CataloguesRepo) AddCatalog(ctx context.Context, cat *DbCatalog) (ID interface{}, err error) {
	bs, er := bson.Marshal(cat)
	if er != nil {
		logrus.Debug(er)
		return 0, er
	}
	result, err := c.collection.InsertOne(ctx, bs)
	if err != nil {
		return 0, err
	}
	return result.InsertedID, err
}

func (c *CataloguesRepo) DelCatalog(ctx context.Context, catName string, holder string) (int64, error) {
	res, err := c.collection.DeleteOne(ctx, bson.M{"name": catName, "holder": holder})
	return res.DeletedCount, err
}

func (c *CataloguesRepo) GetCatalogs(ctx context.Context, holder string) ([]string, error) {
	filter := bson.M{"holder": holder}
	res, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	type OnlyGetCatalogs struct {
		Name string `bson:"name"`
	}
	var data []string
	for res.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem OnlyGetCatalogs
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, elem.Name)
	}
	return data, err
}

func (c *CataloguesRepo) AddDataInCatalog(ctx context.Context, holder string, CatalogName string, dataName string) (int64, error) {
	filter := bson.M{"holder": holder, "name": CatalogName}
	update := bson.D{
		{"$push", bson.D{
			{"data", dataName},
		}},
	}
	res, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	fmt.Println(res.ModifiedCount, res.UpsertedCount, res.MatchedCount) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return res.ModifiedCount, err
}

func (c *CataloguesRepo) GetDataInCatalog(ctx context.Context, holder string, catalogName string) ([]string, error) {
	filter := bson.M{"holder": holder, "name": catalogName}
	var CataloguesLimitOnlyGetData struct {
		Data []string `bson:"data"`
	}
	err := c.collection.FindOne(ctx, filter).Decode(&CataloguesLimitOnlyGetData)
	return CataloguesLimitOnlyGetData.Data, err
}

func (c *CataloguesRepo) DelDataInCatalog(ctx context.Context, holder string, catalogName string, nameDelData string) (int64, error) {
	filter := bson.M{"holder": holder, "name": catalogName}
	update := bson.D{
		{"$pull", bson.D{
			{"data", nameDelData},
		}},
	}
	res, err := c.collection.UpdateOne(ctx, filter, update)
	fmt.Println(res)
	return res.ModifiedCount, err
}
