package main

import "fmt"

func main() {

	var map1 map[string]int
	map1 = make(map[string]int)

	printMap(map1)

	map2 := map[string]int{"prov": 1, "other": 2}
	printMap(map2)

	fmt.Println(map2["prov"])

	map2["prov"] = 5
	fmt.Println(map2["prov"])

	delete(map2, "prov")
	_, isPresent := map2["prov"]
	fmt.Printf("is present? %t", isPresent)
}

func printMap(myMap map[string]int) {
	if len(myMap) == 0 {
		fmt.Println("empty")
	}
	for key, val := range myMap {
		fmt.Printf("%s=%d\n", key, val)
	}
}
