package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/http"
)

type Env struct {
	client   *ethclient.Client
	keystore *keystore.KeyStore
	account  *accounts.Account
	cfg      *Config
}

type Config struct {
	WalletAddress             common.Address
	Base, Quote               *arbiter.Token
	BaseBalance, QuoteBalance float64
	RPC                       string
	Router                    common.Address
	Slippage                  float64
	GasGwei                   int64
}

func NewEnv(c *ethclient.Client, rpcURL string, ks *keystore.KeyStore, acc *accounts.Account) (*Env, error) {
	return &Env{
		client:   c,
		keystore: ks,
		account:  acc,
		cfg: &Config{
			WalletAddress: acc.Address,
			RPC:           rpcURL,
		},
	}, nil
}
func (e *Env) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"address":           e.cfg.WalletAddress,
		"baseTokenAddress":  e.cfg.Base.Address(),
		"quoteTokenAddress": e.cfg.Quote.Address(),
		"slippage":          e.cfg.Slippage,
		"gwei":              e.cfg.GasGwei,
		"routerAddress":     e.cfg.Router,
		"rpcURL":            e.cfg.RPC,
	})
}
func (e *Env) swapHandler(c *gin.Context) {

}
func Token0isBase(baseToken common.Address, lp *arbiter.LiquidityPool) bool {
	return lp.Token0().Address() == baseToken
}

type UpdateRequest struct {
	Address    string  `json:"address,omitempty"`
	Base       string  `json:"base,omitempty"`
	Quote      string  `json:"quote,omitempty"`
	Router     string  `json:"router,omitempty"`
	Slippage   float64 `json:"slippage,omitempty"`
	Gwei       int     `json:"gwei,omitempty"`
	GetPrice   bool    `json:"get_price,omitempty"`
	GetBalance bool    `json:"get_balance,omitempty"`
}
type UpdateResponse struct {
	BaseBalance  float64 `json:"base_balance,omitempty"`
	QuoteBalance float64 `json:"quote_balance,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

func (e *Env) getInitialData(c *gin.Context) {

}
func (e *Env) getUpdate(c *gin.Context) {
	var request UpdateRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Err(err).Msg("failed to bind data")
		c.AbortWithError(500, err)
		return
	}

	var response UpdateResponse

	if request.GetPrice {
		price, err := e.getPrice(
			common.HexToAddress(request.Base),
			common.HexToAddress(request.Quote),
			common.HexToAddress(request.Router))
		if err != nil {
			log.Err(err).Msg("failed to get price")
			c.AbortWithError(500, err)
			return
		}
		response.Price = price
	}
	if request.GetBalance {
		baseBal, quoteBal, err := e.getBalances(
			common.HexToAddress(request.Base),
			common.HexToAddress(request.Quote),
			common.HexToAddress(request.Address))
		if err != nil {
			log.Err(err).Msg("failed to get balances")
			c.AbortWithError(500, err)
			return
		}
		response.QuoteBalance = quoteBal
		response.BaseBalance = baseBal
	}

	c.JSON(200, response)
}
func (e *Env) getBalances(baseTokenAddress, quoteTokenAddress, walletAddress common.Address) (float64, float64, error) {
	base, err := arbiter.NewToken(e.client, baseTokenAddress)
	if err != nil {
		return 0.0, 0.0, err
	}
	quote, err := arbiter.NewToken(e.client, quoteTokenAddress)
	if err != nil {
		return 0.0, 0.0, err
	}

	baseBal, err := base.BalanceOf(nil, walletAddress)
	if err != nil {
		return 0.0, 0.0, err
	}

	quoteBal, err := quote.BalanceOf(nil, walletAddress)
	if err != nil {
		return 0.0, 0.0, err
	}

	return base.ToFloat64(baseBal), quote.ToFloat64(quoteBal), nil
}

func (e *Env) getPrice(baseTokenAddress, quoteTokenAddress, routerAddress common.Address) (float64, error) {
	fmt.Println(baseTokenAddress, quoteTokenAddress, routerAddress)
	r, err := arbiter.NewSwap(routerAddress, e.client)
	if err != nil {
		return 0.0, err
	}
	lp, err := r.GetLP(baseTokenAddress, quoteTokenAddress)
	if err != nil {
		return 0.0, err
	}
	reserves, err := lp.GetReserves(nil)
	if err != nil {
		return 0.0, err
	}

	var baseReserves, quoteReserves *big.Int
	var quoteToken, baseToken *arbiter.Token
	if Token0isBase(baseTokenAddress, lp) {
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

	midPrice := baseToken.ToFloat64(baseReserves) / quoteToken.ToFloat64(quoteReserves)
	return midPrice, nil
}

func (e *Env) getPriceHandler(c *gin.Context) {
	baseTokenAddress := common.HexToAddress(c.Query("base"))
	quoteTokenAddress := common.HexToAddress(c.Query("quote"))
	routerAddress := common.HexToAddress(c.Query("router"))

	midPrice, err := e.getPrice(baseTokenAddress, quoteTokenAddress, routerAddress)
	if err != nil {
		log.Err(err).Msg("failed getting price")
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, midPrice)

}
