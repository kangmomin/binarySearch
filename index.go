package main

import "fmt"

func main() {
	var scan int
	_, err := fmt.Scan(&scan)
	if err != nil {
		fmt.Println(err)
		return
	}
	finder(scan, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func finder(x int, list []int) {
	low := 0
	high := len(list) - 1
	mid := len(list) / 2

	for x != list[mid] {
		mid = (low + high) / 2
		if x > list[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println(list[mid])
}
