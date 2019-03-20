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
		return "DAI-BTC"
	case PairDAIETH:
		return "DAI-ETH"
	case PairDAIREN:
		return "DAI-REN"
	case PairDAITUSD:
		return "DAI-TUSD"
	case PairBTCETH:
		return "BTC-ETH"
	case PairBTCREN:
		return "BTC-REN"
	case PairBTCTUSD:
		return "BTC-TUSD"
	default:
		return ErrInvalidTokenPair.Error()
	}
}

// ParsePair parses the given string into a token pair. It returns
// `ErrInvalidTokenPair` if the pair is invalid.
func ParsePair(pair string) (Pair, error) {
	pair = strings.ToUpper(strings.TrimSpace(pair))
	switch pair {
	case "DAI-BTC":
		return PairDAIBTC, nil
	case "DAI-ETH":
		return PairDAIETH, nil
	case "DAI-REN":
		return PairDAIREN, nil
	case "DAI-TUSD":
		return PairDAITUSD, nil
	case "BTC-ETH":
		return PairBTCETH, nil
	case "BTC-REN":
		return PairBTCREN, nil
	case "BTC-TUSD":
		return PairBTCTUSD, nil
	default:
		return 0, ErrInvalidTokenPair
	}
}
