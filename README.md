# PAI Example Code Usage
Overall PAI usage is very similar to Bitcoin, the RPC layer is the same and the P2P is very similar differing only in the protocol specifics like genesis hash, magic bytes, address formats, etc.

## Global settings
###P2P Ports:
```
8567 - Mainnet
18567 - Testnet
19567 - Regtest
```

###Default RPC Ports
```
8566 - Mainnet
18566 - Testnet
19566 - Regtest
```

## Python via RPC
In order to interact with paicoind via RPC, you can use any of the same tools that work with Bitcoin.  In this example we use bitcoinrpc (https://github.com/jgarzik/python-bitcoinrpc) which is a simple wrapper layer to the standard Bitcoin RPC.  All of the expected RPC requests are still present and can be seen with `paicoin-cli help`.

Please refer to `python_rpc/simple.py` for this example.  You can use the included `requirements.txt` file to get `python-bitcoinrpc` installed via `pip`.


## Java P2P Example
For peer-to-peer usage and Java we have modified `bitcoinj` to fit the requirements of PAI, including all of the various network settings.  The repository: https://github.com/projectpai/bitcoinj/tree/paicoin on the paicoin branch contains this project as well as a new module called `pai-demo` which has a usage example with extensive comments.  In general though, anything that can be achieved with the standard `bitcoinj` should also be possible with this modified version as the only relevant changes are in the `NetworkParams` section.
Please refer to the `README` here: https://github.com/projectpai/bitcoinj/blob/paicoin/paidemo/src/main/java/com/oben/README.md

