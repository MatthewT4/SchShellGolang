package main

import (
	"github.com/MatthewT4/SchShellGolang/internal"
)

func main() {
	internal.Start()
	/*server.Test("ss")
	var db data.DataBase = data.NewData()
	db.Add("ddd")e
	db.Print()
	err := ttt.Prt("aaa")
	if err == nil {
		fmt.Println("nil OK!")
	}
	/*client, err := mongoDB2.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	for i := 0; i < 1000; i++ {
		res, err := collection.InsertOne(ctx, bson.M{"name": "User", "value": "Anton"})
		if err != nil {
			panic(fmt.Errorf("I panic %v", err))
		}
		id := res.InsertedID
		fmt.Println(i, id)

	}
	internal.Start()
	conf := server.Config{Addr: ":8080"}
	server.Start(conf)*/
}
