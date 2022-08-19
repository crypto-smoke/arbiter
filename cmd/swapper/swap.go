package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	WalletAddress             common.Address
	Base, Quote               *arbiter.Token
	BaseBalance, QuoteBalance float64
	RPC                       string
	Router                    common.Address
	Slippage                  float64
	GasGwei                   int64
}

type SwapData struct {
	RPCURL        string           `json:"rpc_url,omitempty"`
	RouterAddress string           `json:"router_address,omitempty"`
	Address       string           `json:"address,omitempty"`    // address of wallet to send tx from
	SwapPath      []common.Address `json:"swap_path,omitempty"`  // path of token swap, should be []string{base,quote} if buy
	AmountIn      float64          `json:"amount_in,omitempty"`  // input amount of the first token in the path
	AmountOut     float64          `json:"amount_out,omitempty"` // output amount of the last token in the path (slippage not included)
	GasGwei       int64            `json:"gas_gwei,omitempty"`
	Slippage      float64          `json:"slippage,omitempty"` // slippage in percent (1% == 1.0)
	IsBuy         bool             `json:"is_buy,omitempty"`
}

func (e *Env) checkApproval(amt *big.Int) bool {
	allowance, err := e.cfg.Base.Allowance(nil, e.cfg.WalletAddress, e.cfg.Router)
	if err != nil {
		panic(err)
	}
	fmt.Println("amount:", amt.String())
	fmt.Println("approval:", allowance.String())
	// if the allowance is greater than or equal to the amount we want to send, return true
	return allowance.Cmp(amt) >= 0
}
func (e *Env) getTxnOpts() *bind.TransactOpts {
	chainID, err := e.client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	opts, err := bind.NewKeyStoreTransactorWithChainID(e.keystore, e.account, chainID)
	if err != nil {
		panic(err)
	}
	return opts
}
func (e *Env) submitApproval(token *arbiter.Token, spender common.Address, amount *big.Int) {
	opts := e.getTxnOpts()
	if amount == nil {
		b, err := hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
		if err != nil {
			panic(err)
		}
		amount = new(big.Int).SetBytes(b)

	}
	fmt.Println("submitting approval for", amount.String())
	tx, err := token.Approve(opts, spender, amount)
	if err != nil {
		panic(err)
	}
	fmt.Println("approval sent")
	rcpt, err := bind.WaitMined(context.Background(), e.client, tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("mined")
	if rcpt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("success")
	}
}
func (e *Env) swapHandler(c *gin.Context) {
	var data SwapData
	err := c.BindJSON(&data)
	if err != nil {
		panic(err)
	}
	if strings.ToLower(e.cfg.RPC) != strings.ToLower(data.RPCURL) {
		// update client and reconnect
		fmt.Println("new rpc")
	}
	var amountIn, amountOut *big.Int
	var inToken, outToken *arbiter.Token
	if data.IsBuy {
		amountIn = e.cfg.Base.FromFloat64(data.AmountIn)
		amountOut = e.cfg.Quote.FromFloat64(data.AmountOut - (data.AmountOut * (data.Slippage / 100)))
		inToken = e.cfg.Base
		outToken = e.cfg.Quote
	} else {
		amountIn = e.cfg.Quote.FromFloat64(data.AmountIn)
		amountOut = e.cfg.Base.FromFloat64(data.AmountOut)
		inToken = e.cfg.Quote
		outToken = e.cfg.Base
	}
	pretty.Print(data)
	s, err := arbiter.NewSwap(common.HexToAddress(data.RouterAddress), e.client)
	if err != nil {
		panic(err)
	}

	fmt.Println("pre", amountOut.String())
	amountsOut, err := s.GetAmountsOut(nil, amountIn, data.SwapPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed getting amounts out")
	}
	for i, a := range amountsOut {
		fmt.Println(i, ":", a.String())
	}
	amountOut = amountsOut[len(amountsOut)-1]
	fmt.Println("post", amountOut.String())

	_ = outToken
	isApproved := e.checkApproval(amountOut)
	if !isApproved {
		fmt.Println("need approval")
		e.submitApproval(inToken, e.cfg.Router, nil)
		return
	}
	chainID, err := e.client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	opts, err := bind.NewKeyStoreTransactorWithChainID(e.keystore, e.account, chainID)
	if err != nil {
		panic(err)
	}
	opts.GasPrice = new(big.Int).Mul(big.NewInt(data.GasGwei), big.NewInt(1000000000))

	fmt.Println(amountIn.String(), amountOut.String())

	tx, err := s.SwapExactTokensForTokens(opts, amountIn, amountOut, data.SwapPath, e.cfg.WalletAddress, big.NewInt(time.Now().Add(5*time.Minute).Unix()))
	if err != nil {
		panic(err)
	}
	rcpt, err := bind.WaitMined(context.Background(), e.client, tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("done", rcpt.Status, rcpt.TxHash)
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
