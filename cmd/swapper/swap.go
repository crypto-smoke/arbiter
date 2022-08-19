package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/http"
	"strconv"
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
	RouterAddress common.Address   `json:"router_address,omitempty"`
	Address       common.Address   `json:"address,omitempty"`    // address of wallet to send tx from
	SwapPath      []common.Address `json:"swap_path,omitempty"`  // path of token swap, should be []string{base,quote} if buy
	AmountIn      float64          `json:"amount_in,omitempty"`  // input amount of the first token in the path
	AmountOut     float64          `json:"amount_out,omitempty"` // output amount of the last token in the path (slippage not included)
	GasGwei       int64            `json:"gas_gwei,omitempty"`
	Slippage      float64          `json:"slippage,omitempty"` // slippage in percent (1% == 1.0)
	IsBuy         bool             `json:"is_buy,omitempty"`
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
func (e *Env) checkApproval(token *arbiter.Token, spender common.Address, amount *big.Int) bool {
	allowance, err := token.Allowance(nil, e.cfg.WalletAddress, spender)
	if err != nil {
		panic(err)
	}
	fmt.Println("amount:", amount.String())
	fmt.Println("approval:", allowance.String())
	// if the allowance is greater than or equal to the amount we want to send, return true
	return allowance.Cmp(amount) >= 0
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
		amountOut = e.cfg.Base.FromFloat64(data.AmountOut - (data.AmountOut * (data.Slippage / 100)))
		inToken = e.cfg.Quote
		outToken = e.cfg.Base
	}
	s, err := arbiter.NewSwap(data.RouterAddress, e.client)
	if err != nil {
		panic(err)
	}

	amountsOut, err := s.GetAmountsOut(nil, amountIn, data.SwapPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed getting amounts out")
	}

	amountOut = amountsOut[len(amountsOut)-1]

	_ = outToken
	isApproved := e.checkApproval(inToken, e.cfg.Router, amountOut)
	if !isApproved {
		fmt.Println("need approval")
		e.submitApproval(inToken, e.cfg.Router, nil)
	}

	opts, err := bind.NewKeyStoreTransactorWithChainID(e.keystore, e.account, e.chainID)
	if err != nil {
		panic(err)
	}

	// only set gas price if we specified something
	if data.GasGwei != 0 {
		opts.GasFeeCap = new(big.Int).Mul(big.NewInt(data.GasGwei), big.NewInt(1000000000))
		opts.GasTipCap = opts.GasFeeCap
	}

	//opts.GasPrice = suggestedGas //new(big.Int).Mul(big.NewInt(data.GasGwei), big.NewInt(1000000000))
	//opts.GasLimit = 500000
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
	BaseBalance  json.Number `json:"base_balance"`
	BaseName     string      `json:"base_name"`
	BaseSymbol   string      `json:"base_symbol"`
	QuoteBalance json.Number `json:"quote_balance"`
	QuoteName    string      `json:"quote_name"`
	QuoteSymbol  string      `json:"quote_symbol"`
	Price        json.Number `json:"price"`
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
		response.Price = json.Number(strconv.FormatFloat(price, 'g', 18, 64))

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
	response.BaseName = e.cfg.Base.Name()
	response.BaseSymbol = e.cfg.Base.Symbol()
	response.QuoteName = e.cfg.Quote.Name()
	response.QuoteSymbol = e.cfg.Quote.Symbol()

	c.JSON(200, response)
}
func (e *Env) getBalances(baseTokenAddress, quoteTokenAddress, walletAddress common.Address) (json.Number, json.Number, error) {
	base, err := arbiter.NewToken(e.client, baseTokenAddress)
	if err != nil {
		return "0", "0", err
	}
	quote, err := arbiter.NewToken(e.client, quoteTokenAddress)
	if err != nil {
		return "0", "0", err
	}

	baseBal, err := base.BalanceOf(nil, walletAddress)
	if err != nil {
		return "0", "0", err
	}

	quoteBal, err := quote.BalanceOf(nil, walletAddress)
	if err != nil {
		return "0", "0", err
	}
	baseFloat := base.ToFloat64(baseBal)

	quoteFloat := quote.ToFloat64(quoteBal)
	basefmtString := fmt.Sprintf("%%.%vf", base.Decimals())
	quotefmtString := fmt.Sprintf("%%.%vf", quote.Decimals())
	return json.Number(fmt.Sprintf(basefmtString, baseFloat)),
		//	strconv.FormatFloat(baseFloat, 'e', base.Decimals(), 64)),
		json.Number(fmt.Sprintf(quotefmtString, quoteFloat)),
		nil
}

func (e *Env) getPrice(baseTokenAddress, quoteTokenAddress, routerAddress common.Address) (float64, error) {
	//	fmt.Println(baseTokenAddress, quoteTokenAddress, routerAddress)
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
