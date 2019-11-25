package gofab

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

func GetChannelClient(sdk *fabsdk.FabricSDK, channelID string, user string, orgName string) *channel.Client {
	//prepare channel client context using client context
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(user), fabsdk.WithOrg(orgName))

	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)

	if err != nil {
		log.Panicf("Failed to create new channel client: %s", err)
	}
	return client
}
