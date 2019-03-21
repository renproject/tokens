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
	if send.Code < receive.Code {
		return Pair((uint64(send.Code) << 32) | uint64(receive.Code))
	}
	return Pair((uint64(receive.Code) << 32) | uint64(send.Code))
}

// Pair values.
var (
	PairDAIBTC  = Pair((uint64(DAI.Code) << 32) | uint64(BTC.Code))
	PairDAIETH  = Pair((uint64(DAI.Code) << 32) | uint64(ETH.Code))
	PairDAIREN  = Pair((uint64(DAI.Code) << 32) | uint64(REN.Code))
	PairDAITUSD = Pair((uint64(DAI.Code) << 32) | uint64(TUSD.Code))

	PairBTCETH  = Pair((uint64(BTC.Code) << 32) | uint64(ETH.Code))
	PairBTCREN  = Pair((uint64(BTC.Code) << 32) | uint64(REN.Code))
	PairBTCTUSD = Pair((uint64(BTC.Code) << 32) | uint64(TUSD.Code))
)

// Pairs is a list of all supported token pairs.
var Pairs = []Pair{
	PairDAIBTC,
	PairDAIETH,
	PairDAIREN,
	PairDAITUSD,
	PairBTCETH,
	PairBTCREN,
	PairBTCTUSD,
}

// BaseToken returns the base token for a given pair.
func (pair Pair) BaseToken() Token {
	return ParseTokenCode(Code(pair >> 32))
}

// QuoteToken returns the quote token for a given pair.
func (pair Pair) QuoteToken() Token {
	return ParseTokenCode(Code(pair & 0x00000000FFFFFFFF))
}

// String returns a human-readable representation for a given pair.
func (pair Pair) String() string {
	switch pair {
	case PairDAIBTC:
		return "BTC-DAI"
	case PairDAIETH:
		return "ETH-DAI"
	case PairDAIREN:
		return "REN-DAI"
	case PairDAITUSD:
		return "TUSD-DAI"
	case PairBTCETH:
		return "ETH-BTC"
	case PairBTCREN:
		return "REN-BTC"
	case PairBTCTUSD:
		return "TUSD-BTC"
	default:
		return ErrInvalidTokenPair.Error()
	}
}

// ParsePair parses the given string into a token pair. It returns
// `ErrInvalidTokenPair` if the pair is invalid.
func ParsePair(pair string) (Pair, error) {
	pair = strings.ToUpper(strings.TrimSpace(pair))
	switch pair {
	case "BTC-DAI":
		return PairDAIBTC, nil
	case "ETH-DAI":
		return PairDAIETH, nil
	case "REN-DAI":
		return PairDAIREN, nil
	case "TUSD-DAI":
		return PairDAITUSD, nil
	case "ETH-BTC":
		return PairBTCETH, nil
	case "REN-BTC":
		return PairBTCREN, nil
	case "TUSD-BTC":
		return PairBTCTUSD, nil
	default:
		return 0, ErrInvalidTokenPair
	}
}
