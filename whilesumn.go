// while n!=0 sum all input number
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for n != 0 {
		sum += n
		fmt.Scan(&n)
	}
	fmt.Println(sum)
}
