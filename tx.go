package shell

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransactionStatusReceipt(ctx context.Context, client *ethclient.Client, hash string) (bool, error) {
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(hash))
	if err != nil {
		return false, err
	}
	if receipt.Status != 1 {
		return false, nil
	}
	return true, nil
}
