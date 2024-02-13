package main

import (
	"fmt"
	"sort"
)

/*
Design a system that search prefix into syste and return three words in lexicographical order.
Return the suggested words, that will be list of lists after each character of the searched word is typed.

Example
Suppose n = 5, words = ["carpet", "cart", "car"
"camera", "crate"], and search = "camera".
Search
Prefix					Matches      									Lexicographically Smallest 3
c			["carpet", "cart","car", "camera","crate"]				["camera", "car","carpet"]
ca			['carpet", "cart","car", "camera"]						['camera", "car","carpet"]
cam			['camera"]												['camera"]
came		['camera"]												['camera"]
camer		['camera"]												['camera"]
camera		['camera"] 												['camera"]

Hence the answer is [['camera" "car" "carpet"], ['camera"
"car", "carpet"], ["camera"], ["camera"], ["camera"], ['camera"]].
*/
func main() {
	products := []string{"carpet", "cart", "car", "camera", "crate"}
	search := "camera"
	fmt.Println(getProductSuggestions(products, search))
}

func getProductSuggestions(products []string, search string) [][]string {
	dp := make(map[string][]string)
	n := len(products)
	var p []string
	remo := make(map[string]bool)
	for i := 0; i < n; i++ {
		if !remo[products[i]] {
			p = append(p, products[i])
			remo[products[i]] = true
		}
	}
	for i := 0; i < n; i++ {
		size := len(p[i])
		str := p[i]
		for j := 1; j <= size; j++ {
			dp[str[:j]] = append(dp[str[:j]], str)
		}
	}
	var result [][]string
	for i := 1; i <= len(search); i++ {
		str := search[:i]
		list := dp[str]
		if len(list) == 0 {
			continue
		}
		sort.Strings(list)
		if len(list) >= 3 {
			result = append(result, []string{list[0], list[1], list[2]})
		} else {
			result = append(result, list)
		}
	}
	return result

}
