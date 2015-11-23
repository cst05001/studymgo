package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func mapreduce(collection *mgo.Collection) {
	fmt.Println("mapreduceStart")
	job := &mgo.MapReduce{
		Map: `
function() {
	emit(this.type, {count: 1});
}
`,
		Reduce: `
function(key, values) {
	var sum = 0;
	values.forEach(function(item) {
		sum += item.count;
	});
	return {count: sum};
}
`,
	}

	result := make([]MapReduceResult, 0)
	_, err := collection.Find(nil).MapReduce(job, &result)
	if err != nil {
		panic(err)
	}
	for _, item := range result {
		fmt.Printf("_id: %s\tvalue: %d\n", item.Id, item.Value)
	}
	fmt.Println("mapreduceEnd")
}
