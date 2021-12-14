package shell

import (
	"strings"

	api "github.com/ipfs/go-ipfs-api"
)

type Client struct {
	IpfsCli *api.Shell
}

func NewClient(uri string) *Client {
	return &Client{
		IpfsCli: api.NewShell(uri),
	}
}

func (i *Client) Add(str string) (hash string, err error) {
	hash, err = i.IpfsCli.Add(strings.NewReader(str))
	return
}
