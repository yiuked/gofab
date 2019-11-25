package gofab

import (
	"encoding/hex"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"log"
)

func GetBlocksByTxId(client ledger.Client, hash string) (Block, error) {
	blockResource, err := client.QueryBlockByTxID(fab.TransactionID(hash))
	if err != nil {
		return Block{}, fmt.Errorf("QueryBlockByTxID err :%s", err.Error())
	}
	listTx := make([]TransactionDetail, 0)
	for _, data := range blockResource.Data.Data {
		txDetail, err := DecodeTransactionFromBlock(data, true)
		if err != nil {
			return Block{}, fmt.Errorf("QueryBlockByTxID err :%s", err)
		}
		listTx = append(listTx, *txDetail)
	}
	block := Block{}
	block.DataHash = hex.EncodeToString(blockResource.Header.GetDataHash())
	block.PreviousHash = hex.EncodeToString(blockResource.Header.PreviousHash)
	block.Number = blockResource.Header.Number
	// add create time
	if len(listTx) > 0 {
		block.CreateTime = listTx[0].CreateTime
	}
	block.TxList = listTx
	return block, nil
}

func GetInfo(client ledger.Client, options ...ledger.RequestOption) (*fab.BlockchainInfoResponse, error) {
	info, err := client.QueryInfo(options...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return info, nil
}

func GetBlock(blockNumber uint64, client ledger.Client, options ...ledger.RequestOption) (Block, error) {
	blockResource, err := client.QueryBlock(blockNumber, options...)
	if err != nil {
		return Block{}, fmt.Errorf("QueryBlockByTxID err :%s", err.Error())
	}
	listTx := make([]TransactionDetail, 0)
	for _, data := range blockResource.Data.Data {
		txDetail, err := DecodeTransactionFromBlock(data, true)
		if err != nil {
			return Block{}, fmt.Errorf("QueryBlockByTxID err :%s", err)
		}
		listTx = append(listTx, *txDetail)
	}
	block := Block{}
	block.DataHash = hex.EncodeToString(blockResource.Header.GetDataHash())
	block.PreviousHash = hex.EncodeToString(blockResource.Header.PreviousHash)
	block.Number = blockResource.Header.Number
	// add create time
	if len(listTx) > 0 {
		block.CreateTime = listTx[0].CreateTime
	}
	block.TxList = listTx
	return block, nil
}

func GetBlockByHash(blockHash []byte, client ledger.Client, options ...ledger.RequestOption) (Block, error) {
	blockResource, err := client.QueryBlockByHash(blockHash, options...)
	if err != nil {
		return Block{}, fmt.Errorf("QueryBlockByTxID err :%s", err.Error())
	}
	listTx := make([]TransactionDetail, 0)
	for _, data := range blockResource.Data.Data {
		txDetail, err := DecodeTransactionFromBlock(data, true)
		if err != nil {
			return Block{}, fmt.Errorf("QueryBlockByTxID err :%s", err)
		}
		listTx = append(listTx, *txDetail)
	}
	block := Block{}
	block.DataHash = hex.EncodeToString(blockResource.Header.GetDataHash())
	block.PreviousHash = hex.EncodeToString(blockResource.Header.PreviousHash)
	block.Number = blockResource.Header.Number
	// add create time
	if len(listTx) > 0 {
		block.CreateTime = listTx[0].CreateTime
	}
	block.TxList = listTx
	return block, nil
}

func GetConfig(client ledger.Client, options ...ledger.RequestOption) (fab.ChannelCfg, error) {
	cfg, err := client.QueryConfig(options...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cfg, nil
}
