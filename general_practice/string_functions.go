// Golang Compare function [comparing two strings lexicographically]
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("A", "B"))             // A < B
	fmt.Println(strings.Compare("B", "A"))             // B > A
	fmt.Println(strings.Compare("Japan", "Australia")) // J > A
	fmt.Println(strings.Compare("Australia", "Japan")) // A < J
	fmt.Println(strings.Compare("Germany", "Germany")) // G == G
	fmt.Println(strings.Compare("Germany", "GERMANY")) // GERMANY > Germany
	fmt.Println(strings.Compare("", ""))
	fmt.Println(strings.Compare("", " ")) // Space is less
	containsFunction()
	containsAnyFunction()
}

func containsFunction() {
	fmt.Println(strings.Contains("Australia", "Aus")) // Any part of string
	fmt.Println(strings.Contains("Australia", "Australian"))
	fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
	fmt.Println(strings.Contains("Japan", "JAP"))   // Case sensitive
	fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("  ", " ")) // space also consider as string
	fmt.Println(strings.Contains("12554", "1"))
}

func containsAnyFunction() {
	fmt.Println(strings.ContainsAny("Australia", "a"))
	fmt.Println(strings.ContainsAny("Australia", "r & a"))
	fmt.Println(strings.ContainsAny("JAPAN", "j"))
	fmt.Println(strings.ContainsAny("JAPAN", "J"))
	fmt.Println(strings.ContainsAny("JAPAN", "JAPAN"))
	fmt.Println(strings.ContainsAny("JAPAN", "japan"))
	fmt.Println(strings.ContainsAny("Shell-12541", "1"))

	//  Contains vs ContainsAny
	fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) // true
	fmt.Println(strings.Contains("Shell-12541", "1-2"))    // false
}

func countInstances() {
	fmt.Println(strings.Count("Australia", "a"))
	fmt.Println(strings.Count("Australia", "A"))
	fmt.Println(strings.Count("Australia", "M"))
	fmt.Println(strings.Count("Japanese", "Japan")) // 1
	fmt.Println(strings.Count("Japan", "Japanese")) // 0
	fmt.Println(strings.Count("Shell-25152", "-25"))
	fmt.Println(strings.Count("Shell-25152", "-21"))
	fmt.Println(strings.Count("test", "")) // length of string + 1
	fmt.Println(strings.Count("test", " "))
}