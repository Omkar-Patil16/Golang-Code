//Package word is a word package
package word

import "strings"


//Function UseCount counts no of time EACH WORD APPEARS and RETURNS THEM
func UseCount(s string) (map[string]int) {
	xs := strings.Fields(s)

	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}
//Function Count counts and returns TOTAL no of WORDS
func Count(s string) int {

	xs := strings.Fields(s)
	return len(xs)

}
