package main

import (
	"fmt"
)

type ProgrammingLanguage struct {
	name string
	year int
}

func main() {
	java := ProgrammingLanguage{name: "Java", year: 1996}
	golang := ProgrammingLanguage{name: "Golang", year: 2012}
	c := ProgrammingLanguage{name: "C", year: 1972}
	fmt.Println(java, golang, c)
}
