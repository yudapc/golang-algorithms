package sliceofmapstring

import "fmt"

func SliceOfMapString() {
	fmt.Println("Hello, playground this is sample slice of map")
	users := []map[string]string{
		{
			"id":   "1",
			"name": "Cogati",
			"age":  "30",
		}, {
			"id":   "2",
			"name": "Prabu",
			"age":  "14",
		}, {
			"id":   "3",
			"name": "Yuda",
		},
	}

	for _, val := range users {
		_, found := val["age"]
		if found {
			fmt.Println("Found with id", val["id"])
		}
	}
}
