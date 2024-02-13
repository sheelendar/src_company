package main

import (
	"fmt"
	"sort"
)

/*
Twin Strings:
Two strings are twins if they are equivalent by doing some number of operations on one or
both strings. There are only 2 operation.

• Swap_Even: swap a char at even number index with a character at another even number index.
• Swap_Odd: swap a character at odd number index with a character at another odd number index.

There are two string array. compare a string from each array,of same index and save result into a result
array like "Yes" or "No" by comparing whether the strings compared are twins or not.


Examples
arr1 = [abcd , abcd]
arr2 = [cbad, adbc]

• Compare two strings arr1[0] and arr2[0]
	One Swap Even operation allows us to swap the characters "a" and "c" of the string "abcd" and make
	it equivalent to "bad" ("abcd" -› "cbad"). arr1[0] and arr2[0] are twins and the answer is "Yes".

• Compare the two strings arr1 [1] and arr2[1]
	Swap Odd or Swap Even operations can make the two strings abed and adbc equivalent.
	arr1 [1] and arr2[1] are not twins and the answer is ["No"].

result: ["Yes,"No"].

size n = 3,  arr1 = [cdab, dcba, abed]
size n = 3,  arr2 = [abcd, abcd,abcdcd]
Sample Output 0
result: [Yes, No, No]
*/

func main() {

	/*	firstString := []string{"cdab", "dcba", "abed"}
		secondString := []string{"abcd", "abcd", "abcdcd"}
		fmt.Println(twins(firstString, secondString))
	*/

	/*firstString := []string{"abcd", "abcd"}
	secondString := []string{"cbad", "adbc"}
	fmt.Println(twins(firstString, secondString))
	*/

	firstString := []string{"abcd", "abcd"}
	secondString := []string{"cbad", "adbc"}
	fmt.Println(twins(firstString, secondString))

}

func twins(firstStrings, sedongStrings []string) []string {
	var result []string
	size := len(firstStrings)

	for i := 0; i < size; i++ {
		m := len(firstStrings[i])
		n := len(sedongStrings[i])
		if n != m {
			result = append(result, "No")
		} else {
			str1 := firstStrings[i]
			str2 := sedongStrings[i]
			var fe, fo, se, so string

			for j := 0; j < m; j++ {
				if (j+1)%2 == 1 {
					fo += string(str1[j])
					so += string(str2[j])
				} else {
					fe += string(str1[j])
					se += string(str2[j])
				}
			}
			arr := []byte(fe)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			fe = string(arr)

			arr = []byte(se)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			se = string(arr)

			arr = []byte(fo)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			fo = string(arr)

			arr = []byte(so)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			so = string(arr)

			if fe == se && fo == so {
				result = append(result, "Yes")
			} else {
				result = append(result, "No")
			}
		}
	}
	return result
}

//fe :=db- > bd
//fo:=ca -> ac
//se:= bd -> bd
//so:=ac -> ac
