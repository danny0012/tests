package main

import (
	"fmt"
	"tests/prose"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}