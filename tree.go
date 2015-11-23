package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func children(db *mgo.Database, categoryName string, level int) {
	collectionCategories := db.C("categories")
	query := collectionCategories.Find(bson.M{"_id": categoryName})
	category := &Category{}
	err := query.One(category)
	if err != nil {
		panic(err)
	}

	for i := 0; i < level; i = i + 1 {
		fmt.Print("\t")
	}
	fmt.Println(categoryName)

	for _, child := range category.Children {
		children(db, child, level+1)
	}
}

func tree(db *mgo.Database) {
	collectionCategories := db.C("categories_root")
	iter := collectionCategories.Find(nil).Iter()
	root := &Category{}
	for iter.Next(root) {
		children(db, root.Id, 0)
	}
}
