package main

import (
	"card-bus/deck"
	"card-bus/menu"
)

func main() {
	deck := deck.NewDeck()
	menu := menu.Menu{
		Deck: deck,
	}

	menu.Welcome()

Menu:
	for {
		next := menu.FirstStep()

		if !next {
			break Menu
		}

		next = menu.SecondStep()

		if !next {
			break Menu
		}

		next = menu.ThirdStep()

		if !next {
			break Menu
		}

		next = menu.FourthStep()

		if !next {
			break
		}
	}

}