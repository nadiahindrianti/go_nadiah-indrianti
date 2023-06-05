package main

import "fmt"

func munculSekali(angka string) []int {
	mapping := make(map[int]int)

	for _, x := range angka {
		num, err := strconv.Atoi(string(x))
		if err != nil {
			continue
		}
		mapping[num]++
	}
	var munculsekali []int 
	for num, y := range mapping {
		if y == 1 {
			munculsekali = append(munculsekali, num)
		}
	}
	
	return munculsekali
}

func main() {
	//Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    //[ 0 8 7 2 5 4]

}
