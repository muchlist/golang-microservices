package main

type Product struct {
	Name string
	Len  int //milimeters
	Wid  int
	Hei  int
}

type Box struct {
	Len int //milimeters
	Wid int
	Hei int
}

func getBestBox(availableBoxes []Box, products []Product) Box {
	return Box{}
}
