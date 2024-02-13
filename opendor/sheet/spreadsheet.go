package main

import (
	"fmt"
	"strings"
)

/*Prompt:
make spreadsheet application like Google sheets. Our spreadsheet, it will only support column which hold either integers values or operations that add two column.

You have to write a program that build this functionality for spreadsheet. you need to check if column has operation then you need to add solve this opreations.

Example how this will work
set_cell("C1", "45")
set_cell("B1", "10")
set_cell("A1", "=C1+B1")

you will be given one input that is the column location.

Example how your getter will work
get_cell("A1") # should return 55 in this case*/

var spreadsheet = make(map[string]int)

func main() {

	set_cell("C1", 45)
	set_cell("B1", 10)
	set_cell("A1", "=C1+B1")
	fmt.Println(get_cell("A1")) //55
	set_cell("B1", 20)
	set_cell("A1", "=C1+B1")
	fmt.Println(get_cell("A1")) // 65
}

func get_cell(col string) int {
	if v, ok := spreadsheet[col]; ok {
		return v
	}
	return 0
}

func set_cell(col string, input interface{}) {
	if value, ok := input.(int); ok {
		spreadsheet[col] = value
	} else if op, ok := input.(string); ok {
		operand := strings.Split(op, "+")
		cel2 := operand[1]
		operandWithExtra := strings.Split(operand[0], "=")
		cel1 := operandWithExtra[1]
		calculateOperation(col, cel1, cel2)
		/*fmt.Println(cel1)
		fmt.Println(cel2)
		fmt.Println(spreadsheet[col])*/
	}

}

func calculateOperation(col string, cel1 string, cel2 string) {
	if len(spreadsheet) <= 0 {
		return
	}
	val1 := 0
	val2 := 0

	if v, ok := spreadsheet[cel1]; ok {
		val1 = v
	}

	if v, ok := spreadsheet[cel2]; ok {
		val2 = v
	}

	spreadsheet[col] = val1 + val2
}
