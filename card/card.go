package card

import (
	"card-bus/shared"
	"card-bus/utils"
	"fmt"
)

type Card struct {
	Value shared.CardValue
	Suit  shared.SuitChar
}

func (c *Card) GetColor() string {
	return string(shared.SuitCharColorMap[c.Suit])
}

func (c *Card) GetColorFmt() shared.ColorFmt {
	return shared.SuitCharColorFmtMap[c.Suit]
}

func (c *Card) GetSuitString() string {
	return string(shared.SuitCharSuitStringMap[c.Suit])
}

func NewCard(card string) *Card {
	value, suit := utils.ParseCard(card)
	return &Card{
		Value: value,
		Suit:  suit,
	}
}

func (c *Card) GetImageCard() string {
	value := c.Value
	suit := c.Suit

	suitColor := c.GetColorFmt()

	topVal := value
	if len(value) == 1 {
		topVal += " "
	}
	topValColor := suitColor.Sprint(topVal)

	bottomVal := value
	if len(value) == 1 {
		bottomVal = " " + value
	}
	bottomValColor := suitColor.Sprint(bottomVal)
	colorSuit := suitColor.Sprint(suit)
	mid := fmt.Sprintf("│    %s    │", colorSuit)
	card := fmt.Sprintf(`
┌─────────┐
│%s       │
│%s        │
│         │
%s
│         │
│        %s│
│       %s│
└─────────┘
`, topValColor, colorSuit, mid, colorSuit, bottomValColor)

	return card
}
