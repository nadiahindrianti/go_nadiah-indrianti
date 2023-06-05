package main

import "fmt"

func PairSum(arr []int, target int) [] int {
	x, y := 0, len(arr)-1
	// array x berada pada bagian kiri
	// array y berada pada bagian kanan

	for x < y {
		sum := arr[x] + arr[y]
		if sum == target {
			return []int{x, y}
		} else if sum < target {
			x++
		} else {
			y--
		}
	}
	return []int{}
}


func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1,3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0,2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2,3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1,2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0,1]

}
