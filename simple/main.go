package main

import "fmt"

func main() {
	fmt.Println(Add(0))
}

func Multi(x int, y int) int {
	return x * y
}
func Add(num int) int {
	//num += 2
	return num + 2
}
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}
