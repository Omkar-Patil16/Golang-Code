package Bookkeeping

import (
	"fmt"
	"testing"
)

func TestDiary(t *testing.T) {
	s := Diary("2020")

	if s!="Engineering Diary 2020"{
		t.Error("Expected ","Engineering Diary 2020","Got ", s)
	}
}

func ExampleDiary() {
	fmt.Println(Diary("2020"))
	//Output:
	//Engineering Diary 2020
}

func BenchmarkDiary(b *testing.B) {
	for i:=0;i<b.N;i++{
		Diary("2020")
	}
}