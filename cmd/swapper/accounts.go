package main

import (
	"fmt"
	"github.com/crypto-smoke/arbiter/cmd/swapper/input"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"path/filepath"
)

/*
 if term.IsTerminal(int(syscall.Stdin)) {
        fmt.Println("Terminal is interactive! You're good to use prompts!")
    } else {
        fmt.Println("Terminal is not interactive! Consider using flags or environment variables!")
    }
*/
func shit() {
	directory := utils.HomeDir()
	ks := keystore.NewKeyStore(filepath.Join(directory, ".smokeswap/keystore"), keystore.LightScryptN, keystore.LightScryptN)
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	wallets := am.Wallets()
	if len(wallets) == 0 {
		if input.YesNo("Create new account?", true) {
			err := createNewAccount(ks)
			if err != nil {
				log.Fatal().Err(err)
			}
		}
	}
	fmt.Println(len(wallets), "wallets")
}

var PassphraseMismatch = errors.New("passphrases dont match")

func getPassword() (string, error) {
	first := input.Password("Enter wallet passphrase:")
	second := input.Password("Enter passphrase again:")
	if first != second {
		return "", PassphraseMismatch
	}
	return first, nil
}
func createNewAccount(ks *keystore.KeyStore) error {
	var passphrase string
	var err error
	for {
		passphrase, err = getPassword()
		if err != nil {
			fmt.Println(err)
			if input.YesNo("Try again?", false) {
				continue
			}
			break
		}
		break
	}

	newAcc, err := ks.NewAccount(passphrase)
	if err != nil {
		return errors.Wrap(err, "failed creating new account")
	}
	fmt.Println(newAcc)
	return nil
}
