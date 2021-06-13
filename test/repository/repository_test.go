package repository

import (
	"context"
	"github.com/MatthewT4/SchShellGolang/internal"
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"github.com/MatthewT4/SchShellGolang/internal/service"
	mongodb "github.com/MatthewT4/SchShellGolang/pkg/mongoDB"
	"testing"
)

func TestAddAndFindCatalogues(t *testing.T) {
	client, err := mongodb.NewClient("mongodb+srv://cluster0.lbets.mongodb.net", "Mathew", "8220")
	if err != nil {
		t.Fatal(err)
	}
	name := "test"
	db := client.Database(name)
	colCat := repository.NewCataloguesRepo(db)
	us := repository.User{Login: "Mathew", Password: "12345", Role: internal.Administrator}
	data := []string{"logo.jpg", "OK.jpg"}
	cat := service.Catalog{Name: "Default", Holder: us.Login, Data: data, Type: service.Image}
	{
		result, err := colCat.AddCatalog(context.TODO(), cat)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
	{
		res, err := colCat.GetDataInCatalog(context.TODO(), cat, us)
		if err != nil {
			t.Fatal(err)
		}
		if len(res) != len(data) {
			t.Fatal("len result != len data")
		}
		for i := 0; i < len(data); i++ {
			t.Log(i)
			if res[i] != data[i] {
				t.Error("result GetDataInCatalog != valid result")
			}
		}
	}
	{
		delCount, err := colCat.DelCatalog(context.TODO(), cat, us)
		if delCount != 1 {
			t.Fatalf("Result delete != 1, error: %v", err)
		}
		if err != nil {
			t.Fatal(err)
		}
		_, err = colCat.GetDataInCatalog(context.TODO(), cat, us)
		if err.Error() != "mongo: no documents in result" {
			t.Fatal(err)
		}
	}
	client.Disconnect(context.TODO())
}

/*
func TestAddAndDelData(t *testing.T) {
	client, err := mongodb.NewClient("mongodb+srv://cluster0.lbets.mongodb.net", "Mathew", "8220")
	if err != nil {
		t.Fatal(err)
	}
	name := "test"
	db := client.Database(name)
	colCat := repository.NewCataloguesRepo(db)
	us := repository.User{Login: "Mathew", Password: "12345", Role: internal.Administrator}
	data := []string{"logo.jpg", "OK.jpg"}

}*/
