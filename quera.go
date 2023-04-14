package main

import "fmt"

func main() {
	var a, b, c, khar int
	fmt.Scan(&a, &b)
	fmt.Scan(&c)

	for i := 1; i <= c; i++ {
		sum := 0
		khar := (khar * 2) - b

		fmt.Println(khar)

		sum := sum + khar

	}
	fmt.Println(sum)

}
