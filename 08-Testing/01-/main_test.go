package main

import "testing"

func TestMySub(t *testing.T) {
	x := mySub(9,4)
	if x != 5 {
		t.Error("Expected 5 got ", x)
	}
}