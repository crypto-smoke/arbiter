package main

import (
	"github.com/crypto-smoke/arbiter/uniswapv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"math/big"
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
	lp, err := uniswapv2.NewLP(address, backend)
	if err != nil {
		return nil, err
	}
	t0, err := lp.Token0(nil)
	if err != nil {
		return nil, err
	}
	t1, err := lp.Token1(nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPool{
		address: address,
		token0:  t0,
		token1:  t1,
		LP:      lp,
	}, nil
}

type Swap struct {
	routerAddress  common.Address
	factoryAddress common.Address
	*uniswapv2.Router
	*uniswapv2.Factory
	backend bind.ContractBackend
}

func (s *Swap) GetReserves(token0, token1 common.Address) (reserves0, reserves1 *big.Int, err error) {
	pool, err := s.GetPair(nil, token0, token1)
	if err != nil {
		return
	}
	lp, err := NewLP(pool, s.backend)
	if err != nil {
		return
	}
	reserves, err := lp.GetReserves(nil)
	if err != nil {
		return
	}
	reserves0 = reserves.Reserve0
	reserves1 = reserves.Reserve1
	return
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
		backend:        backend,
	}, nil
}
