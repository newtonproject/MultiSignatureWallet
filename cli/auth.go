package cli

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewKeyedTransactorByAccount(wallet *keystore.KeyStore, account accounts.Account, passphrase string, networkID *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// force update gas to 1.5 * gas
			if tx.To() != nil {
				tx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), tx.Gas()*15/10, tx.GasPrice(), tx.Data())
			}
			fmt.Println("The tx is as follow: ")
			fmt.Println("\tFrom:", account.Address.String())
			if tx.To() == nil {
				fmt.Println("\tTo: ContractCreate")
			} else {
				fmt.Println("\tTo:", tx.To().String())
			}
			fmt.Println("\tValue:", getWeiAmountTextByUnit(tx.Value(), UnitETH))
			fmt.Println("\tData:", hex.EncodeToString(tx.Data()))
			fmt.Println("\tNonce:", tx.Nonce())
			fmt.Println("\tGasPrice:", getWeiAmountTextByUnit(tx.GasPrice(), UnitETH))
			fmt.Println("\tGasLimit:", tx.Gas())
			fmt.Println("\tGasFee:", getWeiAmountTextByUnit(big.NewInt(0).Mul(tx.GasPrice(), big.NewInt(0).SetUint64(tx.Gas())), UnitETH))

			for trials := 0; trials <= 1; trials++ {
				err := wallet.Unlock(account, passphrase)
				if err == nil {
					break
				}
				if trials >= 1 {
					return nil, fmt.Errorf("failed to unlock account %s (%v)", account.Address.String(), err)

				}
				prompt := fmt.Sprintf("Unlocking account %s", account.Address.String())
				passphrase, _ = getPassPhrase(prompt, false)
			}

			return wallet.SignTx(account, tx, networkID)
		},
	}
}
