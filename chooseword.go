package Hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)


func Chooseword() string{
	file, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	var list_word_str []string
	var result string
	list_word_str = strings.Split(string(file), "\n")

	rand.Seed(time.Now().UnixNano())
   	v := rand.Intn(len(list_word_str))
	result = list_word_str[v]
	fmt.Println(result)
	return result[:len(result)-1]
}