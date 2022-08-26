package main

import (
	"bufio"
	"fmt"
)

// そのまま表示
func cat(scanner *bufio.Scanner) {
	// func (s *Scanner) Scan() bool
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// -n:行番号を付けて表示
func cat_n(scanner *bufio.Scanner) {
	line_num := 1
	// func (s *Scanner) Scan() bool
	for scanner.Scan() {
		fmt.Printf("%5d: %s\n", line_num, scanner.Text())
		line_num++
	}
}

// -b:空行以外に行番号を付けて表示
func cat_b(scanner *bufio.Scanner) {
	line_num := 1
	// func (s *Scanner) Scan() bool
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fmt.Println(scanner.Text())
		} else {
			fmt.Printf("%5d: %s\n", line_num, scanner.Text())
			line_num++
		}
	}
}
