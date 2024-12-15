package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var n = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(n["Bell Labs"], n["Google"])
	fmt.Println(n)
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	l := make(map[string]int)
	l["key"] = 42
	fmt.Println(l["key"])
	l["key"] = 48
	fmt.Println(l["key"])
	delete(l, "key")
	fmt.Println(l["key"])
	v, ok := l["key"]
	fmt.Println("The value:", v, "Present?", ok)
}
