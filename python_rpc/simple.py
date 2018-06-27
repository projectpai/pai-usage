from bitcoinrpc.authproxy import AuthServiceProxy, JSONRPCException
rpc_connection = AuthServiceProxy("http://%s:%s@127.0.0.1:8566"%('user', 'pass'))

# create raw transaction

# here we can load raw transaction information from csv and pass it to paicoind-cli using bellow method
tx_hex_string = rpc_connection.createrawtransaction([{"txid":"dbdc2e2c7f143af70c5e7e8725f55d226b3c058d7bf34a303091b3c6a514848c","vout":1}]
                                    ,{"Mj8ntt66KnQniKvRny6o8kxiWgz5xhycGx": 0.00011})
print(tx_hex_string)

#signing raw transacion
rpc_connection.signrawtransaction(tx_hex_string, [{"txid": tx_hex_string,"vout": "%s", "scriptPubKey": "%s"}], ["private_key_data"])

# sending raw transaction
rpc_connection.sendrawtransaction(tx_hex_string)

### getting transaction fee

print(rpc_connection.gettransaction('transaction_id_which_we_want_to_use')['fee'])