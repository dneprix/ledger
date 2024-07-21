# ledger
**ledger** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```
`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.


```
 Blockchain is running
  
  👤 alice's account address: cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd
  👤 bob's account address: cosmos1hpdvdzxhxsgfzxntm4rr2fqu90hryxwjx35e7m
  
  🌍 Tendermint node: http://0.0.0.0:26657
  🌍 Blockchain API: http://0.0.0.0:1317
  🌍 Token faucet: http://0.0.0.0:4500
```

# Run tests

Now you can run the tests to ensure that everything is working correctly:

```
go test ./x/ledger/keeper -v
```

# Ledger module

```
~/go/bin/ledgerd tx ledger create-account cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd 10coin --from alice 
~/go/bin/ledgerd tx ledger update-account cosmos1hpdvdzxhxsgfzxntm4rr2fqu90hryxwjx35e7m 10coin --from alice 

~/go/bin/ledgerd tx ledger update-account cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd 100coin --from alice 

curl -X GET "http://0.0.0.0:1317/dneprix/ledger/ledger/account" -H  "accept: application/json"
{
  "account": [
    {
      "address": "cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd",
      "balance": {
        "denom": "coin",
        "amount": "100"
      },
      "creator": "cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd"
    },
    {
      "address": "cosmos1hpdvdzxhxsgfzxntm4rr2fqu90hryxwjx35e7m",
      "balance": {
        "denom": "coin",
        "amount": "10"
      },
      "creator": "cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd"
    },
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}

curl -X GET "http://0.0.0.0:1317/dneprix/ledger/ledger/account/cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd" -H  "accept: application/json"
{
  "account": {
    "address": "cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd",
    "balance": {
      "denom": "coin",
      "amount": "100"
    },
    "creator": "cosmos15u9nl5n4524lc8zp36r6qvcqlqsma4jxx02fhd"
  }
}

```