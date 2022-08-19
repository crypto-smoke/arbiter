package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	directory := utils.HomeDir()
	ks := keystore.NewKeyStore(filepath.Join(directory, ".smokeswap/keystore"), keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := loadWallet(ks)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load account")
	}

	fmt.Println("Using account", account.Address.String())
	rpcURL := GetEnvOrPanic("RPC_URL")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed dialing rpc")
	}
	env, err := NewEnv(client, rpcURL, ks, *account)
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating env")
	}
	baseToken, err := arbiter.NewToken(client, common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating base token")
	}
	quoteToken, err := arbiter.NewToken(client, common.HexToAddress("0x22a31bD4cB694433B6de19e0aCC2899E553e9481"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating quote token")
	}
	env.cfg = &Config{
		WalletAddress: env.cfg.WalletAddress,
		Base:          baseToken,
		Quote:         quoteToken,
		BaseBalance:   0,
		QuoteBalance:  0,
		RPC:           env.cfg.RPC,
		Router:        common.HexToAddress("0x51aBA405De2b25E5506DeA32A6697F450cEB1a17"),
		Slippage:      69,
		GasGwei:       420,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)
	r.Static("/static", "./ui")

	r.LoadHTMLFiles("./ui/index.html")
	r.GET("/", env.indexHandler)
	r.POST("/swap", env.swapHandler)
	r.GET("/tokenPrice", env.getPriceHandler)
	r.POST("/update", env.getUpdate)

	log.Info().Msg("open web interface at http://localhost:8069/")

	r.Run("localhost:8069") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func oldmain() {
	fs := http.FileServer(http.Dir("./ui"))
	http.Handle("/", http.StripPrefix("/", fs))

	//http.FileServer(http.FS(res))
	log.Info().Msg("open web interface at http://localhost:8069/")
	log.Fatal().Err(http.ListenAndServe(":8069", nil))

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: false,
	})

	client, err := ethclient.Dial(GetEnvOrPanic("RPC_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed dialing rpc")
	}
	_ = client

	s, err := arbiter.NewSwap(common.HexToAddress(GetEnvOrPanic("ROUTER_ADDRESS")), client)
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating swap")
	}
	_ = s
	//getPairs(s)

	baseToken, err := arbiter.NewToken(client, common.HexToAddress(GetEnvOrPanic("BASE_TOKEN")))
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating base token")
	}
	log.Info().Str("name", baseToken.Name()).Str("symbol", baseToken.Symbol()).Msg("base token loaded")
	quoteToken, err := arbiter.NewToken(client, common.HexToAddress(GetEnvOrPanic("QUOTE_TOKEN")))
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating quote token")
	}
	log.Info().Str("name", quoteToken.Name()).Str("symbol", quoteToken.Symbol()).Msg("quote token loaded")

	accountAddress := GetEnvOrPanic("ACCOUNT_ADDRESS")
	for {
		baseBalance, err := baseToken.BalanceOf(nil, common.HexToAddress(accountAddress))
		if err != nil {
			log.Fatal().Err(err).Msg("failed getting base token balance")
		}

		log.Info().Str("symbol", baseToken.Symbol()).
			Float64("balance", baseToken.ToFloat64(baseBalance)).
			Msg("base token balance")

		quoteBalance, err := quoteToken.BalanceOf(nil, common.HexToAddress(accountAddress))
		if err != nil {
			log.Fatal().Err(err).Msg("failed getting quote token balance")
		}

		log.Info().Str("symbol", quoteToken.Symbol()).
			Float64("balance", quoteToken.ToFloat64(quoteBalance)).
			Msg("quote token balance")

		lp, err := s.GetLP(quoteToken.Address(), baseToken.Address())
		if err != nil {
			log.Fatal().Err(err).Msg("failed getting lp")
		}

		reserves, err := lp.GetReserves(nil)
		if err != nil {
			log.Fatal().Err(err).Msg("failed getting lp reserves")
		}
		_ = reserves

		if lp.Token0().Address() == baseToken.Address() {
			log.Info().Msg("token 0 is base")
		} else if lp.Token1().Address() == baseToken.Address() {
			log.Info().Msg("token 1 is base")
			//fmt.Println(quoteBalance.String(), baseBalance.String())
			//fmt.Println(reserves.Reserve0.String(), reserves.Reserve1.String())
			//baseBalance = baseToken.FromFloat64(1000000)

			amountsOut, err := s.GetAmountsOut(nil, baseBalance, []common.Address{baseToken.Address(), quoteToken.Address()})
			//amountsOut, err := s.GetAmountsOut(nil, baseBalance, []common.Address{baseToken.Address(), common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"), quoteToken.Address()})

			if err != nil {
				log.Fatal().Err(err).Msg("failed getting amounts out")
			}
			log.Info().
				Float64("usdc in", baseToken.ToFloat64(baseBalance)).
				Float64("mmf out", quoteToken.ToFloat64(amountsOut[len(amountsOut)-1])).
				Float64("USDC per MMF", baseToken.ToFloat64(baseBalance)/quoteToken.ToFloat64(amountsOut[len(amountsOut)-1])).
				Msg("USDC to MMF")

		}
		break

	}

}

func getPairs(s *arbiter.Swap) {
	pairs, err := s.AllPairsLength(nil)
	if err != nil {
		log.Fatal().Err(err).Msg("failed getting pairs")
	}
	log.Info().Int64("pair count", pairs.Int64()).Msg("got pairs")
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic("environment variable " + key + " not set")
	}
	return value
}
