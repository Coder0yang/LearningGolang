package main

import (
	"fmt"
)

func InsertSort(a []int) {
	var j int
    var temp int
	for i := 1; i < len(a); i++ {
		temp = a[i]
		j = i - 1
		for j = i - 1; j >= 0 && temp < a[j]; j-- {
			a[j+1] = a[j]
		}
        a[j+1] = temp
	}
}

func Show(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t", a[i])
	}
	fmt.Println()
}

func main() {
	a := []int{4, 5, 2, 8, 9, 6, 3, 1, 7}
	Show(a)
	InsertSort(a)
	Show(a)
}
