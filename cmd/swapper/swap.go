package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math/big"
)

type Env struct {
	client *ethclient.Client
}

func NewEnv(c *ethclient.Client) (*Env, error) {
	return &Env{client: c}, nil
}

func (e *Env) swapHandler(c *gin.Context) {

}
func Token0isBase(baseToken common.Address, lp *arbiter.LiquidityPool) bool {
	return lp.Token0().Address() == baseToken
}
func (e *Env) getPrice(c *gin.Context) {
	baseTokenAddresss := common.HexToAddress(c.Query("base"))
	quoteTokenAddresss := common.HexToAddress(c.Query("quote"))
	routerAddresss := common.HexToAddress(c.Query("router"))
	fmt.Println(baseTokenAddresss, quoteTokenAddresss, routerAddresss)
	r, err := arbiter.NewSwap(routerAddresss, e.client)
	if err != nil {
		log.Err(err).Msg("failed creating swap")
		c.AbortWithError(500, err)
		return
	}
	lp, err := r.GetLP(baseTokenAddresss, quoteTokenAddresss)
	if err != nil {
		log.Err(err).Msg("failed creating lp")
		c.AbortWithError(500, err)
		return
	}
	reserves, err := lp.GetReserves(nil)
	if err != nil {
		log.Err(err).Msg("failed getting reserves")
		c.AbortWithError(500, err)
		return
	}

	var baseReserves, quoteReserves *big.Int
	var quoteToken, baseToken *arbiter.Token
	if Token0isBase(baseTokenAddresss, lp) {
		baseReserves = reserves.Reserve0
		quoteReserves = reserves.Reserve1
		quoteToken = lp.Token1()
		baseToken = lp.Token0()
	} else {
		baseReserves = reserves.Reserve1
		quoteReserves = reserves.Reserve0
		quoteToken = lp.Token0()
		baseToken = lp.Token1()
	}
	_ = quoteToken
	_ = baseToken
	midPrice := baseToken.ToFloat64(baseReserves) / quoteToken.ToFloat64(quoteReserves)
	//value, err := r.GetAmountsOut(nil, big.NewInt(1), []common.Address{baseTokenAddresss, quoteTokenAddresss})
	/*value, err := r.Quote(nil, big.NewInt(1), baseReserves, quoteReserves)
	if err != nil {
		log.Err(err).Msg("failed getting quote")
		c.AbortWithError(500, err)
		return
	}

	*/
	fmt.Println(baseReserves.String(), quoteReserves.String())
	fmt.Println(midPrice)
	//fmt.Println(value[0].String(), value[1].String())
	//value := new(big.Int).Div(baseReserves, quoteReserves)
	//fmt.Println(baseToken.ToFloat64(value))
	//fmt.Println(baseToken.ToFloat64(value[1]))
}
