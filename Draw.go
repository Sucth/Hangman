package Hangman

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"strings"
)

func Draw(n int) {
	file, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(file), "\n")

	for j := len(lines) - (25 * (n + 1)); j < len(lines)-(25*(n+1))+25; j++ {
		color.Yellow.Println(string(lines[j]))
	}
}
