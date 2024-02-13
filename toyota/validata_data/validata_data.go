package main

/*

Write a program to validate the first sheet using the summary data from the second sheet.

Two sheets in CSV format given below like this
1. year : year when forest fires done
2. state : Brazilian state
3. month : month when forest fires done
4. number : number of reported fires in that year, state, and month

The second sheet has the following columns:
1. year: year when forest fires done
2. number : total number of reported fires in that year*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr1 := []string{"year,state,month,number",
		"2001,sdfsddf,sdfs,20",
		"2002,sdfsddf,sdfs,10",
		"2005,sdfsddf,sdfs,30",
		"2001,sdfsddf,sdfs,2",
		"2005,sdfsddf,sdfs,2"}
	arr2 := []string{
		"year,number",
		"2001, 22",
		"2002, 10",
		"2005, 32",
	}
	//first := true
	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "\n" {
			first = false
		}
		if first {
			arr1 = append(arr1, data)
		} else {
			arr2 = append(arr2, data)
		}
	}*/
	//	fmt.Println(arr1)
	//	fmt.Println(arr2)
	m := make(map[string]int)

	for i := 0; i < len(arr2); i++ {
		if i == 0 {
			continue
		}
		str := strings.Split(arr2[i], ",")
		num, err := strconv.Atoi(strings.TrimSpace(str[1]))
		if err != nil {
			fmt.Println(err)
		}
		m[str[0]] = num
	}
	for i := 0; i < len(arr1); i++ {
		if i == 0 {
			continue
		}
		str := strings.Split(arr1[i], ",")
		num, _ := strconv.Atoi(str[3])
		v, ok := m[str[0]]
		if !ok {
			fmt.Println(false)
			return
		} else {
			v = v - num
			if v < 0 {
				fmt.Println(false)
				return
			} else if v == 0 {
				delete(m, str[0])
			} else {
				m[str[0]] = v
			}
		}
	}
	if len(m) == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
