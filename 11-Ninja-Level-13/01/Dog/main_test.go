package Dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i:=0;i<b.N;i++{
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i:=0;i<b.N;i++{
		YearsTwo(20)
	}
}

func ExampleYears() {
	fmt.Println(Years(10));
	//Output:
	//70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	//Output:
	//140
}
func TestYearsTwo(t *testing.T) {
	xn := YearsTwo(20)
	if xn!=140{
		fmt.Println("got ", xn ,"expected " , 140)
	}
}

func TestYears(t *testing.T) {
	xn := Years(10)
	if xn!=70{
		fmt.Println("got ", xn ,"expected " , 70)

	}

}