package shared

type CardValue string

const (
	Two   CardValue = "2"
	Three CardValue = "3"
	Four  CardValue = "4"
	Five  CardValue = "5"
	Six   CardValue = "6"
	Seven CardValue = "7"
	Eight CardValue = "8"
	Nine  CardValue = "9"
	Ten   CardValue = "10"
	Jack  CardValue = "J"
	Queen CardValue = "Q"
	King  CardValue = "K"
	Ace   CardValue = "A"
)

var CardList = []CardValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

var CardWeightMap = map[CardValue]int{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}

func (v *CardValue) ToWeight() int {
	return CardWeightMap[*v]
}
