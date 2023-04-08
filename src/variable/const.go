package main

import (
	"fmt"
)

func main() {
	const a int = 10
	//错误示范
	// const a := 10 

	fmt.Printf("a = %d \n", a)

	const (
		m_a int = 666
		m_b string = "ljk"
		m_c float64 = 10.2
	)

	//错误示范
	//m_a = 777

	fmt.Printf("m_a=%d, m_b=%s, m_c=%.2f", 
	m_a,
	m_b,
	m_c)
}