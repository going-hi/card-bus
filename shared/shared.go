package shared

import (
	"github.com/fatih/color"
)

type ColorFmt = *color.Color

var (
	RedColorFmt   ColorFmt = color.New(color.FgRed)
	BlackColorFmt ColorFmt = color.New(color.FgBlack)
)

type SuitString string

const (
	SpadesSuit   SuitString = "spades"
	ClubsSuit    SuitString = "clubs"
	HeartsSuit   SuitString = "hearts"
	DiamondsSuit SuitString = "diamonds"
)

type SuitChar string

const (
	SpadesSuitChar   SuitChar = "♠"
	ClubsSuitChar    SuitChar = "♣"
	HeartsSuitChar   SuitChar = "♥"
	DiamondsSuitChar SuitChar = "♦"
)

type ColorCard string

const (
	BlackColorCard ColorCard = "Black"
	RedColorCard   ColorCard = "Red"
)

var SuitCharColorFmtMap = map[SuitChar]*color.Color{
	SpadesSuitChar:   BlackColorFmt,
	ClubsSuitChar:    BlackColorFmt,
	HeartsSuitChar:   RedColorFmt,
	DiamondsSuitChar: RedColorFmt,
}

var SuitCharColorMap = map[SuitChar]ColorCard{
	SpadesSuitChar:   BlackColorCard,
	ClubsSuitChar:    BlackColorCard,
	HeartsSuitChar:   RedColorCard,
	DiamondsSuitChar: RedColorCard,
}

var SuitCharSuitStringMap = map[SuitChar]SuitString{
	SpadesSuitChar:   SpadesSuit,
	ClubsSuitChar:    ClubsSuit,
	HeartsSuitChar:   HeartsSuit,
	DiamondsSuitChar: DiamondsSuit,
}

var SuitList = []SuitChar{
	SpadesSuitChar, ClubsSuitChar, HeartsSuitChar, DiamondsSuitChar,
}
