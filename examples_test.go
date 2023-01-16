package tsyncmap

import "fmt"

func ExampleMap_Load() {
	myMap := Map[string, string]{}
	myMap.Store("a", "b")

	result, ok := myMap.Load("a")
	fmt.Println(result, ok)
}

func ExampleMap_Store() {
	myMap := Map[string, string]{}
	myMap.Store("a", "b")

	result, ok := myMap.Load("a")
	fmt.Println(result, ok)
}
