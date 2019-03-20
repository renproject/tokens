package tokens

import (
	"errors"
	"math/rand"
	"reflect"
	"strings"
)

// ErrUnsupportedTokenPair is returned when the given token pair is unsupported.
var ErrUnsupportedTokenPair = errors.New("unsupported token pair")

// Name is a string representation of the token supported by Ren.
type Name string

var (
	NameDAI  = Name("DAI")
	NameBTC  = Name("BTC")
	NameETH  = Name("ETH")
	NameREN  = Name("REN")
	NameDGX  = Name("DGX")
	NameZRX  = Name("ZRX")
	NameOMG  = Name("OMG")
	NamePAX  = Name("PAX")
	NameUSDC = Name("USDC")
	NameGUSD = Name("GUSD")
	NameTUSD = Name("TUSD")
	NameWBTC = Name("WBTC")
)

// Code is a numerical representation of a token supported by RenEx.
type Code uint32

var (
	// Base tokens are ranging from 0 to 1023. Base tokens can be a quote
	// token in a Pair if the other base token has a higher token rank.
	CodeDAI Code = 100
	CodeBTC Code = 200

	// Quote tokens are ranging from 1024 to Max_Uint32
	CodeETH  Code = 1024
	CodeREN  Code = 1025
	CodeDGX  Code = 1026
	CodeZRX  Code = 1027
	CodeOMG  Code = 1028
	CodePAX  Code = 1029
	CodeGUSD Code = 1030
	CodeTUSD Code = 1031
	CodeUSDC Code = 1032
	CodeWBTC Code = 1033
)

// Token represents the token we are trading.
type Token struct {
	Name       Name           `json:"name"`
	Code       Code           `json:"code"`
	Decimals   int64          `json:"decimals"`
	Blockchain BlockchainName `json:"blockchain"`
}

var (
	DAI  = Token{NameDAI, CodeDAI, 18, ERC20}
	BTC  = Token{NameBTC, CodeBTC, 8, BITCOIN}
	ETH  = Token{NameETH, CodeETH, 18, ETHEREUM}
	REN  = Token{NameREN, CodeREN, 18, ERC20}
	DGX  = Token{NameDGX, CodeDGX, 9, ERC20}
	ZRX  = Token{NameZRX, CodeZRX, 18, ERC20}
	OMG  = Token{NameOMG, CodeOMG, 18, ERC20}
	PAX  = Token{NamePAX, CodePAX, 18, ERC20}
	GUSD = Token{NameGUSD, CodeGUSD, 2, ERC20}
	TUSD = Token{NameTUSD, CodeTUSD, 18, ERC20}
	USDC = Token{NameUSDC, CodeUSDC, 6, ERC20}
	WBTC = Token{NameWBTC, CodeWBTC, 8, ERC20}
)

// SupportedTokens contains all the tokens supported by Ren.
var SupportedTokens = []Token{
	DAI, BTC, ETH, REN, DGX, ZRX, PAX, OMG, GUSD, TUSD, USDC, WBTC,
}

// String returns the string representation of the token.
func (token Token) String() string {
	return string(token.Name)
}

// Generate implements the `Generator` interface used by quickCheck.
func (Token) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(SupportedTokens[rand.Int()%len(SupportedTokens)])
}

// ParseToken parses a string to a Token. It returns `ErrUnsupportedTokenPair`
// if the given string cannot be parsed to a supported token.
func ParseToken(token string) (Token, error) {
	token = strings.TrimSpace(strings.ToLower(token))
	switch token {
	case "dai", "maker-dai", "makerdai":
		return DAI, nil
	case "bitcoin", "btc", "xbt":
		return BTC, nil
	case "ethereum", "eth", "ether":
		return ETH, nil
	case "ren", "republictoken", "republic token":
		return REN, nil
	case "digix-gold-token", "dgx", "dgt":
		return DGX, nil
	case "zerox", "zrx", "0x":
		return ZRX, nil
	case "omisego", "omg", "omise-go":
		return OMG, nil
	case "pax", "paxosstandardtoken", "paxos-standard-token":
		return PAX, nil
	case "gusd", "gemini-dollar", "geminidollar":
		return GUSD, nil
	case "tusd", "trueusd", "true-usd":
		return TUSD, nil
	case "usdc", "usd-coin", "usdcoin":
		return USDC, nil
	case "wrappedbtc", "wbtc", "wrappedbitcoin":
		return WBTC, nil
	default:
		return Token{}, ErrUnsupportedTokenPair
	}
}

// ParseTokenCode parses a token Code to a Token. It returns `ErrUnsupportedTokenPair`
// if the given code cannot be parsed to a supported token.
func ParseTokenCode(code Code) (Token, error) {
	switch code {
	case CodeDAI:
		return DAI, nil
	case CodeBTC:
		return BTC, nil
	case CodeETH:
		return ETH, nil
	case CodeREN:
		return REN, nil
	case CodeDGX:
		return DGX, nil
	case CodeZRX:
		return ZRX, nil
	case CodeOMG:
		return OMG, nil
	case CodePAX:
		return PAX, nil
	case CodeGUSD:
		return GUSD, nil
	case CodeTUSD:
		return TUSD, nil
	case CodeUSDC:
		return USDC, nil
	case CodeWBTC:
		return WBTC, nil
	default:
		return Token{}, ErrUnsupportedTokenPair
	}
}
