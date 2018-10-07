/* https://tour.golang.org/moretypes/23
Implement WordCount, which returns a map of the counts of each word in the string s
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	count := make(map[string]int)

	for _, val := range words {
		if _, ok := count[val]; ok {
			count[val] += 1
		} else {
			count[val] = 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
