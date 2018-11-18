package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	// dictPath = "./american-english"
	dictPath = "/usr/share/dict/american-english"
)

func main() {
	c, _ := net.Dial("tcp", fmt.Sprintf("%v:%v", os.Args[1], os.Args[2]))
	defer c.Close()
	s := bufio.NewScanner(c)

	send := make(chan string)
	get := make(chan []string)
	go wordFinder(send, get)

	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	isFlag := regexp.MustCompile(`^RCN{.*}$`).MatchString

	fmt.Fprintf(c, "Y\n")
	for s.Scan() {
		line := s.Text()
		// fmt.Println(line)
		if len(strings.Split(line, " ")) == 1 {
			if isLetter(line) {
				time.Sleep(100 * time.Millisecond)
				send <- line
				ret := calcHash(line, <-get)
				// fmt.Println(ret)
				fmt.Fprintf(c, ret)
			} else if isFlag(line) {
				fmt.Println(line)
				break
			}
		}
	}

	close(send)
	close(get)
}

func initGame(c *net.Conn) {
	var buf bytes.Buffer
	io.Copy(&buf, *c)
	fmt.Print(buf)
}

func calcHash(word string, similar []string) string {
	h1 := md5.New()
	h2 := md5.New()
	io.WriteString(h1, word)
	io.WriteString(h2, strings.Join(similar, " "))
	return fmt.Sprintf("%x %x\n", h1.Sum(nil), h2.Sum(nil))
}

func wordFinder(in <-chan string, out chan<- []string) {
	for word := range in {
		ri := make(chan string)
		var ret []string
		go readLines(dictPath, ri)
		for line := range ri {
			if strings.Contains(strings.ToLower(line), word) {
				ret = append(ret, line)
			}
		}
		out <- ret
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
