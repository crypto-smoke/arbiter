package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
)

type Env struct {
	client   *ethclient.Client
	chainID  *big.Int
	keystore *keystore.KeyStore
	account  accounts.Account
	cfg      *Config
}

func NewEnv(c *ethclient.Client, rpcURL string, ks *keystore.KeyStore, acc accounts.Account) (*Env, error) {
	chainID, err := c.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed getting chain ID")
	}
	return &Env{
		client:   c,
		chainID:  chainID,
		keystore: ks,
		account:  acc,
		cfg: &Config{
			WalletAddress: acc.Address,
			RPC:           rpcURL,
		},
	}, nil
}
