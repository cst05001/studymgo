package main

type Test1 struct {
	Name  string
	Type  string
	Price float64
}

type MapReduceResult struct {
	Id    string "_id"
	Value struct {
		Count int64 "count"
	} "value"
}
