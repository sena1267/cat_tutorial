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
func catN(scanner *bufio.Scanner) {
	lineNum := 1
	// func (s *Scanner) Scan() bool
	for scanner.Scan() {
		fmt.Printf("%5d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
}

// -b:空行以外に行番号を付けて表示
func catB(scanner *bufio.Scanner) {
	lineNum := 1
	// func (s *Scanner) Scan() bool
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fmt.Println(scanner.Text())
		} else {
			fmt.Printf("%5d: %s\n", lineNum, scanner.Text())
			lineNum++
		}
	}
}
