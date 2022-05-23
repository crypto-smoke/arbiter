package database

import "github.com/ethereum/go-ethereum/common"

type Database interface {
	GetDex(name string) (common.Address, error)
}
