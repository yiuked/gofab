package gofab

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	ptCommon "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	ptUtils "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/utils"
	ptPeer "github.com/hyperledger/fabric/protos/peer"
	"log"
	"time"
)

func GetTransactionByTxID(txID string, client *ledger.Client, options ...ledger.RequestOption) (*Transaction, error) {
	tx, err := client.QueryTransaction(fab.TransactionID(txID), options...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return DecodeTransaction(tx.TransactionEnvelope)
}

// 从SDK中Block.BlockDara.Data中提取交易具体信息
func DecodeTransactionFromBlock(data []byte, needArgs bool) (*Transaction, error) {
	txd := &Transaction{}
	env, err := ptUtils.GetEnvelopeFromBlock(data)
	if err != nil {
		return txd, err
	}
	return DecodeTransaction(env)
}

func DecodeTransaction(e *ptCommon.Envelope) (*Transaction, error) {
	args := make([]string, 0)

	payload, err := ptUtils.GetPayload(e)
	if err != nil {
		return nil, err
	}

	channelHeaderBytes := payload.Header.ChannelHeader
	channelHeader := &ptCommon.ChannelHeader{}

	if err := proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
		return nil, err
	}

	tx, err := ptUtils.GetTransaction(payload.Data)
	if err != nil {
		return nil, err
	}
	actionPayload, err := ptUtils.GetChaincodeActionPayload(tx.Actions[0].Payload)
	if err != nil {
		return nil, err
	}

	propPayload := &ptPeer.ChaincodeProposalPayload{}

	if err := proto.Unmarshal(actionPayload.ChaincodeProposalPayload, propPayload); err != nil {
		return nil, errors.New(fmt.Sprintf("Unmarshal ChaincodeProposalPayload Error:%s", err.Error()))
	}

	invokeSpec := &ptPeer.ChaincodeInvocationSpec{}
	err = proto.Unmarshal(propPayload.Input, invokeSpec)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unmarshal ChaincodeInvocationSpec Error:%s", err.Error()))
	}

	for _, v := range invokeSpec.ChaincodeSpec.Input.Args {
		args = append(args, string(v))
	}

	result := &Transaction{
		TransactionId: channelHeader.TxId,
		Args:          args,
		CreateTime:    time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05"),
	}
	return result, nil
}
