package main

import "fmt"

func main() {
	var myBill bill = bill {
		items: map[string]float64{"pie": 3.4, "car": 2.3},
		tip: 0,
	}

	fmt.Println(myBill.format())
}