// Массивы
package main

import "fmt"



func main() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	months := [...]string{1: "Январь", 5: "Май", 12: "Декабрь"}
	fmt.Println(months)
	fmt.Println(len(months))
}