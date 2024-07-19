// Fetch выполняет параллельную выборку URL и записывает в файл
// информацию о затраченном времени и размере ответа для каждого из них.
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"os"
)

func main() {
	urls := []string{
		"",
		"",
		"",
	}
	start := time.Now()
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		result := <- ch
		fmt.Fprintln(file, result)
		// Получение из канала ch и запись в файл
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}