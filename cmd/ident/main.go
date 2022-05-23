package main

import (
	"bytes"
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/crypto-smoke/arbiter/arb"
	"github.com/crypto-smoke/arbiter/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var tokens = map[string]arbiter.Token{}
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
	tokensList = []string{
		"0x5C7F8A570d578ED84E63fdFA7b1eE72dEae1AE23", // wCRO
		"0x97749c9B61F878a880DfE312d2594AE07AEd7656", // MMF
		"0xc21223249CA28397B4B6541dfFaEcC539BfF0c59", // USDC
		"0x66e428c3f67a68878562e79A0234c1F83c208770", // USDT
		"0x95aEaF383E2e86A47c11CffdE1F7944eCB2C38C2", // MUSD
		"0xf2001b145b43032aaf5ee2884e456ccd805f677d", // DAI

	}
	/*
		pools = []string{
			"0xbA452A1c0875D33a440259B1ea4DcA8f5d86D9Ae", // MMF wCRO
			"0x722f19bd9A1E5bA97b3020c6028c279d27E4293C", // MMF USDC
			"0x5801d37e04ab1f266c35a277e06c9d3afa1c9ca2", // MMF USDT
			"0xeF2dC4849bDCC120acB7274cd5A557B5145DA149", // MMF MUSD
		}

	*/
)

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
		Start: []string{"usdc", "usdt", "musd"},
		Mid:   []string{"wcro", "mmf", "usdc", "usdt", "musd"}, // "crow"},
		//End:   []string{"usdc", "usdt", "musd","dai"},
	}
	for _, start := range paths.Start {
		for _, mid := range paths.Mid {
			if mid == start {
				continue
			}
			for _, mmid := range paths.Mid {
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
	//fmt.Println(routes)
}

type PoolList struct {
	sync.RWMutex
	p map[common.Address]map[common.Address]common.Address
}

func NewPoolList() *PoolList {
	return &PoolList{p: map[common.Address]map[common.Address]common.Address{}}
}
func (p *PoolList) sortAddresses(addressA, addressB common.Address) (common.Address, common.Address) {
	res := bytes.Compare(addressA.Bytes(), addressB.Bytes())
	if res > 0 {
		return addressB, addressA
	}
	return addressA, addressB
}
func (p *PoolList) SavePair(addressA, addressB, poolAddress common.Address) {
	p.Lock()
	defer p.Unlock()
	addressA, addressB = p.sortAddresses(addressA, addressB)
	//fmt.Println("SAVING:", addressA, addressB)

	if p.p[addressA] == nil {
		p.p[addressA] = make(map[common.Address]common.Address)
	}
	p.p[addressA][addressB] = poolAddress

}
func (p *PoolList) GetPair(addressA, addressB common.Address) common.Address {
	p.RLock()
	defer p.RUnlock()
	addressA, addressB = p.sortAddresses(addressA, addressB)
	return p.p[addressA][addressB]
}
func doThing(client *ethclient.Client) {
	util, err := utils.NewThing(common.HexToAddress("0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"), client)
	if err != nil {
		log.Err(err).Msg("failed dialing rpc")
		return
	}

	// returns false 894308421969433
	res, err := util.ComputeProfitMaximizingTrade(nil, big.NewInt(1000000000000000), big.NewInt(2000000000000000), big.NewInt(3000000000000000), big.NewInt(4000000000000000))
	if err != nil {
		log.Err(err).Msg("failed dialing rpc")
		return
	}
	fmt.Println(res)
	atob, amt := computeProfitMaximizingTrade(big.NewInt(1000000000000000), big.NewInt(2000000000000000), big.NewInt(3000000000000000), big.NewInt(4000000000000000))
	fmt.Println(atob, amt)
}
func aOrB(aToB bool, token1, token2 *big.Int) *big.Int {
	if aToB {
		return token1
	}
	return token2
}

//var fee int64 = 9970 // 9970 == 3% == uniswap default
var fee int64 = 9983 // 1.7% == meerkat default

func computeProfitMaximizingTrade(truePriceTokenA, truePriceTokenB, reserveA, reserveB *big.Int) (aToB bool, amountIn *big.Int) {
	//	aToB = FullMath.mulDiv(reserveA, truePriceTokenB, reserveB) < truePriceTokenA;
	mul := new(big.Int).Mul(reserveA, truePriceTokenB)
	div := new(big.Int).Div(mul, reserveB)
	//fmt.Println("atob math", div.String())
	aToB = div.Cmp(truePriceTokenA) == -1

	// uint256 invariant = reserveA.mul(reserveB);
	invariant := new(big.Int).Mul(reserveA, reserveB)
	//fmt.Println("invariant", invariant.String())
	/*uint256 leftSide = Babylonian.sqrt(
			FullMath.mulDiv(
				invariant.mul(1000),
				aToB ? truePriceTokenA : truePriceTokenB,
		(aToB ? truePriceTokenB : truePriceTokenA).mul(997)
		)
	);		*/
	invariant1000 := new(big.Int).Mul(invariant, big.NewInt(10000))
	mulDiv := new(big.Int).Mul(invariant1000, aOrB(aToB, truePriceTokenA, truePriceTokenB))
	mulDiv = new(big.Int).Div(mulDiv, new(big.Int).Mul(aOrB(aToB, truePriceTokenB, truePriceTokenA), big.NewInt(fee)))
	//fmt.Println("mdiv", mulDiv.String())
	leftSide := new(big.Int).Sqrt(mulDiv)

	//uint256 rightSide = (aToB ? reserveA.mul(1000) : reserveB.mul(1000)) / 997;
	rightSide := aOrB(aToB, new(big.Int).Mul(reserveA, big.NewInt(10000)), new(big.Int).Mul(reserveB, big.NewInt(10000)))
	rightSide = rightSide.Div(rightSide, big.NewInt(fee))
	//fmt.Println(leftSide.String(), rightSide.String())
	// if (leftSide < rightSide) return (false, 0);
	if leftSide.Cmp(rightSide) == -1 {
		return false, new(big.Int)
	}

	// compute the amount that must be sent to move the price to the profit-maximizing price
	//amountIn = leftSide.sub(rightSide);
	amountIn = leftSide.Sub(leftSide, rightSide)
	return
}
func main() { // appease the heroku gods
	go func() { _ = http.ListenAndServe(":"+os.Getenv("PORT"), nil) }()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: false,
	})

	client, err := ethclient.Dial("https://cronosrpc-2.xstaking.sg") //"https://rpc.nebkas.ro/")
	//client, err := ethclient.Dial("http://127.0.0.1:8545/") //"https://rpc.nebkas.ro/")
	if err != nil {
		log.Err(err).Msg("failed dialing rpc")
		return
	}
	//doThing(client)
	//return
	for _, t := range tokensList {
		tok, err := arbiter.NewToken(client, common.HexToAddress(t))
		if err != nil {
			log.Err(err).Msg("failed creating token")
			return
		}
		err = tok.Populate()
		if err != nil {
			log.Err(err).Msg("failed populating token")
			return
		}
		tokens[strings.ToLower(tok.Symbol())] = *tok
	}
	//fmt.Println(tokensList)
	// meerkat finance 0x145677FC4d9b8F19B5D56d1820c48e0443049a30
	swap, err := arbiter.NewSwap(common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"), client)
	if err != nil {
		log.Err(err).Msg("failed to create router client")
		return
	}

	poolList := NewPoolList()
	for _, t1 := range tokens {
		for _, t2 := range tokens {
			if t1 == t2 {
				continue
			}
			pair, err := swap.GetPair(nil, t1.Address(), t2.Address())
			if err != nil {
				log.Err(err).Msg("failed to get pair from factory")
				return
			}
			if isZeroAddress(pair) {
				//fmt.Printf("%s - %s: %s\n", t1.Symbol(), t2.Symbol(), "NO POOL")
				continue
			}
			existing := poolList.GetPair(t1.Address(), t2.Address())
			if !isZeroAddress(existing) {
				//fmt.Println("EXISTS:", existing)
				continue
			}
			poolList.SavePair(t1.Address(), t2.Address(), pair)
			fmt.Printf("%v - %v: %v\n", t1.Symbol(), t2.Symbol(), pair.String())
		}
	}
	//return
	// 0x5FbDB2315678afecb367f032d93F642f64180aa3 // localhost laptop
	// cronos mainnet 0xc98B790B00Bf8b4b3818a1432d7fAe64Fdee0aaB
	a, err := arb.NewArb(common.HexToAddress("0xc98B790B00Bf8b4b3818a1432d7fAe64Fdee0aaB"), client)
	if err != nil {
		log.Err(err).Msg("failed to create arb client")
		return
	}

	fmt.Println(routes)
	ticker := time.NewTicker(6 * time.Second)
	for ; true; <-ticker.C {
		for _, path := range routes {
			//go func(path []string) {
			canArb := checkTri(a, path, 100)
			if !canArb {
				continue
			}

			var r []arbiter.Token
			for _, t := range path {
				r = append(r, tokens[t])
			}
			var swaps arbiter.SwapChain
			for i := 0; i < len(r)-1; i++ {
				pool, err := swap.GetPair(nil, r[i].Address(), r[i+1].Address())
				if err != nil {
					log.Err(err).Msg("failed")
					continue
				}
				lp, err := arbiter.NewLP(pool, client)
				if err != nil {
					log.Err(err).Msg("failed")
					continue
				}
				reserves, err := lp.GetReserves(nil)
				if err != nil {
					log.Err(err).Msg("failed")
					continue
				}

				reserves0 := reserves.Reserve0
				reserves1 := reserves.Reserve1

				var inputReserves, outputReserves *big.Int
				inputAddress := r[i].Address()
				if lp.Token0() == inputAddress {
					inputReserves = reserves0
					outputReserves = reserves1
				} else if lp.Token1() == inputAddress {
					inputReserves = reserves1
					outputReserves = reserves0
				} else {
					panic("wtf")
				}

				if err != nil {
					log.Err(err).Msg("failed getting reserves 1")
				}
				var token0, token1 string
				for name, address := range coins {
					if address == lp.Token0() {
						token0 = name
					}
					if address == lp.Token1() {
						token1 = name
					}
				}
				swaps = append(swaps, arbiter.NewSingleSwap(inputReserves.String(), outputReserves.String(), 0.0017))
				fmt.Printf("%s: %s - %s: %s\n", token0, reserves0.String(), token1, reserves1.String())
			}
			amountIn, profit := swaps.FindInput()
			amountInInt, _ := big.NewFloat(amountIn).Int(nil)
			profitInt, _ := big.NewFloat(profit).Int(nil)

			fmt.Printf("Best input: %v will profit %v\n", r[0].ToFloat64(amountInInt), r[0].ToFloat64(profitInt))
			/*
				continue
				reserves0, reserves1, err := swap.GetReserves(r[0].Address(), r[1].Address())
				if err != nil {
					log.Err(err).Msg("failed getting reserves 1")
				}
				reserves2, reserves3, err := swap.GetReserves(r[1].Address(), r[2].Address())
				if err != nil {
					log.Err(err).Msg("failed getting reserves 2")
				}

				atob, amountIn := computeProfitMaximizingTrade(reserves0, reserves1, reserves2, reserves3)
				fmt.Println(atob, amountIn, reserves0.String(), reserves1.String(), reserves2.String(), reserves3.String())
				inputToken := r[0]
				if atob {
					inputToken = r[2]
				}
				fmt.Printf("input %0.2f %s\n", inputToken.ToFloat64(amountIn), inputToken.Symbol())

			*/
			//	}(path)
			//	go checkDual(a, r)
		}
	}
}

func isZeroAddress(a common.Address) bool {
	return a == common.Address{0}
}
func checkTri(a *arb.Arb, path []string, inputAmount float64) bool {
	var r []arbiter.Token
	for _, t := range path {
		r = append(r, tokens[t])
	}
	//fmt.Println(r)

	//amount := new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(r[0].Decimals())), nil))
	amount := r[0].FromFloat64(inputAmount)
	//amount = big.NewInt(10000000000)
	output, err := a.EstimateTriDexTrade(nil,
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),
		common.HexToAddress("0x145677FC4d9b8F19B5D56d1820c48e0443049a30"),

		r[0].Address(),
		r[1].Address(),
		r[2].Address(),
		amount,
	)

	if err != nil {
		log.Err(err).Msg("failed to create arb client")
		return false
	}

	//fmt.Println("checking tri", r[0], r[1], r[2], r[0], balToFloat(amount, int(inputDecimals)), balToFloat(output, int(outputDecimals)))
	if balToFloat(output, r[0].Decimals())-balToFloat(amount, r[0].Decimals()) > 0.3 {
		//if balToFloat(new(big.Int).Sub(output, amount), 6) >= 1 {
		log.Info().
			Float64("input", balToFloat(amount, r[0].Decimals())).
			Float64("output", balToFloat(output, r[0].Decimals())).
			Msgf("%v -> %v -> %v -> %v", r[0].Symbol(), r[1].Symbol(), r[2].Symbol(), r[0].Symbol())

		return true
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
	return false
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
