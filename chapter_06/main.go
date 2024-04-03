package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declaring an array
	// var array_name [size_of_array]type_of_elements_in_array
	var num_array [5]int
	// Array of name: `num_array`, size: `5`, type: `int`
	fmt.Println(num_array)

	// Initialising an array
	num2_array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num2_array)

	// Indexing in arrays: starts from 0
	fmt.Println(num2_array[0])

	num3_array := [...]int{0, 9, 8, 7, 6}
	fmt.Println(num3_array)

	// Get length of an array
	fmt.Println(len(num3_array))

	// Multidimensional Array
	// A 2D array : row, column
	//         r  c
	var table [5][6]string
	for table_row := 0; table_row < 5; table_row++ {
		for table_column := 0; table_column < 6; table_column++ {
			table[table_row][table_column] = strconv.Itoa(table_row) + ", " + strconv.Itoa(table_column)
		}
	}
	fmt.Println(table)

	// A 3D array : row, column, depth
	//        r  c  d
	var cube [4][3][3]string
	// cube_row, cube_column, cube_depth := 0, 0, 0

	for cube_row := 0; cube_row < 4; cube_row++ {
		for cube_column := 0; cube_column < 3; cube_column++ {
			for cube_depth := 0; cube_depth < 3; cube_depth++ {
				cube[cube_row][cube_column][cube_depth] = strconv.Itoa(cube_row) + strconv.Itoa(cube_column) + strconv.Itoa(cube_depth)
			}
		}
	}
	fmt.Println(cube)

}
