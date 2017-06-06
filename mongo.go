package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
	"time"
)

type Product struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Description     string
	Timestamp time.Time
}

var (
	IsDrop = true
	db = "go"
	collection = "product"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)


	c := session.DB(db).C(collection)

	index := mgo.Index{
		Key:        []string{"title", "description"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	err = c.Insert(&Product{Title: "Echo", Description: "Amazon Echo Description", Timestamp: time.Now()},
		&Product{Title: "Iphone 7", Description: "Apple IPhone 7", Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}

	result := Product{}
  //db.product.findOne({"tiltle":"Echo"})
	err = c.Find(bson.M{"title": "Echo"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data", result)

	//Retrieve all documents
	var results []Product
	err = c.Find(bson.M{"title": "Echo"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("All documents: ", results)




}
