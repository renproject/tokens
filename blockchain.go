package tokens

import "fmt"

// NewErrUnsupportedBlockchain returns an error when the given blockchain is
// not supported by Ren.
func NewErrUnsupportedBlockchain(blockchain BlockchainName) error {
	return fmt.Errorf("unsupported blockchain: %s", blockchain)
}

// BlockchainName is the name of the blockchain.
type BlockchainName string

var (
	BITCOIN  = BlockchainName("bitcoin")
	ETHEREUM = BlockchainName("ethereum")
	ERC20    = BlockchainName("erc20")
)
