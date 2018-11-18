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

	largest := -999999999.9999
	for s := range in {
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		if i > largest {
			largest = i
		}
	}
	fmt.Printf("%v\n", largest)
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
