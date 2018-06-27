// keys.go - Generating public/private key pairs.
package multisig

import (
	"github.com/prettymuchbryce/hellobitcoin/base58check"
	"github.com/projectpai/pai-usage/go-paicoin-multisig/btcutils"

	"encoding/hex"
	"fmt"
	"log"
)

//OutputKeys formats and prints relevant outputs to the user.
func OutputKeys(flagKeyCount int, flagConcise bool, flagTestNet bool) {
	if flagKeyCount < 1 || flagKeyCount > 100 {
		log.Fatal("--count <count> must be between 1 and 100")
	}

	if !flagConcise {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("Disclaimer: These key pairs are cryptographically secure to the limits of the crypto/rand cryptography package in Golang. They should not be used without further security audit in production systems.")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("Each generated key pair includes: ")
		fmt.Println("* Your private key\t\t\t-- Keep this private, needed to spend received Paicoins.")
		fmt.Println("* Your public key\t\t\t-- in HEX format. This is required to generate multisig destination address.")
		fmt.Println("* Your public destination address\t-- Give this to other people to send you Paicoins.")
		fmt.Println("----------------------------------------------------------------------")
	}

	privateKeyWIFs, publicKeyHexs, publicAddresses := generateKeys(flagKeyCount, flagTestNet)

	for i := 0; i <= flagKeyCount-1; i++ {

		//Output private key in WIF format, public key as hex and P2PKH public address
		fmt.Println("-------------------------------------------------------------")
		fmt.Printf("KEY #%d\n", i+1)
		if !flagConcise {
			fmt.Println("")
		}
		fmt.Println("Private key: ")
		fmt.Println(privateKeyWIFs[i])
		if !flagConcise {
			fmt.Println("")
		}
		fmt.Println("Public key hex: ")
		fmt.Println(publicKeyHexs[i])
		if !flagConcise {
			fmt.Println("")
		}
		fmt.Println("Public Paicoin address: ")
		fmt.Println(publicAddresses[i])
		fmt.Println("-------------------------------------------------------------")
	}
}

// generateKeys is the high-level logic for generating public/private key pairs with the 'go-paicoin-multisig keys' subcommand.
// Takes flagCount (desired number of key pairs) and flagConcise (true hides warnings and helpful messages for conciseness)
// as arguments.
func generateKeys(flagKeyCount int, flagTestNet bool) ([]string, []string, []string) {
	publicKeyHexs := make([]string, flagKeyCount)
	publicAddresses := make([]string, flagKeyCount)
	privateKeyWIFs := make([]string, flagKeyCount)

	for i := 0; i <= flagKeyCount-1; i++ {
		//Generate private key
		privateKey := btcutils.NewPrivateKey()
		//Generate public key from private key
		publicKey, err := btcutils.NewPublicKey(privateKey)
		if err != nil {
			log.Fatal(err)
		}
		//Get hex encoded version of public key
		publicKeyHexs[i] = hex.EncodeToString(publicKey)
		//Get public address by hashing with SHA256 and RIPEMD160 and base58 encoding with prefix
		publicKeyHash, err := btcutils.Hash160(publicKey)
		if err != nil {
			log.Fatal(err)
		}
		publicAddresses[i] = base58check.Encode(Prefix(PublicKey, flagTestNet), publicKeyHash)
		//Get private key in Wallet Import Format (WIF) by base58 encoding with prefix
		privateKeyWIFs[i] = base58check.Encode(Prefix(PrivateKey, flagTestNet), privateKey)
	}

	return privateKeyWIFs, publicKeyHexs, publicAddresses
}
