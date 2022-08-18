package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const AddressToCheck = "0xfe2a94ecd89745f3C3099067f9738296941DbC9D"

func main() {
	client, err := ethclient.Dial("https://arb1.arbitrum.io/rpc")
	if err != nil {
		panic(err)
	}

	m, err := NewMarinate(common.HexToAddress("0x2adabd6e8ce3e82f52d9998a7f64a90d294a92a4"), client)
	if err != nil {
		panic(err)
	}
	values, err := m.MarinatorInfo(nil, common.HexToAddress(AddressToCheck))
	if err != nil {
		panic(err)
	}
	fmt.Println("Checking rewards for mUMAMI holder", AddressToCheck)
	fmt.Println("Amount", values.Amount.String())
	fmt.Println("MultipliedAmount", values.MultipliedAmount.String())
}
