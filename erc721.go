package shell

import (
	"context"
	"github.com/dl-nice/go-ethereum-contract/contract/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func DeployErc721(ctx context.Context, client *ethclient.Client, tokenName, tokenSymbol, pvk string, limit uint64) (contractAddress, hash string, err error) {
	auth, err := getAuth(ctx, client, pvk)
	if err != nil {
		return "", "", err
	}
	auth.GasLimit = limit
	address, tx, _, err := erc721.DeployErc721(auth, client, tokenName, tokenSymbol)
	if err != nil {
		return "", "", err
	}
	contractAddress = address.Hex()
	hash = tx.Hash().Hex()
	return
}

func AwardItemErc721(ctx context.Context, client *ethclient.Client, contractAddress, fromAddress, cid, pvk string, limit uint64) (hash string, tokenId *big.Int, err error) {
	auth, err := getAuth(ctx, client, pvk)
	if err != nil {
		return "", nil, err
	}
	auth.GasLimit = limit
	nft, err := erc721.NewErc721(common.HexToAddress(contractAddress), client)
	if err != nil {
		return "", nil, err
	}
	address := common.HexToAddress(fromAddress)
	tx, err := nft.AwardItem(auth, address, cid)
	if err != nil {
		return "", nil, err
	}
	raw := erc721.Erc721Raw{Contract: nft}
	var out []interface{}
	err = raw.Call(&bind.CallOpts{From: address}, &out, "awardItem", common.HexToAddress(fromAddress), cid)
	if err != nil {
		return "", nil, err
	}
	return tx.Hash().Hex(), out[0].(*big.Int), nil
}

func TokenURIErc721(client *ethclient.Client, contractAddress, fromAddress string, tokenId *big.Int) (string, error) {
	nft, err := erc721.NewErc721(common.HexToAddress(contractAddress), client)
	if err != nil {
		return "", nil
	}
	return nft.TokenURI(&bind.CallOpts{From: common.HexToAddress(fromAddress)}, tokenId)
}
