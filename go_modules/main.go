package main

import (
	"fmt"
	"studyGo/go_modules/requesthtml"
)

func main() {
	html := requesthtml.RequestHtml("https://www.iqiyi.com/")
	fmt.Println(html)
}
