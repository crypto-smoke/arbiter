package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter/uniswapv2"
	"github.com/ethereum/go-ethereum/common"
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
	router, err := uniswapv2.NewRouter(common.HexToAddress("0xB2C8D1e0626245A06BD5901063aB52f86F772BC4"), client)
	if err != nil {
		log.Err(err).Msg("failed")
		return
	}

	factoryAddy, err := router.Factory(nil)
	if err != nil {
		log.Err(err).Msg("failed")
		return
	}
	factory, err := uniswapv2.NewFactory(factoryAddy, client)
	if err != nil {
		log.Err(err).Msg("failed")
		return
	}
	pairs, err := factory.AllPairsLength(nil)
	if err != nil {
		log.Err(err).Msg("failed")
		return
	}
	fmt.Println("pairs", pairs.String())
}
