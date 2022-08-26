package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	var (
		n = flag.Bool("n", false, "行番号を付けて表示する")
		b = flag.Bool("b", false, "空行以外に行番号を付けて表示する")
	)

	flag.Parse()

	/*
		// 引数の数が合わないとき
		if len(os.Args) != 2 {
			fmt.Println("１つのファイル名を指定してください")
			os.Exit(1)
		}
	*/
	var scanner *bufio.Scanner
	if len(os.Args) == 2 {
		fp, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}
		defer fp.Close()
		scanner = bufio.NewScanner(fp)
	} else if len(os.Args) == 3 {
		fp, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		defer fp.Close()
		scanner = bufio.NewScanner(fp)
	} else {
		os.Exit(1)
	}

	// func NewScanner(r io.Reader) *Scanner
	//scanner := bufio.NewScanner(fp)

	if *n {
		cat_n(scanner)
	} else if *b {
		cat_b(scanner)
	} else {
		cat(scanner)
	}
}
