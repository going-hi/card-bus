package deck

import (
	"card-bus/card"
	"card-bus/shared"
	"card-bus/utils"
	"math/rand"
)

type Deck struct {
	cards []string
}

func baseCardCreate() []string {
	var cards = []string{}
	for _, m := range shared.SuitList {
		for _, v := range shared.CardList {
			card := string(m) + string(v)
			cards = append(cards, card)
		}
	}
	return cards
}

func NewDeck() *Deck {
	var cards = baseCardCreate()

	return &Deck{
		cards: cards,
	}
}

func (d *Deck) GetRandomCard() string {
	randomIndex := rand.Intn(len(d.cards) - 1)
	card := d.cards[randomIndex]
	d.cards = utils.RemoveElement(d.cards, randomIndex)
	return card
}

func (d *Deck) RebootDeck() {
	d.cards = baseCardCreate()
}

// first > second = 1; === 0 ; first < second = -1
func (d *Deck) ComparisonCards(cardFirst *card.Card, cardSecond *card.Card) int {
	firstWeight := d.getWeightCard(*cardFirst)
	secondWeight := d.getWeightCard(*cardSecond)

	if firstWeight < secondWeight {
		return -1
	}

	if firstWeight > secondWeight {
		return 1
	}

	return 0
}

func (d *Deck) getWeightCard(card card.Card) int {
	return shared.CardWeightMap[card.Value]
}

func (d *Deck) IsBetweenCards(currentCard, firstCard, secondCard card.Card) bool {
	weightCurrentCard := d.getWeightCard(currentCard)
	weightFirstCard := d.getWeightCard(firstCard)
	weightSecondCard := d.getWeightCard(secondCard)

	if weightCurrentCard == weightFirstCard || weightCurrentCard == weightSecondCard {
		return true
	}

	if weightCurrentCard > weightFirstCard && weightCurrentCard < weightSecondCard {
		return true
	}

	if weightCurrentCard < weightFirstCard && weightCurrentCard > weightSecondCard {
		return true
	}

	return false
}
