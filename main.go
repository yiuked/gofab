package gofab

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

var (
	channelID = "zsjr"
	user      = "Admin"
	orgName   = "Org1"
)

func main() {
	sdk := SDK("config.yaml")
	context := sdk.ChannelContext(channelID, fabsdk.WithUser(user), fabsdk.WithOrg(orgName))

	client, err := ledger.New(context)
	ErrFinal(err)

	cfg, err := client.QueryConfig()
	ErrLog(err)

	log.Println(cfg)
}
