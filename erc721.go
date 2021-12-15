package shell

import (
	"github.com/dl-nice/go-ethereum-contract/contract/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func DeployErc721(client *ethclient.Client, auth *bind.TransactOpts, tokenName, tokenSymbol string) (contractAddress, hash string, err error) {
	address, tx, _, err := erc721.DeployErc721(auth, client, tokenName, tokenSymbol)
	if err != nil {
		return "", "", err
	}
	contractAddress = address.Hex()
	hash = tx.Hash().Hex()
	return
}

func AwardItemErc721(client *ethclient.Client, auth *bind.TransactOpts, contractAddress, fromAddress, cid string) (hash string, tokenId *big.Int, err error) {
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

	}
	return nft.TokenURI(&bind.CallOpts{From: common.HexToAddress(fromAddress)}, tokenId)
}
