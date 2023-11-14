package chain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func EtherToWei(amount int64) *big.Int {
	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Mul(big.NewInt(amount), ether)
}

func Has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func IsValidAddress(address string, checksummed bool) bool {
	if !common.IsHexAddress(address) {
		return false
	}
	return !checksummed || common.HexToAddress(address).Hex() == address
}

// StringEthToWei converts a string representation of Ether to a *big.Int representation of Wei.
func StringEthToWei(amount string) *big.Int {
	ether, _, err := big.ParseFloat(amount, 10, 0, big.ToNearestEven)
	if err != nil {
		// Handle the error according to your needs.
		return nil
	}

	wei := new(big.Float).Mul(ether, big.NewFloat(1e18))
	result := new(big.Int)
	wei.Int(result) // Convert big.Float to big.Int
	return result
}
