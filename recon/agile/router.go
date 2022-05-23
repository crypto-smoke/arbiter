package agile

import (
	"github.com/crypto-smoke/arbiter/swap"
	"github.com/crypto-smoke/arbiter/swap/driver"
	"github.com/ethereum/go-ethereum/common"
)

var Router = common.HexToAddress("0x9AadB4A3BFaCff8aa60a2c63735Cb8B94De7C57d")

var name = "agile"

type Conn struct{}

func (c *Conn) Close() {}

type Driver struct{}

func (d *Driver) New() driver.Conn {
	return &Conn{}
}
func init() {
	swap.Register(name, &Driver{})
}
