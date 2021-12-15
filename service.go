package shell

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	api "github.com/ipfs/go-ipfs-api"
	"log"
	"math/big"
	"strings"
)

func Mint(auth *bind.TransactOpts, contractAddress, eUrl, iUrl, fromAddress, json string) (hash string, tokenId *big.Int, err error) {
	ethClient, err := ethclient.Dial(eUrl)
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewShell(iUrl)
	cid, err := client.Add(strings.NewReader(json))
	return AwardItemErc721(ethClient, auth, contractAddress, fromAddress, cid)
}
