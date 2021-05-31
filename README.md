# GCODB

GCODB is an [Ethereum-compatible](https://github.com/ethereum/go-ethereum) project. It uses a new consensus and new block reward for JINBAO ecosystem devices and IOT. And you can view the transactions on the [


Since the list of signers is 11, it is recommended that the confirmation number of general transfer transaction block be set to 11 (one round), and that of exchange block be set to 22 (two rounds).

## List of Chain ID's:
| Chain(s)    |  CHAIN_ID  |
| ----------  | :-----------:|
| mainnet     | 111            |
| testnet     | 3            |
| devnet      | 4            |

## Warning

We suggest that the GasPrice should not be less than 18Gwei, otherwise the transaction may not be packaged into the block.

## Build the source

Building GCODB requires both a Go (version 1.13.0 or later) and a C compiler. You can install them using your favourite package manager.

### MacOS & Linux

```
$ make gcodb
```

## Run node

> By default will run on mainnet , add `--testnet` options to join the testnet

    $ ./build/bin/gcodb console

## Create new account
    Users can create new account:

    > personal.newAccount()

## Security-related

### Encrypt your nodekey

     $ ./build/bin/gcodb security --passwd

### Decrypt your nodekey

     $ ./build/bin/gcodb security --unlock

# blockchain
