package main

import "fmt"

/*Сравнение map
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
*/

func main() {
	//ajes := make(map[string]int)
	ages := map[string]int{
		"alice":   31,
		"charlie": 35,
	}
	fmt.Println(ages)

	delete(ages, "alice")
	ages["carol"] = 23
	fmt.Println(ages)

	age, ok := ages["bob"]
	if !ok {
		fmt.Println(age)
	}
	//age, ok := ages["bob"]; if !ok {...}

}