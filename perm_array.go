package main

import "fmt"

func perm(head []int, rest []int) {
	if len(rest) == 0 {
		fmt.Println(head)

	} else {
		for i, v := range rest {

			restx := make([]int, len(rest))
			copy(restx, rest)

			restx = append(restx[:i], restx[i+1:]...)
			headx := append(head, v)

			perm(headx, restx)
		}
	}
}

func main() {
	perm([]int{}, []int{1, 2, 3, 4})
}
