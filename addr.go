package shell

import (
	"context"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ValidAddr(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func ValidContractAddr(ctx context.Context, client *ethclient.Client, address string) bool {
	if !ValidAddr(address) {
		return false
	}
	addr := common.HexToAddress(address)
	byteCode, err := client.CodeAt(ctx, addr, nil)
	if err != nil {
		return false
	}
	return len(byteCode) > 0
}
