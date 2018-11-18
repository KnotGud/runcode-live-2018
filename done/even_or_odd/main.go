package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fn := os.Args[1]
	in := make(chan string)
	go readLines(fn, in)
	for s := range in {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			fmt.Printf("%v True\n", s)
		} else {
			fmt.Printf("%v False\n", s)
		}

	}
}

func readLines(path string, out chan<- string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		out <- s.Text()
	}
	close(out)
}
