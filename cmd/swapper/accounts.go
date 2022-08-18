package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/crypto-smoke/arbiter/cmd/swapper/input"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"path/filepath"
	"strings"
)

/*
 if term.IsTerminal(int(syscall.Stdin)) {
        fmt.Println("Terminal is interactive! You're good to use prompts!")
    } else {
        fmt.Println("Terminal is not interactive! Consider using flags or environment variables!")
    }
*/
func importAccount(ks *keystore.KeyStore) error {
	var key *ecdsa.PrivateKey
	for {
		keyInput := input.Password("Paste private key:")
		if strings.HasPrefix(keyInput, "0x") {
			keyInput = strings.TrimPrefix(keyInput, "0x")
		}
		keyBytes, err := hex.DecodeString(keyInput)
		if err != nil {
			log.Fatal().Err(err).Msg("failed importing key")
		}
		key, err = crypto.ToECDSA(keyBytes)
		if err != nil {
			log.Fatal().Err(err).Msg("failed creating ecdsa key")
		}

		fmt.Println("Imported account with address:", crypto.PubkeyToAddress(key.PublicKey))
		if input.YesNo("Is this the correct account?", false) {
			break
		}
	}
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
	fmt.Println("importing key")
	_, err = ks.ImportECDSA(key, passphrase)
	if err != nil {
		log.Fatal().Err(err).Msg("failed saving new key")
	}
	fmt.Println("key imported")
	return nil
}
func loadWallet() *accounts.Account {
	directory := utils.HomeDir()
	ks := keystore.NewKeyStore(filepath.Join(directory, ".smokeswap/keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
	//	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	ksAccounts := ks.Accounts()

	if len(ksAccounts) == 0 {
		fmt.Println("No wallets saved. You can create a new account or import a private key.")
		if input.YesNo("Create new account?", true) {
			err := createNewAccount(ks)
			if err != nil {
				log.Fatal().Err(err)
			}
		} else {
			err := importAccount(ks)
			if err != nil {
				log.Fatal().Err(err)
			}
		}
	}
	ksAccounts = ks.Accounts()
	switch len(ksAccounts) {
	case 0:
		log.Fatal().Msg("no accounts saved")
	case 1:
		err := ks.Unlock(ksAccounts[0], input.Password("Enter passphrase for account "+ksAccounts[0].Address.String()))
		if err != nil {
			log.Fatal().Err(err).Msg("failed unlocking account")
		}
		fmt.Println("account unlocked")
		return &ksAccounts[0]
	default:
		log.Fatal().Msg("too many accounts")
	}
	return nil
}

func shit() *accounts.Account {
	directory := utils.HomeDir()
	ks := keystore.NewKeyStore(filepath.Join(directory, ".smokeswap/keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
	//	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	ksAccounts := ks.Accounts()

	if len(ksAccounts) == 0 {
		fmt.Println("No wallets saved. You can create a new account or import a private key.")
		if input.YesNo("Create new account?", true) {
			err := createNewAccount(ks)
			if err != nil {
				log.Fatal().Err(err)
			}
		} else {
			var key *ecdsa.PrivateKey
			for {
				keyInput := input.Password("Paste private key:")
				if strings.HasPrefix(keyInput, "0x") {
					keyInput = strings.TrimPrefix(keyInput, "0x")
				}
				keyBytes, err := hex.DecodeString(keyInput)
				if err != nil {
					log.Fatal().Err(err).Msg("failed importing key")
				}
				key, err = crypto.ToECDSA(keyBytes)
				if err != nil {
					log.Fatal().Err(err).Msg("failed creating ecdsa key")
				}

				fmt.Println("Imported account with address:", crypto.PubkeyToAddress(key.PublicKey))
				if input.YesNo("Is this the correct account?", false) {
					break
				}
			}
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
			fmt.Println("importing key")
			_, err = ks.ImportECDSA(key, passphrase)
			if err != nil {
				log.Fatal().Err(err).Msg("failed saving new key")
			}
			fmt.Println("key imported")
		}
	}
	ksAccounts = ks.Accounts()
	switch len(ksAccounts) {
	case 0:
		log.Fatal().Msg("no accounts saved")
	case 1:
		err := ks.Unlock(ksAccounts[0], input.Password("Enter passphrase for account "+ksAccounts[0].Address.String()))
		if err != nil {
			log.Fatal().Err(err).Msg("failed unlocking account")
		}
		fmt.Println("account unlocked")
		return &ksAccounts[0]
	default:
		log.Fatal().Msg("too many accounts")
	}
	return nil
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
	fmt.Println("Account created: " + newAcc.Address.String())
	return nil
}
