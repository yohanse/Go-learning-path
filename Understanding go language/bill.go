package main

import "fmt"

type bill struct {
	items map[string]float64
	tip int
}

func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key, value)
		total += value
	}

	fs += fmt.Sprintf("%-25v ...$%0.11v", "total:", total)
	return fs

}
