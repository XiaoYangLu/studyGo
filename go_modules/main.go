package main

import (
	"fmt"
	"sync"
)

func main() {
	//h := requesthtml.RequestHtml("www.baidu.com")
	//fmt.Println(h)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			fmt.Println("+", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			fmt.Println("-", i)
		}
	}()
	wg.Wait()
}
