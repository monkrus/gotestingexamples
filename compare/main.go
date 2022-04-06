package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []string{"strawberry", "raspberry"}
	b := []string{"strawberry", "raspberry"}

	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(a, b))

	b = append(b, "test")
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(a, b))
	testEq(a, b)
}
func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	fmt.Println("test")
	return true
}
