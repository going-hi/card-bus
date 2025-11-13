package utils

import (
	"card-bus/shared"
	"fmt"
)

func RemoveElement(list []string, index int) []string {
	return append(list[:index], list[index+1:]...)
}

// input "♠10", output "10", "♠"
func ParseCard(card string) (shared.CardValue, shared.SuitChar) {
	runes := []rune(card)
	suit := shared.SuitChar(runes[0])
	value := shared.CardValue(runes[1:])

	return value, suit
}

func PromptData(prompt ...string) string {

	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scanln(&res)
	return res
}
