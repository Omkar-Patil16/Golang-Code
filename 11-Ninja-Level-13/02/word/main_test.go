package word

import (
	"github.com/Omkar-Patil16/Golang-Code/11-Ninja-Level-13/02/quote"
	"testing"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("expected :", 3, "got :",n)
	}
}

func TestUseCount(t *testing.T) {
	for k,v := range(UseCount("one two two three three three")) {
			switch k {
			case "one" :
				if v != 1 {t.Error("expected :", 1, "got :", v)}
			case "two" :
				if v != 2 {t.Error("expected :", 2, "got :", v)}
			case "three" :
				if v != 3 {t.Error("expected :", 3, "got :", v)}
			}
	}
}


func BenchmarkCount(b *testing.B) {
	for i:=0;i<b.N;i++{
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i:=0;i<b.N;i++{
		UseCount(quote.SunAlso)
	}
}