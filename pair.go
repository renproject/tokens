package tokens

import (
	"errors"
	"strings"
)

// ErrInvalidTokenPair is returned when the given token pair is invalid.
var ErrInvalidTokenPair = errors.New("invalid token pair")

// Pair is a numerical representation for a token pairing.
type Pair uint64

// NewPair creates a new pair from the given tokens.
func NewPair(send, receive Token) Pair {
	if send.Code > receive.Code {
		return Pair((uint64(send.Code) << 32) | uint64(receive.Code))
	}
	return Pair((uint64(receive.Code) << 32) | uint64(send.Code))
}

// Pair values.
var (
	PairBTCDAI  = Pair((uint64(BTC.Code) << 32) | uint64(DAI.Code))
	PairETHDAI  = Pair((uint64(ETH.Code) << 32) | uint64(DAI.Code))
	PairRENDAI  = Pair((uint64(REN.Code) << 32) | uint64(DAI.Code))
	PairTUSDDAI = Pair((uint64(TUSD.Code) << 32) | uint64(DAI.Code))

	PairETHBTC  = Pair((uint64(ETH.Code) << 32) | uint64(BTC.Code))
	PairRENBTC  = Pair((uint64(REN.Code) << 32) | uint64(BTC.Code))
	PairTUSDBTC = Pair((uint64(TUSD.Code) << 32) | uint64(BTC.Code))
)

// Pairs is a list of all supported token pairs.
var Pairs = []Pair{
	PairBTCDAI,
	PairETHDAI,
	PairRENDAI,
	PairTUSDDAI,
	PairETHBTC,
	PairRENBTC,
	PairTUSDBTC,
}

// QuoteToken returns the base token for a given pair.
func (pair Pair) QuoteToken() Token {
	return ParseToken(Code(pair & 0x00000000FFFFFFFF))
}

// BaseToken returns the quote token for a given pair.
func (pair Pair) BaseToken() Token {
	return ParseToken(Code(pair >> 32))
}

// String returns a human-readable representation for a given pair.
func (pair Pair) String() string {
	switch pair {
	case PairBTCDAI:
		return "BTC-DAI"
	case PairETHDAI:
		return "ETH-DAI"
	case PairRENDAI:
		return "REN-DAI"
	case PairTUSDDAI:
		return "TUSD-DAI"
	case PairETHBTC:
		return "ETH-BTC"
	case PairRENBTC:
		return "REN-BTC"
	case PairTUSDBTC:
		return "TUSD-BTC"
	default:
		panic(ErrInvalidTokenPair.Error())
	}
}

// PatchPair parses the given string into a token pair. It returns
// `ErrInvalidTokenPair` if the pair is invalid.
func PatchPair(pair string) (Pair, error) {
	pair = strings.ToUpper(strings.TrimSpace(pair))
	switch pair {
	case "BTC-DAI":
		return PairBTCDAI, nil
	case "ETH-DAI":
		return PairETHDAI, nil
	case "REN-DAI":
		return PairRENDAI, nil
	case "TUSD-DAI":
		return PairTUSDDAI, nil
	case "ETH-BTC":
		return PairETHBTC, nil
	case "REN-BTC":
		return PairRENBTC, nil
	case "TUSD-BTC":
		return PairTUSDBTC, nil
	default:
		return 0, ErrInvalidTokenPair
	}
}
