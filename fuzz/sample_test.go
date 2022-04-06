package main

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}) {
		t.Error("expected true, got false")
	}
}

func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}

func FuzzMult1(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int, y int) {
		Mult(x, y)
	})
}

// `go test` or `go test -v`
//  go test -run=TestName

func FuzzMult(f *testing.F) {
	f.Add(10, 20)
	f.Fuzz(func(t *testing.T, x int, y int) {
		Mult(x, y)
	})

}
