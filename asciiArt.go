package Hangman

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"strings"
)

func AsciiArt(word []rune, filename string) {
	if filename != "normal" {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}

		lines := strings.Split(string(file), "\r\n")

		for j := 0; j < 9; j++ {
			for i := 0; i < len(word); i++ {
				if word[i] != '_' {
					nblettre := rune(word[i]) - 'A'
					color.Cyan.Print(lines[297+int(nblettre)*9+j])
				} else {
					color.Red.Print(lines[568+j])
				}
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println(string(word))
	}
}
