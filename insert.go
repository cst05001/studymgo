package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func insert(collection *mgo.Collection) {
	fmt.Println("insertStart")
	test1 := &Test1{
		Name:  "iPhone 6s 128G",
		Price: 7888,
		Type:  "phone",
	}
	err := collection.Insert(test1)
	if err != nil {
		panic(err)
	}
	fmt.Println("insertEnd")
}
