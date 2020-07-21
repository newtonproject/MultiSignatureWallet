package cli

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"
)

func (cli *CLI) buildBroadcastCmd() *cobra.Command {
	signMesgCmd := &cobra.Command{
		Use:                   "broadcast <signTxFilePath>",
		Short:                 "Broadcast sign transacion hex in the signTxFilePath to blockchain",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			infileStr := args[0]

			signTxStr, err := readLineFromFile(infileStr)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(signTxStr))

			signTxByte := common.FromHex(string(signTxStr))
			signTx := new(types.Transaction)
			if err := rlp.DecodeBytes(signTxByte, signTx); err != nil {
				fmt.Println("DecodeBytes signTxHex error: ", err)
				return
			}

			ctx := context.Background()
			client, err := rpc.DialContext(ctx, cli.rpcURL)
			if err != nil {
				fmt.Println("DialContext: ", err)
				return
			}
			if err := client.CallContext(ctx, nil, "eth_sendRawTransaction", signTxStr); err != nil {
				fmt.Println("CallContext Error: ", err)
				return
			}
			fmt.Println("Waiting for transaction receipt...")
			txp, err := waitMined(ctx, client, signTx.Hash())
			if err != nil {
				fmt.Printf("Error: wait tx mined error(%v)\n", err)
				return
			}
			showTransactionReceipt(cli.rpcURL, signTx.Hash().String())
			cli.showSubmitID(txp)
		},
	}
	return signMesgCmd
}

func readLineFromFile(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			return text, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func waitMined(ctx context.Context, client *rpc.Client, hash common.Hash) (*types.Receipt, error) {
	transactionReceipt := func() (*types.Receipt, error) {
		var r *types.Receipt
		err := client.CallContext(ctx, &r, "eth_getTransactionReceipt", hash)
		if err == nil {
			if r == nil {
				return nil, ethereum.NotFound
			}
		}
		return r, err
	}

	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := transactionReceipt()
		if receipt != nil {
			return receipt, nil
		}
		if err != nil {
			// logger.Trace("Receipt retrieval failed", "err", err)
		} else {
			// logger.Trace("Transaction not yet mined")
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (cli *CLI) showSubmitID(txp *types.Receipt) bool {
	if txp != nil && len(txp.Logs) > 0 {
		log := txp.Logs[0]
		if log != nil {
			topics := log.Topics
			if topics != nil && len(topics) > 1 {
				// event Submission(uint indexed transactionId)
				// sha3('Submission(uint256)'): 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51
				if topics[0] == common.HexToHash("0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51") {
					data := topics[1]
					transferID := new(big.Int).SetBytes(data.Bytes())
					fmt.Println("TransferID: ", transferID)
					return true
				}
			}
		}
	}
	return false
}
