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

func (t *Token) FromFloat64(amount float64) *big.Int {
	f := big.NewFloat(amount)
	exp := Pow(big.NewFloat(10), uint64(t.decimals))
	f = f.Mul(f, exp)
	i, _ := f.Int(nil)
	return i
}

func Pow(a *big.Float, e uint64) *big.Float {
	result := Zero().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
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
