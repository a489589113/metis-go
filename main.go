package main

import (
	"encoding/json"
	"log"

	// "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	// "github.com/btcsuite/btcutil"
	"fmt"
	"metis-go/logger"
)

func main() {
	// create new client instance
	connCfg := &rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "10.6.9.226:18332",
		User:         "bitnode",
		Pass:         "123456",
	}
	// txHash tx = ;
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}
	// var hash chainhash.Hash
	// chainhash.Decode(&hash, "000000000000004c592914735324e42c8ec3cb33f667fc01df746041a2d3fa53")
	// blockCount, err := client.GetBlock(&hash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Block info: %d", blockCount)

	var hash2 chainhash.Hash
	chainhash.Decode(&hash2, "4362785b9a52cea3591341384f28df422f22e35ceb84108daf1d47f3de465b66")
	st, err := client.GetRawTransactionVerbose(&hash2)
	if err != nil {
		log.Fatal(err)
	}
	txRawRetJ, err := json.Marshal(st)
	if err != nil {
		logger.Error(fmt.Sprintf("GetRawTx, fail to Marshal(txRawRet),err:%v", err))
	}
	logger.Info(fmt.Sprintf("GetRawTx, Marshal(txRawRet) ok, txRawRetJ:%v\n", string(txRawRetJ)))
	// transaction, err := client.DecodeRawTransaction(st)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("transaction:", *st)
	// list accounts
	// transaction, err := client.GetRawTransaction("4362785b9a52cea3591341384f28df422f22e35ceb84108daf1d47f3de465b66")
	// if err != nil {
	// 	log.Fatalf("error listing accounts: %v", err)
	// }
	// // iterate over accounts (map[string]btcutil.Amount) and write to stdout
	// for label, amount := range accounts {
	// 	log.Printf("%s: %s", label, amount)
	// }

	// // prepare a sendMany transaction
	// receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
	// if err != nil {
	// 	log.Fatalf("address receiver1 seems to be invalid: %v", err)
	// }
	// receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
	// if err != nil {
	// 	log.Fatalf("address receiver2 seems to be invalid: %v", err)
	// }
	// receivers := map[btcutil.Address]btcutil.Amount{
	// 	receiver1: 42,  // 42 satoshi
	// 	receiver2: 100, // 100 satoshi
	// }

	// // create and send the sendMany tx
	// txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
	// if err != nil {
	// 	log.Fatalf("error sendMany: %v", err)
	// }
	// log.Printf("sendMany completed! tx sha is: %s", transaction.String())
}
