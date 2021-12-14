package shell

import (
	"context"
	"math/big"

	"github.com/dl-nice/go-ethereum-contract/contract/erc20"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployErc20(ctx context.Context, client *ethclient.Client, tokenName string, tokenSymbol string, initialSupply *big.Int, pvk string, limit uint64) (contractAddress string, hash string, err error) {
	auth, err := getAuth(ctx, client, pvk)
	if err != nil {
		return "", "", err
	}
	auth.GasLimit = limit
	address, tx, _, err := erc20.DeployErc20(auth, client, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return "", "", err
	}
	contractAddress = address.Hex()
	hash = tx.Hash().Hex()
	return
}
