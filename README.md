# CNS Client
CNS(Chain Naming Service) client is connected to CNS gateway (TEST VERSION)

## Introduction
Most addresses of blockchain are difficult to recognize as address is generated randomly by using key algorithm like secp256k1 and is too long. For easily reminding address, Naming service is that short name(as 'domain') is mapped to blockchain account address similarly DNS. This client supports to operate methods are creating/retrieving domain and send coin using domain (not account address) in the CNS system as frontend. It is inspired by [ENS](https://ens.domains/).
This CNS client, but also CNS gateway and CNS Cosmwasm contract composed of CNS system, is able to operate on Tendermint based blockchains. Also, it supports chains included ethermint module for using EVM by selecting secp256k1 is used to general cosmos/tendermint and eth_secp256k1 is used to ethermint module. 

## Prerequisites
- Set config file
  - `./config/config.json`
  ```yaml
  {
    "clientPort":"13000",
    "gatewayEndpoint":"http://localhost:12000"
  }
  ```
  - `clientPort` : CNS client TCP port number
  - `gatewayEndpoint` : Target CNS gateway information (IP address and port number) 
  
## Start
```shell
go mod tidy
go build cnsclient.go
./cnsclient
```

## Usage
First of all, This client is test version and not wallet. For testing, user's mnemonic words is needed to use CNS system. The user private key is generated by mnemonic words, and is saved CNS gateway session DB. In session end, gateway deletes info in the DB.
After CNS client runs, user can check functionalities in the browser. 

### Create wallet(private key and account address)
At `Wallet` tap, an user can create mnemonic words or recover private key and account address. After recover address, it might be different user's already known address as `coinType` of HD derivation path in order to create private key set 60(ETH) not 118(Cosmos) in the CNS gateway.

### Mapping domain
At `Domain` tap, a recovered account address can be mapped a domain which user want. However, Character length of the domain limit 20 and domains cannot reduplicate.

### Retrieve domain
At `Retrieve` tap, can retrieve mapped account address and domain.

### Send coin
At `Send` tap, user send coin to others by using domain.