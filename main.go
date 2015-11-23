package main

import (
	"flag"
	"gopkg.in/mgo.v2"
)

func main() {
	mongodAddr := flag.String("addr", "127.0.0.1:27017", "-addr=127.0.0.1:27017")
	switchFind := flag.Bool("find", false, "-find")
	switchMapReduce := flag.Bool("mapreduce", false, "-mapreduce")
	switchInsert := flag.Bool("insert", false, "-insert")
	switchInit := flag.Bool("init", false, "-init")
	flag.Parse()

	session, err := mgo.Dial(*mongodAddr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	dbName := "study"
	collectionName := "test1"
	collection := session.DB(dbName).C(collectionName)

	if *switchFind {
		find(collection)
	} else if *switchMapReduce {
		mapreduce(collection)
	} else if *switchInsert {
		insert(collection)
	} else if *switchInit {
		initCollection(collection, 10000)
	} else {
		flag.PrintDefaults()
	}
}
