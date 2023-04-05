package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	a := "http://github.com"
	b := "http://baidu.com"

	start1 := time.Now()
	http.Get(a)
	d1 := time.Since(start1)

	fmt.Printf("a url req duration = %.2f ms\n", float64(d1/1000000.0))

	start2 := time.Now()
	http.Get(b)
	d2 := time.Since(start2)

	fmt.Printf("a url req duration = %.2f ms\n", float64(d2/1000000.0))

}
