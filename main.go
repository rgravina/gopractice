package main

import (
	"fmt"
	"strconv"
)

type ProgrammingLanguage struct {
	name string
	year int
}

func (pl ProgrammingLanguage) description() string {
	return pl.name + "\t- Released in " + strconv.Itoa(pl.year) + "."
}

func main() {
	c := ProgrammingLanguage{name: "C", year: 1972}
	java := ProgrammingLanguage{name: "Java", year: 1996}
	golang := ProgrammingLanguage{name: "Golang", year: 2012}

	oldestLanguage := &c
	fmt.Println(oldestLanguage.description())
	fmt.Println(java.description())
	fmt.Println(golang.description())

}
