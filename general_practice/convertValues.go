package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println("\nConvert String into Boolean Data type\n")
	str1 := "japan"
	fmt.Println("Before :", reflect.TypeOf(str1))
	bolStr, _ := strconv.ParseBool(str1)
	fmt.Println("After :", reflect.TypeOf(bolStr))
	fmt.Println(bolStr)

	fmt.Println("\nConvert String into Float64 Data type\n")
	str2 := "japan"
	fmt.Println("Before :", reflect.TypeOf(str2))
	fltStr, _ := strconv.ParseFloat(str2, 64)
	fmt.Println("After :", reflect.TypeOf(fltStr))
	fmt.Println(fltStr)

	fmt.Println("\nConvert String into Integer Data type\n")
	str3 := "japan"
	fmt.Println("Before :", reflect.TypeOf(str3))
	intStr, _ := strconv.ParseInt(str3, 10, 64)
	fmt.Println("After :", reflect.TypeOf(intStr))
	fmt.Println(intStr)

}
