package main

import (
	"github.com/crypto-smoke/arbiter/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

type Token struct {
	address  common.Address
	name     string
	symbol   string
	decimals int
	*erc20.Token
	populated bool
}

func (t *Token) Address() common.Address {
	return t.address
}
func (t *Token) Name() string {
	return t.name
}
func (t *Token) Symbol() string {
	return t.symbol
}
func (t *Token) Decimals() int {
	return t.decimals
}
func (t *Token) ToFloat64(balance *big.Int) float64 {
	return balToFloat(balance, t.decimals)
}
func (t *Token) populate() error {
	decimals, err := t.Token.Decimals(nil)
	if err != nil {
		return err
	}
	name, err := t.Token.Name(nil)
	if err != nil {
		return err
	}

	symbol, err := t.Token.Symbol(nil)
	if err != nil {
		return err
	}

	t.decimals = int(decimals)

	t.name = name

	t.symbol = symbol
	t.populated = true
	return nil
}

func NewToken(backend bind.ContractBackend, a common.Address) (*Token, error) {
	if backend == nil {
		return nil, errors.New("nil client")
	}
	t, err := erc20.NewToken(a, backend)
	if err != nil {
		return nil, err
	}
	return &Token{
		address: a,
		Token:   t,
	}, nil
}
