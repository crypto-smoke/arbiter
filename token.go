package main

import (
	"github.com/crypto-smoke/arbiter/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

type Token struct {
	Address   common.Address
	Decimals  big.Int
	Name      string
	Symbol    string
	Token     *erc20.Token
	populated bool
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

	dec := new(big.Int).SetUint64(uint64(decimals))
	t.Decimals = *dec

	t.Name = name

	t.Symbol = symbol
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
		Address: a,
		Token:   t,
	}, nil
}
