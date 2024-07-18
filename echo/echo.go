// Echo - выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Now())
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now())
	fmt.Println(os.Args[1:])
}