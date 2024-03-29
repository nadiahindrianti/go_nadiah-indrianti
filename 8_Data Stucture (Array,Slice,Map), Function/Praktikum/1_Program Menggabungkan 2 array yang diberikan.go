package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	ArrayMerge := append(arrayA, arrayB...)
	result := make ([]string, 0)
	arrayAandB := make(map[string]bool)

	for _, u := range arrayMerge {
		if !arrayAandB[u] {
			arrayAandB[u] = true
			result = append(result, u)
		}
	}

	return result
	
}

func main() {
	//Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin","steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu","devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// [ "devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// [ "hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []

}
