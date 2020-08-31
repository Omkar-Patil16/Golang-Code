package mySub

import (
	"fmt"
	"testing"
)

func TestMySub(t *testing.T) {
	type test struct {
		input []int
		val   int
	}

	tests := []test{
		{[]int{9, 4, 1}, 4},
		{[]int{8, 5}, 3},
		{[]int{7, 3, 4}, 0},
		{[]int{6, 4}, 2},
	}

	for _, v := range tests {
		x := Sub(v.input...)
		if x != v.val {
			t.Error("Expected", v.val, "got", x)
		}

	}

}
func ExampleSub() {
	fmt.Println(Sub(23,10))
	//Output:
	//13
}
