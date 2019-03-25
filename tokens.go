package tokens

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// ErrUnsupportedTokenPair is returned when the given token pair is not
// supported.
var ErrUnsupportedTokenPair = errors.New("unsupported token pair")

// NewErrUnsupportedToken returns an error when the given token is not
// supported by Ren.
func NewErrUnsupportedToken(token Name) error {
	return fmt.Errorf("unsupported token: %s", token)
}

// Name is a string representation of a token.
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

	NameInvalid = Name("INVALID")
)

// Code is a numerical representation of a token.
type Code uint32

var (
	CodeInvalid Code = 0

	// Quote tokens range from 1 to 1023. Quote tokens can be used as a quote
	// token in a Pair if the other quote token has a lower token code.
	CodeDAI Code = 100
	CodeBTC Code = 200

	// Base tokens range from 1024 to Max_Uint32.
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

// Token provides a token representation.
type Token struct {
	Name       Name           `json:"name"`
	Code       Code           `json:"code"`
	Decimals   int64          `json:"decimals"`
	Blockchain BlockchainName `json:"blockchain"`
}

var InvalidToken = Token{NameInvalid, CodeInvalid, 0, ""}

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
func (token Token) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(SupportedTokens[rand.Int()%len(SupportedTokens)])
}

func (token Token) AdditionalTransactionFee(amount *big.Int) *big.Int {
	switch token {
	case DGX:
		return calculateFeesFromBips(amount, 13)
	default:
		return nil
	}
}

func calculateFeesFromBips(value *big.Int, bips int64) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(value, big.NewInt(bips)), new(big.Int).Sub(big.NewInt(10000), big.NewInt(bips)))
}

// ParseToken parses a string to a token. It returns an `InvalidToken` if the
// given string cannot be parsed to a token.
func ParseToken(token string) Token {
	token = strings.TrimSpace(strings.ToLower(token))
	switch token {
	case "dai", "maker-dai", "makerdai":
		return DAI
	case "bitcoin", "btc", "xbt":
		return BTC
	case "ethereum", "eth", "ether":
		return ETH
	case "ren", "republictoken", "republic token":
		return REN
	case "digix-gold-token", "dgx", "dgt":
		return DGX
	case "zerox", "zrx", "0x":
		return ZRX
	case "omisego", "omg", "omise-go":
		return OMG
	case "pax", "paxosstandardtoken", "paxos-standard-token":
		return PAX
	case "gusd", "gemini-dollar", "geminidollar":
		return GUSD
	case "tusd", "trueusd", "true-usd":
		return TUSD
	case "usdc", "usd-coin", "usdcoin":
		return USDC
	case "wrappedbtc", "wbtc", "wrappedbitcoin":
		return WBTC
	default:
		return InvalidToken
	}
}

// ParseTokenCode parses a token code to a token. It returns an `InvalidToken`
// if the given code cannot be parsed to a token.
func ParseTokenCode(code Code) Token {
	switch code {
	case CodeDAI:
		return DAI
	case CodeBTC:
		return BTC
	case CodeETH:
		return ETH
	case CodeREN:
		return REN
	case CodeDGX:
		return DGX
	case CodeZRX:
		return ZRX
	case CodeOMG:
		return OMG
	case CodePAX:
		return PAX
	case CodeGUSD:
		return GUSD
	case CodeTUSD:
		return TUSD
	case CodeUSDC:
		return USDC
	case CodeWBTC:
		return WBTC
	default:
		return InvalidToken
	}
}
