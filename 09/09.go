package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayList := []int{3, 12, 4, 5, 8, 9}
	sort.Sort(sort.IntSlice(arrayList))
	fmt.Println(arrayList)
}