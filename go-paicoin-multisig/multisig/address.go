// Package multisig contains the main starting threads for each of the subcommands for go-paicoin-multisig.
//
// address.go - Generating P2SH addresses.
package multisig

import (
	"github.com/prettymuchbryce/hellobitcoin/base58check"
	"github.com/projectpai/pai-usage/go-paicoin-multisig/btcutils"

	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

type PrefixType byte

const (
	PublicKey PrefixType = iota
	P2shScript
	PrivateKey
)

func Prefix(prefix PrefixType, flagTestNet bool) string {
	if flagTestNet {
		switch prefix {
			case PublicKey:
				return "33"
			case P2shScript:
				return "B4"
			case PrivateKey:
				return "E2"
		}
	} else {
		switch prefix {
			case PublicKey:
				return "38"
			case P2shScript:
				return "82"
			case PrivateKey:
				return "F7"
		}
	}
	return "00"
}

//OutputAddress formats and prints relevant outputs to the user.
func OutputAddress(flagM int, flagN int, flagPublicKeys string, flagTestNet bool) {
	P2SHAddress, redeemScriptHex := generateAddress(flagM, flagN, flagPublicKeys, flagTestNet)

	if flagM*73+flagN*66 > 496 {
		fmt.Printf(`
-----------------------------------------------------------------------------------------------------------------------------------
WARNING: 
%d-of-%d multisig transaction is valid but *non-standard* for Bitcoin v0.9.x and earlier.
It may take a very long time (possibly never) for transaction spending multisig funds to be included in a block.
To remain valid, choose smaller m and n values such that m*73+n*66 <= 496, as per standardness rules.
See http://bitcoin.stackexchange.com/questions/23893/what-are-the-limits-of-m-and-n-in-m-of-n-multisig-addresses for more details.
------------------------------------------------------------------------------------------------------------------------------------
`,
			flagM,
			flagN,
		)
	}
	//Output P2SH and redeemScript
	fmt.Printf(`
-----------------------------------------------------------------------------------------------------------------------------------
Your *P2SH ADDRESS* is:
%v
Give this to sender funding multisig address with Paicoin.
-----------------------------------------------------------------------------------------------------------------------------------
-----------------------------------------------------------------------------------------------------------------------------------
Your *REDEEM SCRIPT* is:
%v
Keep private and provide this to redeem multisig balance later.
-----------------------------------------------------------------------------------------------------------------------------------
`,
		P2SHAddress,
		redeemScriptHex,
	)
}

// generateAddress is the high-level logic for creating P2SH multisig addresses with the 'go-paicoin-multisig address' subcommand.
// Takes flagM (number of keys required to spend), flagN (total number of keys)
// and flagPublicKeys (comma separated list of N public keys) as arguments.
func generateAddress(flagM int, flagN int, flagPublicKeys string, flagTestNet bool) (string, string) {
	//Convert public keys argument into slice of public key bytes with necessary tidying
	flagPublicKeys = strings.Replace(flagPublicKeys, "'", "\"", -1) //Replace single quotes with double since csv package only recognizes double quotes
	publicKeyStrings, err := csv.NewReader(strings.NewReader(flagPublicKeys)).Read()
	if err != nil {
		log.Fatal(err)
	}
	publicKeys := make([][]byte, len(publicKeyStrings))
	for i, publicKeyString := range publicKeyStrings {
		publicKeyString = strings.TrimSpace(publicKeyString)   //Trim whitespace
		publicKeys[i], err = hex.DecodeString(publicKeyString) //Get private keys as slice of raw bytes
		if err != nil {
			log.Fatal(err, "\n", "Offending publicKey: \n", publicKeyString)
		}
	}
	//Create redeemScript from public keys
	redeemScript, err := btcutils.NewMOfNRedeemScript(flagM, flagN, publicKeys)
	if err != nil {
		log.Fatal(err)
	}
	redeemScriptHash, err := btcutils.Hash160(redeemScript)
	if err != nil {
		log.Fatal(err)
	}
	//Get P2SH address by base58 encoding with P2SH prefix
	P2SHAddress := base58check.Encode(Prefix(P2shScript, flagTestNet), redeemScriptHash)
	//Get redeemScript in Hex
	redeemScriptHex := hex.EncodeToString(redeemScript)

	return P2SHAddress, redeemScriptHex
}
