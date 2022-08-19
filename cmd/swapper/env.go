package main

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Env struct {
	client   *ethclient.Client
	keystore *keystore.KeyStore
	account  accounts.Account
	cfg      *Config
}

func NewEnv(c *ethclient.Client, rpcURL string, ks *keystore.KeyStore, acc accounts.Account) (*Env, error) {
	return &Env{
		client:   c,
		keystore: ks,
		account:  acc,
		cfg: &Config{
			WalletAddress: acc.Address,
			RPC:           rpcURL,
		},
	}, nil
}
