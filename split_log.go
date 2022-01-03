package main

import (
	"bufio"
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
	args := os.Args
	filePath := args[1]
	start := args[2]
	end := "END"
	if len(args) > 3 {
		end = args[3]
	}
	//fmt.Println(filePath)
	//fmt.Println(start)
	//fmt.Println(end)

	f, err := os.Open(filePath)
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
}
