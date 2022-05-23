package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: false,
	})

	client, err := ethclient.Dial("https://cronosrpc-2.xstaking.sg")
	if err != nil {
		log.Err(err).Msg("failed dialing rpc")
		return
	}
	_ = client

	swaps := arbiter.SwapChain{
		arbiter.NewSingleSwap("100", "200", 0.0017),
		arbiter.NewSingleSwap("190", "100", 0.0017),
		arbiter.NewSingleSwap("100", "100", 0.0017),
	}

	amountIn, profit := swaps.FindInput()
	fmt.Println("ideal input:", amountIn, "proft:", profit)

}
