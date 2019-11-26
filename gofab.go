package gofab

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func SDK(configPath string) *fabsdk.FabricSDK {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	ErrFinal(err)
	return sdk
}
