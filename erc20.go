package shell

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

	"github.com/dl-nice/go-ethereum-contract/contract/erc20"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployErc20(client *ethclient.Client, auth *bind.TransactOpts, tokenName, tokenSymbol string, initialSupply *big.Int) (contractAddress string, hash string, err error) {
	address, tx, _, err := erc20.DeployErc20(auth, client, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return "", "", err
	}
	contractAddress = address.Hex()
	hash = tx.Hash().Hex()
	return
}
