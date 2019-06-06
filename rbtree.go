package main

import (
	"fmt"
)

func check(err error) {
	if err != nil {
		fmt.Println("here man")
		return
	}
}
func main() {
	var N_fields int8
	_, err := fmt.Scanf("%d", &N_fields)
	fields_lacks := make([]int, N_fields*8)
	N_lacks := make([]int, N_fields)
}
