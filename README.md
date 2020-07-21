
## MultiSignatureWallet 

`MultiSignatureWallet` project contains the following:
* Deploy Multisignature wallet contract
* Submit, confirm, execute, revoke or check transactionId
* Get basic information
* Manage owners, dailyLimit and required

# Table of Contents

- [QuickStart](#quickstart)
   * [Download from releases](#download-from-releases)
   * [Building the source](#building-the-source)
   * [Windows](#windows)
   * [Linux or Mac](#linux-or-mac)
- [Usage](#installing)
   * [Contract](#contract)
   * [Commandline](#commandline)
      * [Help](#help)
      * [Use config.toml](#use-configtoml)
      * [Initialize config file](#initialize-config-file)
      * [Create account](#create-account)
      * [Deploy contract](#deploy-contract)
      * [Submit transaction](#submit-transaction)
      * [Confirm transactionID](#confirm-transactionid)
      * [Revoke transactionID](#revoke-transactionid)
      * [Execute transactionID](#execute-transactionid)
      * [List transactionIDs](#list-transactionids)
      * [Get basic info](#get-basic-info)
      * [Manage owners](#manage-owners)
      * [Update daily limit or the number of required](#update-daily-limit-or-the-number-of-required)
      * [Build transaction online](#build-transaction-online)
      * [Sign transaction offline](#sign-transaction-offline)
      * [Broadcast signed transaction online](#broadcast-signed-transaction-online)
- [Examples](#examples)
   * [Online Example (3/3)](#online-example-33)
      * [For owner A](#for-owner-a)
         * Submit transaction
      * [For owner B and C](#for-owner-b-and-c)
         * Confirm transactionID
   * [Offline Example (3/3)](#offline-example-33)
      * [For owner A](#for-owner-a-1)
         1. Build Submit Transaction Information (Online Computer)
         2. Sign Transaction (Offline Computer)
         3. Broadcast Signed Transaction (Online Computer)
      * [For owner B and C](#for-owner-b-and-c-1)
         1. Build Confirm TransactionID Information (Online Computer)
         2. Sign Transaction (Offline Computer)
         3. Broadcast Signed Transaction (Online Computer)
      * [Offline Example Tips](#offline-example-tips)


## QuickStart

### Download from releases

Binary archives are published at https://release.cloud.diynova.com/newton/MultiSignatureWallet/.

### Building the source

#### Windows

install:

```bash
git clone https://github.com/newtonproject/MultiSignatureWallet.git && cd MultiSignatureWallet && make install
```

run MultiSignatureWallet:

```bash
%GOPATH%/bin/MultiSignatureWallet.exe
```

#### Linux or Mac

install:

```bash
git clone https://github.com/newtonproject/MultiSignatureWallet.git && cd MultiSignatureWallet && make install
```

run MultiSignatureWallet:

```bash
$GOPATH/bin/MultiSignatureWallet
```

## Usage

### Contract

Use commands `go generate` or `abigen` to generate MultiSigWalletWithDailyLimit.go from [MultiSigWalletWithDailyLimit.sol](https://github.com/gnosis/MultiSigWallet/blob/v1.6.0/contracts/MultiSigWalletWithDailyLimit.sol).

```bash
abigen --sol contract/MultiSigWalletWithDailyLimit.sol --pkg cli --out cli/MultiSigWalletWithDailyLimit.go
```

### Commandline

#### Help

Use command `MultiSignatureWallet help` to display the usage.

```bash
Usage:
  MultiSigWallet [flags]
  MultiSigWallet [command]

Available Commands:
  account     Manage NewChain accounts
  broadcast   Broadcast sign transacion hex in the signTxFilePath to blockchain
  build       Build transaction
  confirm     Confirm transactionId
  deploy      Deploy NewChain contract
  execute     Execute transactionId
  help        Help about any command
  info        Show the basic info of contract wallet or a transaction ID
  init        Initialize config file
  list        List of transaction IDs in defined range
  owner       Manage contract owners
  revoke      Revoke transactionId
  sign        Sign the transaction in the file
  submit      Submit a transaction, pay amount in unit to target address
  update      Manage daily limit and requirement
  version     Get version of MultiSigWallet CLI

Flags:
  -c, --config path               The path to config file (default "./config.toml")
  -a, --contractAddress address   Contract address
  -f, --from address              the from address who pay gas
  -h, --help                      help for MultiSigWallet
  -i, --rpcURL url                NewChain json rpc or ipc url (default "https://rpc1.newchain.newtonproject.org")
  -w, --walletPath directory      Wallet storage directory (default "./wallet/")

Use "MultiSigWallet [command] --help" for more information about a command.
```

#### Use config.toml

You can use a configuration file to simplify the command line parameters.

One available configuration file `config.toml` is as follows:


```conf
contractaddress = "0xf09E6759c2588eE8435902d16350E321CBD27af3"
from = "0xdDeB86Dd09F16316B67322199E288d7AF35E0806"
rpcurl = "https://rpc1.newchain.newtonproject.org"
walletpath = "./wallet/"
```

#### Initialize config file

```bash
# Initialize config file
$ MultiSignatureWallet init
```

Just press Enter to use the default configuration, and it's best to create a new user.


```bash
$ MultiSignatureWallet init
Initialize config file
Enter file in which to save (./config.toml):
Enter the wallet storage directory (./wallet/):
Enter NewChain json rpc or ipc url (https://rpc1.newchain.newtonproject.org):
Create a default account or not: [Y/n]
Your new account is locked with a password. Please give a password. Do not forget this password.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
New accout is  0xdDeB86Dd09F16316B67322199E288d7AF35E0806
Get faucet for 0xdDeB86Dd09F16316B67322199E288d7AF35E0806
Your configuration has been saved in  ./config.toml
```

#### Create account

```bash
# Create an account with faucet
MultiSignatureWallet account new --faucet

# Create 10 accounts
MultiSignatureWallet account new -n 10 --faucet

# Create an account with the standard scrypt for keystore
MultiSignatureWallet account new -s
```

#### Deploy contract

```bash
# Deploy multisignature wallet contract
# with 3 owner, 2 required confirmations
MultiSignatureWallet deploy -o 0xdDeB86Dd09F16316B67322199E288d7AF35E0806,0x9B3deA9C636BA262f870f98a1c64d444BF0f6544,0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31 -r 2

# with 3 owner, 2 required confirmations and 1024 NEW daily withdraw without confirmations
MultiSignatureWallet deploy -o 0xdDeB86Dd09F16316B67322199E288d7AF35E0806,0x9B3deA9C636BA262f870f98a1c64d444BF0f6544,0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31 -r 2 -l 1024

# with 3 owner, 2 required confirmations and 10 WEI daily withdraw without confirmations
MultiSignatureWallet deploy -o 0xdDeB86Dd09F16316B67322199E288d7AF35E0806,0x9B3deA9C636BA262f870f98a1c64d444BF0f6544,0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31 -r 2 -l 1024 -u WEI
```

#### Submit transaction

```bash
# Submit transaction, pay 10 NEW from the contract address to 0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31
MultiSignatureWallet submit 10 -t 0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31
```

This will get a message back from the wallet contract with a number or hash code which is the transactionID.

You have to communicate this transactionID to other owners to get them to confirm it.

#### Confirm transactionID
```bash
# Confirm transaction ID
MultiSignatureWallet confirm 1
```

#### Revoke transactionID
```bash
# Revoke a confirmation for a transaction
MultiSignatureWallet revoke 1
```

#### Execute transactionID
```bash
# Execute a confirmed transaction
MultiSignatureWallet execute 1
```

#### List transactionIDs
```
# List transaction IDs
MultiSignatureWallet list

# List pending transaction IDs
MultiSignatureWallet list --pending

# List executed transaction IDs
MultiSignatureWallet list --executed

# List transaction IDs from index 10 to 20
MultiSignatureWallet list --fromindex 10 --toindex 20
```

#### Get basic info

```bash
# Get the basic information of the contract wallet
MultiSignatureWallet info

# Get the basic infofmation of a transaction ID
MultiSignatureWallet info 1

# Force display value in unit
MultiSignatureWallet info -u NEW
MultiSignatureWallet info 1 -u NEW

# Get the basic information of the specified contract address
MultiSignatureWallet info -a 0x8CFA0D92673bECC7A4B480844376A82b942E469b
```

#### Manage owners

```bash
# List all owners
MultiSignatureWallet owner list

# Check whether an address is an owner
MultiSignatureWallet owner check 0x9B3deA9C636BA262f870f98a1c64d444BF0f6544

# Add an owner
MultiSignatureWallet owner add 0x9B3deA9C636BA262f870f98a1c64d444BF0f6544

# Remove an owner
MultiSignatureWallet owner remove 0x9B3deA9C636BA262f870f98a1c64d444BF0f6544

# Replace an owner(0x9B3deA9C636BA262f870f98a1c64d444BF0f6544) with a new one
MultiSignatureWallet owner replace 0x9B3deA9C636BA262f870f98a1c64d444BF0f6544 0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31
```

#### Update daily limit or the number of required

```bash
# Update the daily limit to 1024 NEW
MultiSignatureWallet update dailylimit 1024

# Update the daily limit to 10 WEI
MultiSignatureWallet update dailylimit 10 -u WEI

# Update the number of required to 2 
MultiSignatureWallet update required 2
```

#### Build transaction online
```bash
# Build transaction
MultiSignatureWallet build
```

#### Sign transaction offline
```bash
# Sign transaction from file and save sign transaction hex to file
MultiSignatureWallet sign tx.txt

# Sign transaction save to the specified file
MultiSignatureWallet sign tx.txt --out tx.sign
```

#### Broadcast signed transaction online
```bash
# Broadcast signed transacion hex to NewChain system
MultiSignatureWallet broadcast tx.sign
```

## Examples

### Online Example (3/3)

##### For owner A

##### Submit transaction
```bash
# Submit transaction, pay 10 NEW from the contract address to 0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31
MultiSignatureWallet submit 10 -t 0xc8B5c4cB6DB7254d082b24A96627F143E8A80c31
```

This will get a message back from the wallet contract with a number or hash code which is the transactionID.

You have to communicate this transactionID to other owners to get them to confirm it.

#### For owner B and C

##### Confirm transaction
```bash
# Confirm transaction ID
MultiSignatureWallet confirm 1
```

### Offline Example (3/3)

`MultiSignatureWallet` provides offline signature, the process is as follows:

#### For owner A

##### Build Submit Transaction Information (Online Computer)
```bash
# Build transaction, use the defalult setting
$ MultiSignatureWallet build
Enter contract address (default: 0xc264cdD35F99C72b78315D25E6E353dBc7AA0430):
Enter from address who sign tx (default: 0x7c1d845a0CC7E24352A59FEF437eB27b504769DE):
Which action to use? (default = submit)
 1. Submit - Submit transaction, pay to address
 2. Confirm - Confirm transaction ID
 3. Revoke - Revoke a confirmation for a transaction
 4. Execute - Execute a confirmed transaction
 5. OwnerAdd - Add an owner
 6. OwnerRemove - Remove an owner
 7. OwnerReplace - Replace an owner
 8. DailyLimit - Update the daily limit
 9. Required - Update the number of required
Enter the number of action (default: 1):
Enter to address: 0x7c1d845a0CC7E24352A59FEF437eB27b504769DE
Enter unit for amount (NEW or WEI, default NEW):
Enter amount to pay in NEW (default: 0):
Enter text message (default is empty):
Transaction details are as follows:
{
 "from": "0x7c1d845a0cc7e24352a59fef437eb27b504769de",
 "to": "0xc264cdd35f99c72b78315d25e6e353dbc7aa0430",
 "value": "0",
 "unit": "WEI",
 "data": "0xc64274740000000000000000000000007c1d845a0cc7e24352a59fef437eb27b504769de000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000",
 "nonce": 52,
 "gasPrice": 100,
 "gas": 132155,
 "networkID": 16888
}
Enter file to save transaction (default: 20180929180404.tx):
Successfully save transaction to file 20180929180404.tx
```

##### Sign Transaction (Offline Computer)
```bash
# Sign transaction from file tx.txt and save sign transaction hex to file tx.txt.sign
MultiSignatureWallet sign tx.txt
```

##### Broadcast Signed Transaction (Online Computer)
```bash
# Broadcast signed transacion hex to NewChain system
# This will return an transacionID and tell it to the other owners.
MultiSignatureWallet broadcast tx.txt.sign
```

#### For owner B and C

##### Build Confirm TransactionID Information (Online Computer)
```bash
# Build transaction and select 2
MultiSignatureWallet build
```

##### Sign Transaction (Offline Computer)
```bash
# Sign transaction from file tx.txt and save sign transaction hex to file tx.txt.sign
MultiSignatureWallet sign tx.txt
```

##### Broadcast Signed Transaction (Online Computer)
```bash
# Broadcast signed transacion hex to NewChain system
MultiSignatureWallet broadcast tx.txt.sign
```

#### Offline Example Tips:
* The last owner should `build` transaction after the other owners `broadcast`, otherwise it will causes the error `intrinsic gas too low`.