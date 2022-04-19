// supports only v1.18 Go

package main

import (
	"testing"
)

func TestTestEq(t *testing.T) {
	a := []string{"strawberry", "raspberry"}
	b := []string{"strawberry", "raspberry"}

	for i := 0; i < len(a); i++ {
		if !testEq(a, b) {
			t.Fail()
		}
	}

}
