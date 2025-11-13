package menu

import (
	"card-bus/card"
	"card-bus/deck"
	"card-bus/map-step"
	"card-bus/utils"
	"fmt"
)

type Menu struct {
	Deck       *deck.Deck
	firstCard  *card.Card
	secondCard *card.Card
}

func (m *Menu) Welcome() {
	fmt.Println("Привет! Это игра карточный автобус! Давай начнем!")
}

// next
func (m *Menu) FirstStep() bool {
	mapMenu := GetMapFirstStep()

	varNumber := utils.PromptData(mapMenu.InputMenu...)

	variant, err := mapstep.GetVariantMapStep[string](mapMenu, varNumber)
	if err != nil {
		panic(err)
	}

	if variant == "Exit" {
		fmt.Println("Выход")
		fmt.Println("-------------------------------")
		return false
	}

	randomCard := card.NewCard(m.Deck.GetRandomCard())

	colorCard := randomCard.GetColor()

	fmt.Println(randomCard.GetImageCard())
	if variant == colorCard {
		m.firstCard = randomCard
		fmt.Println("Ты угадал!")
		fmt.Println("-------------------------------")
	} else {
		fmt.Println("Не угадал")
		fmt.Println("-------------------------------")
		return false
	}

	return true
}

func ResultComparison(d *deck.Deck, comparison int, variant int) bool {
	if comparison == 0 {
		return true
	}

	return comparison == variant
}

func (m *Menu) SecondStep() bool {
	mapMenu := GetMapSecondStep()
	varNumber := utils.PromptData(mapMenu.InputMenu...)

	if varNumber == "3" {
		fmt.Println("Выход")
		fmt.Println("-------------------------------")
		return false
	}

	variant, err := mapstep.GetVariantMapStep[int](mapMenu, varNumber)

	if err != nil {
		panic(err)
	}

	randomCard := card.NewCard(m.Deck.GetRandomCard())

	comparison := m.Deck.ComparisonCards(m.firstCard, randomCard)

	isRight := ResultComparison(m.Deck, comparison, variant)

	fmt.Println(randomCard.GetImageCard())
	if isRight {
		m.secondCard = randomCard
		fmt.Println("Ты угадал!")
		fmt.Println("-------------------------------")
	} else {
		fmt.Println("Не угадал")
		fmt.Println("-------------------------------")
		return false
	}

	return true
}

func (m *Menu) ThirdStep() bool {
	mapMenu := GetMapThirthStep()
	varNumber := utils.PromptData(mapMenu.InputMenu...)

	if varNumber == "3" {
		fmt.Println("Выход")
		fmt.Println("-------------------------------")
		return false
	}

	variant, err := mapstep.GetVariantMapStep[bool](mapMenu, varNumber)
	if err != nil {
		panic(err)
	}

	randomCard := card.NewCard(m.Deck.GetRandomCard())
	isBetween := m.Deck.IsBetweenCards(*randomCard, *m.firstCard, *m.secondCard)

	fmt.Println(randomCard.GetImageCard())
	if variant == isBetween {
		fmt.Println("Ты угадал!")
		fmt.Println("-------------------------------")
		return true
	} else {
		fmt.Println("Не угадал")
		fmt.Println("-------------------------------")
		return false
	}
}

func (m *Menu) FourthStep() bool {
	mapMenu := GetMapFourthStep()
	varNumber := utils.PromptData(mapMenu.InputMenu...)

	if varNumber == "5" {
		fmt.Println("Выход")
		fmt.Println("-------------------------------")
		return false
	}

	variant, err := mapstep.GetVariantMapStep[string](mapMenu, varNumber)
	if err != nil {
		panic(err)
	}

	randomCard := card.NewCard(m.Deck.GetRandomCard())
	suitString := randomCard.GetSuitString()

	fmt.Println(randomCard.GetImageCard())
	if suitString == variant {
		fmt.Println("Ты угадал! Ты прошел все 4 раунда!!!")
		fmt.Println("-------------------------------")
		return true
	} else {
		fmt.Println("Ты не угадал!")
		fmt.Println("-------------------------------")
		return false
	}

}
