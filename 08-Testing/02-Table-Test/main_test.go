package main

import "testing"

func TestMySub(t *testing.T) {
	type test struct{
		input []int
		val int
	}

	tests := []test{
		test{[]int{9,4,1},4},
		test{[]int{8,5},3},
		test{[]int{7,3,4},0},
		test{[]int{6,4},2},

	}

	for _,v := range tests{
		x := mySub(v.input...)
		if x!=v.val {
			t.Error("Expected",v.val,"got",x)
		}

	}


}