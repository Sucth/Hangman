package main

import (
	"Hangman"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	var win bool
	var found bool
	var letter_used bool
	var letter string
	var lifes int
	var alphabname string
	var letters_used []rune
	wordstofind := Hangman.Chooseword()
	lwordstofind := Hangman.StartGame(wordstofind)
	nblifes := 10
	if len(args) != 1 {
		for i := range args {
			if args[i] == "--startWith" {
				Hangman.AsciiArt([]rune("Welcome Back !!"), alphabname)
				nblifes, lwordstofind = Hangman.Start(args[i+1])
			} else if args[i] == "--letterFile" {
				alphabname = args[i+1]
			}
		}
	} else {
		alphabname = "normal"
	}
	Hangman.AsciiArt([]rune("        the HANGMAN"), alphabname)
	Hangman.AsciiArt([]rune("        ==========="), alphabname)

	fmt.Print("\n")
	Hangman.AsciiArt([]rune("Good Luck, you have 10"), alphabname)
	Hangman.AsciiArt([]rune("attempts."), alphabname)
	fmt.Print("\r\n    ")
	Hangman.AsciiArt(lwordstofind, alphabname)
	for lifes = nblifes; lifes >= 0; lifes-- {
		fmt.Print("\n")
		found = false
		fmt.Print("Choose : ")
		fmt.Scanf("%s\n", &letter)
		if len(letter) > 1 {
			if letter == "STOP" {
				Hangman.AsciiArt([]rune("Game Saved !"), alphabname)
				Hangman.Stop(lwordstofind, lifes)
				break
			}
			if Hangman.Wordtest(wordstofind, letter) {
				Hangman.AsciiArt([]rune("You won GG !"), alphabname)
				break
			} else if !Hangman.Wordtest(wordstofind, letter) {
				lifes -= 1
				if lifes <= 0 {
					Hangman.AsciiArt([]rune("You are dead !"), alphabname)
					break
				} else {
					Hangman.AsciiArt([]rune("Wrong word ! You lost 2 lifes,"), alphabname)
					Hangman.AsciiArt([]rune{rune(lifes) + '0', ' ', 'a', 't', 't', 'e', 'm', 'p', 't', 's', ' ', 'r', 'e', 'm', 'a', 'i', 'n', 'i', 'n', 'g'}, alphabname)
					fmt.Print("\n")
					Hangman.Draw(lifes)
					continue
				}
			}
		} else if len(letter) == 1 {
			letter_used = false
			if len(letters_used) == 0 {
				letters_used = append(letters_used, rune(letter[0]))
			} else {
				for i := 0; i < len(letters_used); i++ {
					if letters_used[i] == rune(letter[0]) {
						letter_used = true
						break
					}
				}
			}
			if letter_used {
				Hangman.AsciiArt([]rune("You already used this letter !"), alphabname)
				lifes++
				continue
			} else {
				letters_used = append(letters_used, rune(letter[0]))
			}
			for i := 0; i < len(wordstofind); i++ {
				if rune(wordstofind[i]-32) == rune(letter[0]) {
					lwordstofind[i] = rune(letter[0])
					found = true
				}
			}
			if found {
				if lifes == 10 {
				} else {
					lifes++
				}
				Hangman.AsciiArt(lwordstofind, alphabname)
				fmt.Print("\n")
			} else {
				if lifes == 10 {
					lifes--
				}
				Hangman.AsciiArt([]rune("Not present in the word, "), alphabname)
				Hangman.AsciiArt([]rune{rune(lifes) + '0', ' ', 'a', 't', 't', 'e', 'm', 'p', 't', 's', ' ', 'r', 'e', 'm', 'a', 'i', 'n', 'i', 'n', 'g'}, alphabname)
				fmt.Print("\n")
				Hangman.Draw(lifes)
				Hangman.AsciiArt(lwordstofind, alphabname)
			}
			win = true
			for _, n := range lwordstofind {
				if n == '_' {
					win = false
				}
			}
			if win {
				fmt.Print("\n\n")
				Hangman.AsciiArt([]rune("You won, GG !"), alphabname)
				break
			} else if lifes == 0 {
				fmt.Print("\n")
				Hangman.AsciiArt([]rune("You are dead !"), alphabname)
			}
		} else {
			Hangman.AsciiArt([]rune("One letter is expected, (you lost one try)"), alphabname)
			Hangman.AsciiArt(lwordstofind, alphabname)
			fmt.Print("\n")
			Hangman.Draw(lifes)
		}
	}
}
