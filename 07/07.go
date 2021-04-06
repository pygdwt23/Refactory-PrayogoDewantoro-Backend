package main

import (
	"fmt"
	"sort"
)

func main(){
	numbers := []int{9, 4, 2, 4, 1, 5, 3, 0}
	sort.Sort(sort.IntSlice(numbers))
	fmt.Println(numbers)
}