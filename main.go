package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := 0

	temp = *a
	*a = *b
	*b = temp

}

func main() {

	//top k element problem solution
	p := []int{32, 23, 50, 99, 20, 10, 56}
	var k int

	fmt.Scanf("%d", &k)
	counter := 0
	for i := 0; i < len(p); i++ {

		if i+1 >= len(p) {
			i = 0
			counter++
			if counter >= len(p) {
				break
			}
		}
		if p[i] > p[i+1] {
			swap(&p[i], &p[i+1])

		}

	}

	for j := len(p) - 1; j >= len(p)-k; j-- {
		fmt.Printf("%d ", p[j])
	}
}
