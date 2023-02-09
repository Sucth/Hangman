package Hangman

 import (
 	"math/rand"
 	"time"
 )

 func StartGame(s string) []rune {
 	s2 := []rune(s)
 	var tab []rune

 	for i := 0; i < len(s2); i++ {
 		tab = append(tab, '_')

 	}

     rand.Seed(time.Now().UnixNano())
     for i:=0; i<len(s)/2-1; i++ {
         v := rand.Intn(len(s))
 		if tab[v] == '_' {
 			tab[v] = (rune(s[v] - 32))
 		} else {
 			i--
 		}
     }
 	return tab
 }
