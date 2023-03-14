package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	"GG":        {75.12345, -123.45678},
}

func main() {
	fmt.Println(m)
}
