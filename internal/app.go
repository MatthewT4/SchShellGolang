package internal

import (
	"github.com/MatthewT4/SchShellGolang/internal/http"
	mongodb "github.com/MatthewT4/SchShellGolang/pkg/mongoDB"
)

/*
import (
	"github.com/Mat/SchShell/internal/db/repository"
	mongodb "github.com/Mat/SchShell/pkg/mongoDB"
	"log"
)*/

type Role int

const (
	Default = iota
	Moderator
	Administrator
)

func Start() {
	client, err := mongodb.NewClient("mongodb+srv://cluster0.lbets.mongodb.net", "Mathew", "8220")
	if err != nil {
		panic(err)
	}
	name := "test"
	db := client.Database(name)
	r := http.NewRouter(db)
	r.Start() /*
		client, err := mongodb.NewClient("mongodb+srv://cluster0.lbets.mongodb.net","Mathew", "8220")
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		name := "test"
		db := client.Database(name)
		colCat := repository.NewCataloguesRepo(db)
		colCat.AddCatalog(context.TODO(), repository.Catalog{Name: "Default", Holder: "Mathew", Data: []string{"logo.jpg", "OK.jpg"}})
		col := repository.NewUserRepo(db)
		use := repository.GetNullUser()
		use.Login = "logffinTffest"
		use.Email = "sdvddd@bffggkf.fflflf"
		err = col.AddUser(context.TODO(), use)
		if err != nil {
			log.Fatal(err)
		}/*err = col.AddUser(context.TODO(), repository.NewUser("1253", "2345", Administrator, "test@google.com"))
		if err != nil {
			//log.Fatal(err, mongodb.IsDuplicate(err))
		}
		fmt.Println("OK 1")
		err = col.CheckUser(context.TODO(), repository.NewUser("123", "2345", Administrator, "test@google.com"))
		if err != nil {
			fmt.Println("I am err", err)
		}
		fmt.Println("OK 2")
		err = col.ReplayUserPassword(context.TODO(), repository.NewUser("123", "2345", Administrator, "test@google.com"), "15423")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("OK 3")
		err = col.RemoveUser(context.TODO(), repository.NewUser("123", "2345", Administrator, "test@google.com"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("OK 4")*/
}
