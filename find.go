package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func find(collection *mgo.Collection) {
	fmt.Println("findStart")
	iter := collection.Find(nil).Iter()
	test1 := &Test1{}
	for iter.Next(test1) {
		fmt.Printf("Name: %s\tPrice: %f Type: %s\n", test1.Name, test1.Price, test1.Type)
	}
	fmt.Println("findEnd")
}
