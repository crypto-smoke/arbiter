package arbiter

import (
	"github.com/crypto-smoke/arbiter/uniswapv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
)

type LiquidityPool struct {
	address        common.Address
	token0, token1 common.Address
	t0, t1         *Token
	*uniswapv2.LP
}

func (lp *LiquidityPool) Address() common.Address {
	return lp.address
}
func (lp *LiquidityPool) Token0() *Token {
	return lp.t0
}
func (lp *LiquidityPool) Token1() *Token {
	return lp.t1
}
func NewLP(address common.Address, backend bind.ContractBackend) (*LiquidityPool, error) {
	uniswapLP, err := uniswapv2.NewLP(address, backend)
	if err != nil {
		return nil, err
	}
	t0, err := uniswapLP.Token0(nil)
	if err != nil {
		return nil, err
	}
	t1, err := uniswapLP.Token1(nil)
	if err != nil {
		return nil, err
	}

	lp := &LiquidityPool{
		address: address,
		token0:  t0,
		token1:  t1,
		LP:      uniswapLP,
	}
	lp.t0, err = NewToken(backend, t0)
	if err != nil {
		return nil, err
	}
	lp.t1, err = NewToken(backend, t1)
	if err != nil {
		return nil, err
	}
	return lp, nil
}

type Swap struct {
	routerAddress  common.Address
	factoryAddress common.Address
	*uniswapv2.Router
	*uniswapv2.Factory
	backend bind.ContractBackend
}

func (s *Swap) GetLP(token0, token1 common.Address) (*LiquidityPool, error) {
	lpAddr, err := s.GetPair(nil, token0, token1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pair")
	}
	return NewLP(lpAddr, s.backend)
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
