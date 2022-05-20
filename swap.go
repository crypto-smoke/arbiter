package main

import (
	"github.com/crypto-smoke/arbiter/uniswapv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

type LiquidityPool struct {
	address        common.Address
	token0, token1 common.Address
	*uniswapv2.LP
}

func (lp *LiquidityPool) Address() common.Address {
	return lp.address
}
func (lp *LiquidityPool) Token0() common.Address {
	return lp.token0
}
func (lp *LiquidityPool) Token1() common.Address {
	return lp.token1
}
func NewLP(address common.Address, backend bind.ContractBackend) (*LiquidityPool, error) {
	return &LiquidityPool{}, nil
}

type Swap struct {
	routerAddress  common.Address
	factoryAddress common.Address
	*uniswapv2.Router
	*uniswapv2.Factory
}

func (s *Swap) shitfuck() {
	//	fee := 0.03
	//	tokensOut := (1/fee) * ( )
}
func (s *Swap) RouterAddress() common.Address {
	return s.routerAddress
}
func (s *Swap) FactoryAddress() common.Address {
	return s.factoryAddress
}
func NewSwap(router common.Address, backend bind.ContractBackend) (*Swap, error) {
	r, err := uniswapv2.NewRouter(router, backend)
	if err != nil {
		log.Err(err).Msg("failed to create router client")
		return nil, err
	}
	f, err := r.Factory(nil)
	if err != nil {
		log.Err(err).Msg("failed to get factory address")
		return nil, err
	}
	factory, err := uniswapv2.NewFactory(f, backend)
	if err != nil {
		log.Err(err).Msg("failed to create router client")
		return nil, err
	}
	return &Swap{
		routerAddress:  router,
		factoryAddress: f,
		Router:         r,
		Factory:        factory,
	}, nil
}
