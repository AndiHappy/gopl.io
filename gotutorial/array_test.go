package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	var array2 [5]int
	var strArray [5]string
	fmt.Println(reflect.ValueOf(array2).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())

	var theArray [3]string
	theArray[0] = "India"    // Assign a value to the first element
	theArray[1] = "Canada"   // Assign a value to the second element
	theArray[2] = "Japan"    // Assign a value to the third element
	fmt.Println(theArray[0]) // Access the first element value
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2]) // Access the third element valu

	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment
	fmt.Println(x)
	fmt.Println(y)

	x2 := [...]int{10, 20, 30}
	fmt.Println(reflect.ValueOf(x2).Kind())
	fmt.Println(len(x2))

	// 比较新颖的初始化的内容
	x3 := [5]int{1: 10, 3: 30}
	fmt.Println(x3)

	//
	array2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}
	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range array2 {
		fmt.Println(index, "=>", element)
	}
	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range array2 {
		fmt.Println(value)
	}
	j := 0
	fmt.Println("\n---------------Example 4--------------------\n")
	for range array2 {
		fmt.Println(array2[j])
		j++
	}

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence
	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %v\n", *strArray3)
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func TestElementExitInArray(t *testing.T) {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))
}
