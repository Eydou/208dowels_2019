//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// main
//

package main

import (
	"fmt"
	"os"
	"strconv"

	functions "./functions"
)

func help() {
	fmt.Printf("USAGE\n   ./208dowels O0 O1 O2 O3 O4 O5 O6 O7 O8\n")
	fmt.Printf("\nDESCRIPTION\n")
	fmt.Printf("   Oi    size of the observed class\n")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			help()
		}
	}
	if _, err := functions.ErrorArgs(args); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mX\033[0m Error: %s\n", err)
		os.Exit(84)
	}
	numbers := make([]int, 0)
	for i := 1; i != len(args); i++ {
		nb, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println(err, nb)
			os.Exit(84)
		}
		numbers = append(numbers, nb)
	}
	totalArgs := 0
	for _, num := range numbers {
		totalArgs += num
	}
	if totalArgs != 100 {
		os.Exit(84)
	}
	os.Exit(functions.MathParse(numbers))
}
