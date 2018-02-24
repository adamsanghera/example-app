package main

import (
	"fmt"
	"os"
)

import "github.com/adamsanghera/greeter"

func main() {
	var g greeter.Greeter

	if os.Getenv("DEBUG") == "true" {

		g = greeter.NewSimpleGreeter("stormy", "hello I am simple")
	} else {
		g = greeter.NewComplexGreeter(40.6, -73.9)
	}

	// SFG := ComplexGreeter{
	// 	lat: 37.7,
	// 	lon: 122.4,
	// }

	fmt.Println(g.SayHello())
	fmt.Println(g.TellWeather())
}
