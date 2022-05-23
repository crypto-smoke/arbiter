package arbiter

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gonum.org/v1/gonum/optimize"
	"math/big"
)

func NewSingleSwap(inputReserves, outputReserves string, fee float64) SingleSwap {
	return newswap(inputReserves, outputReserves, fee)
}
func newswap(inputReserves, outputReserves string, fee float64) SingleSwap {
	i, _ := new(big.Int).SetString(inputReserves, 10)
	o, _ := new(big.Int).SetString(outputReserves, 10)
	return SingleSwap{
		Fee:                 fee,
		InputTokenReserves:  i,
		OutputTokenReserves: o,
	}
}

func (s SwapChain) FindInput() (idealInput float64, profit float64) {
	p := optimize.Problem{
		Func: s.Minimize,
	}

	x := []float64{1}
	result, err := optimize.Minimize(p, x, nil, nil)
	if err != nil {
		log.Fatal().Err(err)
	}
	if err = result.Status.Err(); err != nil {
		log.Fatal().Err(err)
	}
	fmt.Printf("result.Status: %v\n", result.Status)
	fmt.Printf("result.X: %0.4g\n", result.X)
	fmt.Printf("result.F: %0.4g\n", result.F)
	fmt.Printf("time:%v\n", result.Runtime)
	fmt.Printf("result.Stats.FuncEvaluations: %d\n", result.Stats.FuncEvaluations)
	fmt.Println(result.F * -1)
	return result.X[0], result.F * -1
}

func Sauce() float64 {
	swaps := []SingleSwap{
		newswap("19620506524995", "19588298297531", 0.0017),
		newswap("3537938437212", "15899803120072376261728251", 0.0017),
		newswap("13630671474867253004808338", "3055055887382", 0.0017),
	}
	return sauce(100, swaps)
}

/*
6:53AM INF USDT -> USDC -> MMF -> USDT input=100 output=100.042479
usdt: 19620506524995 - usdc: 19588298297531
mmf: 15899803120072376261728251 - usdc: 3537938437212
usdt: 3055055887382 - mmf: 13630671474867253004808338

*/
type SingleSwap struct {
	Fee                 float64 //0.003 for uniswap .3%
	InputTokenReserves  *big.Int
	OutputTokenReserves *big.Int
}

type SwapChain []SingleSwap

func (s SwapChain) Minimize(initialIn []float64) (profit float64) {
	outputAmount := sauce(initialIn[0], []SingleSwap(s))
	//fmt.Println(initialIn, outputAmount)
	return -1 * (outputAmount - initialIn[0])
}
func MinSauce(initialIn []float64) (profit float64) {
	/* swaps := []SingleSwap{
		newswap("19620506524995", "19588298297531", 0.0017),
		newswap("3537938437212", "15899803120072376261728251", 0.0017),
		newswap("13630671474867253004808338", "3055055887382", 0.0017),
	}

	*/
	swaps := []SingleSwap{
		newswap("100", "200", 0.0017),
		newswap("190", "100", 0.0017),
		newswap("100", "100", 0.0017),
	}
	//fmt.Println(initialIn)
	outputAmount := sauce(initialIn[0], swaps)
	fmt.Println(initialIn, outputAmount)
	return -1 * (outputAmount - initialIn[0])
}

// sauce returns the amount of tokens expected after executing the swaps
func sauce(initialIn float64, swaps []SingleSwap) float64 {
	/*
			def func(initialIn):
			  global n,fee,IpX0,IpX1
			  initalIn=abs(initialIn)
			  swapout=initialIn
			  for i in range(n):
			      swapout=((1-fee[i]) * swapout * IpX0[i] ) / (IpX1[i] + (1-fee[i]) * swapout)
			  return -1.0*(swapout-initialIn)

		(0.997 * amountIn * reserveOut) / (reserveIn + 0.997 * amountIn)
	*/
	swapOutput := initialIn
	for _, s := range swaps {

		//swapOutput = ((1 - s.Fee) * swapOutput * s.OutputTokenReserves) / (s.InputTokenReserves. + (1-s.Fee)*swapOutput)
		reservesOut, outAcc := new(big.Float).SetInt(s.OutputTokenReserves).Float64()
		_ = outAcc
		//	fmt.Println(s.OutputTokenReserves.String(), "to float", reservesOut, "with acc", outAcc)
		reservesIn, inAcc := new(big.Float).SetInt(s.InputTokenReserves).Float64()
		_ = inAcc
		//	fmt.Println(s.InputTokenReserves.String(), "to float", reservesIn, "with acc", inAcc)
		swapOutput = ((1 - s.Fee) * swapOutput * reservesOut) / (reservesIn + (1-s.Fee)*swapOutput)
	}
	return swapOutput
	return -1.0 * (swapOutput - initialIn)
}
