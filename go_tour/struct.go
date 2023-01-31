package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := &Vertex{1, 2}
	fmt.Println(*a)
	fmt.Printf("a.X = %d,%d", a.X, a.Y)
}
