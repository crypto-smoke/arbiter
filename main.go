package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter/arb"
	"github.com/crypto-smoke/arbiter/erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math"
	"math/big"
	"net/http"
	"os"
	"time"
)

var coins = map[string]common.Address{
	"wcro": common.HexToAddress("0x5C7F8A570d578ED84E63fdFA7b1eE72dEae1AE23"), // wCRO
	"mmf":  common.HexToAddress("0x97749c9B61F878a880DfE312d2594AE07AEd7656"), // MMF
	"crow": common.HexToAddress("0x285c3329930a3fd3C7c14bC041d3E50e165b1517"),
	"usdc": common.HexToAddress("0xc21223249CA28397B4B6541dfFaEcC539BfF0c59"), // USDC
	"usdt": common.HexToAddress("0x66e428c3f67a68878562e79A0234c1F83c208770"), // USDT
	"musd": common.HexToAddress("0x95aEaF383E2e86A47c11CffdE1F7944eCB2C38C2"), // MUSD
	"dai":  common.HexToAddress("0xf2001b145b43032aaf5ee2884e456ccd805f677d"), //DAI
}

var (
	tokens = []string{
		"0x5C7F8A570d578ED84E63fdFA7b1eE72dEae1AE23", // wCRO
		"0x97749c9B61F878a880DfE312d2594AE07AEd7656", // MMF
		"0xc21223249CA28397B4B6541dfFaEcC539BfF0c59", // USDC
		"0x66e428c3f67a68878562e79A0234c1F83c208770", // USDT
		"0x95aEaF383E2e86A47c11CffdE1F7944eCB2C38C2", // MUSD
		"0xf2001b145b43032aaf5ee2884e456ccd805f677d", // DAI

	}

	pools = []string{
		"0xbA452A1c0875D33a440259B1ea4DcA8f5d86D9Ae", // MMF wCRO
		"0x722f19bd9A1E5bA97b3020c6028c279d27E4293C", // MMF USDC
		"0x5801d37e04ab1f266c35a277e06c9d3afa1c9ca2", // MMF USDT
		"0xeF2dC4849bDCC120acB7274cd5A557B5145DA149", // MMF MUSD
	}
)

type Coin struct {
	Address  common.Address
	Decimals int
}

type arbs struct {
	Start   []string
	Mid     []string
	MoreMid []string
	End     []string
}

var routes [][]string

func init() {
	/* 	paths := arbs{
	   		Start: []string{"usdc", "usdt", "musd", "dai"},
	   		Mid:   []string{"wcro", "mmf"}, // "crow"},
	   		End:   []string{"usdc", "usdt", "musd"},
	   	}

	*/

	paths := arbs{
		Start: []string{"usdc", "usdt", "musd", "wcro", "mmf"},
		//Mid:   []string{"wcro", "mmf"}, // "crow"},
		//End:   []string{"usdc", "usdt", "musd","dai"},
	}
	for _, start := range paths.Start {
		for _, mid := range paths.Start {
			if mid == start {
				continue
			}
			for _, mmid := range paths.Start {
				if mmid == start {
					continue
				}
				if mmid == mid {
					continue
				}

				routes = append(routes, []string{start, mid, mmid, start})

			}
		}
	}
	/*
		for _, start := range paths.Start {
			for _, mid := range paths.Mid {
				for _, mmid := range paths.Mid {
					if mmid == mid {
						continue
					}
					for _, end := range paths.End {
						if start == end {
							continue
						}
						routes = append(routes, []string{start, mid, mmid, end})
					}
				}
			}
		}

	*/
	fmt.Println(routes)
}
func main() { // appease the heroku gods
	go func() { _ = http.ListenAndServe(":"+os.Getenv("PORT"), nil) }()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: true,
	})

	client, err := ethclient.Dial("https://cronosrpc-2.xstaking.sg") //"https://rpc.nebkas.ro/")
	if err != nil {
		log.Err(err).Msg("failed dialing rpc")
		return
	}
	a, err := arb.NewArb(common.HexToAddress("0xc98B790B00Bf8b4b3818a1432d7fAe64Fdee0aaB"), client)
	if err != nil {
		log.Err(err).Msg("failed to create arb client")
		return
	}
	/*
		router, err := uniswapv2.NewRouter(common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"), client)
		if err != nil {
			log.Err(err).Msg("failed to create router client")
			return
		}

	*/

	ticker := time.NewTicker(6 * time.Second)
	for ; true; <-ticker.C {
		for _, r := range routes {
			go checkTri(a, r)
			//	go checkDual(a, r)
		}
	}
}

func checkTri(a *arb.Arb, r []string) {
	var inputDecimals int64 = 6
	if r[0] == "musd" || r[0] == "dai" {
		inputDecimals = 18
	}
	amount := new(big.Int).Mul(big.NewInt(1000), new(big.Int).Exp(big.NewInt(10), big.NewInt(inputDecimals), nil))
	//
	output, err := a.EstimateTriDexTrade(nil,
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),

		coins[r[0]],
		coins[r[1]],
		coins[r[2]],
		amount,
	)

	if err != nil {
		log.Err(err).Msg("failed to create arb client")
		return
	}

	outputDecimals := inputDecimals
	//fmt.Println("checking tri", r[0], r[1], r[2], r[0], balToFloat(amount, int(inputDecimals)), balToFloat(output, int(outputDecimals)))
	if balToFloat(output, int(outputDecimals))-balToFloat(amount, int(inputDecimals)) > 0.5 {
		//if balToFloat(new(big.Int).Sub(output, amount), 6) >= 1 {
		log.Info().
			Float64("input", balToFloat(amount, int(inputDecimals))).
			Float64("output", balToFloat(output, int(outputDecimals))).
			Msgf("%v -> %v -> %v -> %v", r[0], r[1], r[2], r[0])
		/*
			txn, err := a.DualDexTrade(&bind.TransactOpts{
				From:      common.Address{},
				Nonce:     nil,
				Signer:    nil,
				Value:     nil,
				GasPrice:  nil,
				GasFeeCap: nil,
				GasTipCap: nil,
				GasLimit:  0,
				Context:   nil,
				NoSend:    true,
			},
				common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
				common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
			)

		*/
	}
}

func checkDual(a *arb.Arb, r []string) {
	var inputDecimals int64 = 6
	if r[0] == "musd" {
		inputDecimals = 18
	}
	amount := new(big.Int).Mul(big.NewInt(1000), new(big.Int).Exp(big.NewInt(10), big.NewInt(inputDecimals), nil))

	output, err := a.EstimateSingleDexTrade(nil,
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),

		coins[r[0]],
		coins[r[1]],
		coins[r[2]],
		amount,
	)

	if err != nil {
		log.Err(err).Msg("failed to create arb client")
		return
	}

	var outputDecimals int64 = 6
	if r[2] == "musd" {
		outputDecimals = 18
	}

	if balToFloat(output, int(outputDecimals))-balToFloat(amount, int(inputDecimals)) > 1 {
		//if balToFloat(new(big.Int).Sub(output, amount), 6) >= 1 {
		log.Info().
			Float64("input", balToFloat(amount, int(inputDecimals))).
			Float64("output", balToFloat(output, int(outputDecimals))).
			Msgf("%v -> %v -> %v", r[0], r[1], r[2])
		/*
			txn, err := a.DualDexTrade(&bind.TransactOpts{
				From:      common.Address{},
				Nonce:     nil,
				Signer:    nil,
				Value:     nil,
				GasPrice:  nil,
				GasFeeCap: nil,
				GasTipCap: nil,
				GasLimit:  0,
				Context:   nil,
				NoSend:    true,
			},
				common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
				common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
			)

		*/
	}
}

func balToFloat(balance *big.Int, decimals int) float64 {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(decimals)))
	actBal, _ := ethValue.Float64()
	return actBal
}
