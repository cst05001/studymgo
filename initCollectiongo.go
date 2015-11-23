package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"math/rand"
	"time"
)

func generateString(length int) string {
	rand.Seed(time.Now().UnixNano())
	stringSet := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	result := make([]rune, length)
	for k, _ := range result {
		result[k] = stringSet[rand.Intn(length)]
	}
	return string(result)
}

func initCollection(collection *mgo.Collection, total int) {
	fmt.Println("initCollectionStart")
	rand.Seed(time.Now().UnixNano())
	err := collection.DropCollection()
	if err != nil {
		panic(err)
	}

	//随机生成 type 列表
	typeCount := rand.Intn(5) + 3
	typeList := make([]string, typeCount)
	for k, _ := range typeList {
		typeList[k] = fmt.Sprintf("type_%s", generateString(4))
	}

	//生成 document
	test1 := &Test1{}
	for i := 0; i < total; i = i + 1 {
		test1.Type = typeList[rand.Intn(typeCount)]
		test1.Name = fmt.Sprintf("name_%s", generateString(10))
		test1.Price = float64(int(rand.Float64()*100000*100)) / 100
		err = collection.Insert(test1)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("initCollectionEnd")
}
