//Ultra mean the best package
package ultra

// Sum function adds any given number of ints
func Sum(n ...int) int{
	s := 0
	for _,v := range n{
		 s += v
	}
	return  s
}
