package main

import (
	"github.com/projectpai/pai-usage/go-paicoin-multisig/multisig"

	"os"

	"gopkg.in/alecthomas/kingpin.v1"
)

// Kingpin configurations for command-line subcommands and their respective flags.
var (
	app = kingpin.New("go-paicoin-multisig", "A Paicoin multisig transaction builder")

	//keys subcommand
	cmdKeys        = app.Command("keys", "Generate public/private key pairs valid for use on Paicoin network. **PSEUDORANDOM AND FOR DEMONSTRATION PURPOSES ONLY. DO NOT USE IN PRODUCTION.**")
	cmdKeysCount   = cmdKeys.Flag("count", "No. of key pairs to generate.").Default("1").Int()
	cmdKeysConcise = cmdKeys.Flag("concise", "Turn on concise output. Default is off (verbose output).").Default("false").Bool()
	cmdKeysTestNet = cmdKeys.Flag("testnet", "Use testnet parameters. Default is off.").Default("false").Bool()
	//address subcommand
	cmdAddress           = app.Command("address", "Generate a multisig P2SH address with M-of-N requirements and set of public keys.")
	cmdAddressM          = cmdAddress.Flag("m", "M, the minimum number of keys needed to spend Paicoin in M-of-N multisig transaction.").Required().Int()
	cmdAddressN          = cmdAddress.Flag("n", "N, the total number of possible keys that can be used to spend Paicoin in M-of-N multisig transaction.").Required().Int()
	cmdAddressPublicKeys = cmdAddress.Flag("public-keys", "Comma separated list of private keys to sign with. Whitespace is stripped and quotes may be placed around keys. Eg. key1,key2,\"key3\"").PlaceHolder("PUBLIC-KEYS(Comma separated)").Required().String()
	cmdAddressTestNet    = cmdAddress.Flag("testnet", "Use testnet parameters. Default is off.").Default("false").Bool()
	//fund subcommand
	cmdFund            = app.Command("fund", "Fund multisig address from a standard Paicoin address.")
	cmdFundPrivateKey  = cmdFund.Flag("private-key", "Private key of paicoin to send.").Required().String()
	cmdFundInputTx     = cmdFund.Flag("input-tx", "Input transaction hash of paicoin to send.").Required().String()
	cmdFundAmount      = cmdFund.Flag("amount", "Amount of paicoin to send in satoshi (100,000,000 satoshi = 1 paicoin).").Required().Int()
	cmdFundDestination = cmdFund.Flag("destination", "Destination address. For P2SH, this should start with '3'.").Required().String()
	//spend subcommand
	cmdSpend             = app.Command("spend", "Spend multisig balance by sending to a standard Paicoin address.")
	cmdSpendPrivateKeys  = cmdSpend.Flag("private-keys", "Comma separated list of private keys to sign with. Whitespace is stripped and quotes may be placed around keys. Eg. key1,key2,\"key3\"").PlaceHolder("PRIVATE-KEYS(Comma separated)").Required().String()
	cmdSpendDestination  = cmdSpend.Flag("destination", "Public destination address to send paicoins.").Required().String()
	cmdSpendRedeemScript = cmdSpend.Flag("redeemScript", "Hex representation of redeem script that matches redeem script in P2SH input transaction.").Required().String()
	cmdSpendInputTx      = cmdSpend.Flag("input-tx", "Input transaction hash of paicoin to send.").Required().String()
	cmdSpendAmount       = cmdSpend.Flag("amount", "Amount of paicoin to send in satoshi (100,000,000 satoshi = 1 paicoin).").Required().Int()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	//keys -- Generate public/private key pairs
	case cmdKeys.FullCommand():
		multisig.OutputKeys(*cmdKeysCount, *cmdKeysConcise, *cmdKeysTestNet)

	//address -- Create a multisig P2SH address
	case cmdAddress.FullCommand():
		multisig.OutputAddress(*cmdAddressM, *cmdAddressN, *cmdAddressPublicKeys, *cmdAddressTestNet)

	//address -- Fund a P2SH address
	case cmdFund.FullCommand():
		multisig.OutputFund(*cmdFundPrivateKey, *cmdFundInputTx, *cmdFundAmount, *cmdFundDestination)

	//address -- Spend a multisig P2SH address
	case cmdSpend.FullCommand():
		multisig.OutputSpend(*cmdSpendPrivateKeys, *cmdSpendDestination, *cmdSpendRedeemScript, *cmdSpendInputTx, *cmdSpendAmount)
	}
}
