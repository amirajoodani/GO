// average of two decimal number 3 and 6
package main

import "fmt"

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func main() {
	fmt.Println(avg(3, 6))
}
