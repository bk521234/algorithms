
// GO language program with example of Array Sort Functions 
 
package main
 
import (
    "fmt"
    "sort"
)
func main() {
    num := []int{50,90, 30, 10, 50}
    fmt.Println(num)    
    /*  Check array of interger Is sorted   */
    if sort.IntsAreSorted(num)==false{
        sort.Ints(num)  // Sort Integer Array
    }
    fmt.Println(num)
    fmt.Println(sort.SearchInts(num,30))
 
    text := []string{"US","UK","Germany","Australia","Japan"}
    fmt.Println(text)   
    if sort.StringsAreSorted(text)==false{
        sort.Strings(text) // Sort String Array
    }
    fmt.Println(text)
    fmt.Println(sort.SearchStrings(text,"Japan"))
 
    val := []float64{2.5,6.3,1.5,9.8,4.7,1.1,5.2}
    fmt.Println(val)    
    if sort.Float64sAreSorted(val)==false{
        sort.Float64s(val) // Sort Float Array
    }
    fmt.Println(val)
    fmt.Println(sort.SearchFloat64s(val,5.2))
}
/*
Output
 
C:\golang>go run example25.go
[50 90 30 10 50]
[10 30 50 50 90]
1
[US UK Germany Australia Japan]
[Australia Germany Japan UK US]
2
[2.5 6.3 1.5 9.8 4.7 1.1 5.2]
[1.1 1.5 2.5 4.7 5.2 6.3 9.8]
4
 
 
*/