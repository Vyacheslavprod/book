// reverse обращает порядок чисел на месте
package main

import "fmt"

func r(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	r(a)
	fmt.Println(a)

	//Циклический сдвиг среза на n элементов
	r(a[:2])
	r(a[2:])
	r(a)
	fmt.Println(a)
}