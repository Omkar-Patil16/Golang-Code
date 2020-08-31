//Package mySub provides ACME math Solution
package mySub



//MySub subtracts unlimited no of ints
func Sub(x ...int) int {
	sub := 0
	for _, v := range x {

		sub = v - sub
		if sub < 0 {
			sub = - sub
		}
	}
	return sub
}


