package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var path, start, end string
	flag.StringVar(&path, "p", "", "file path")
	flag.StringVar(&start, "s", "", "start value")
	flag.StringVar(&end, "e", "END", "end value")
	flag.Parse()

	if len(path) == 0 {
		fmt.Println("Usage: -p=\"file_path\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(path)
	check(err)

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)
	isExist := false

	for fs.Scan() {
		l := fs.Text()
		if strings.HasPrefix(l, start) {
			isExist = true
		}
		if strings.HasPrefix(l, end) {
			break
		}
		if isExist {
			fmt.Println(l)
		}
	}
	f.Close()
	os.Exit(0)
}
