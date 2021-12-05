package fabricSDK

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/mocks"
)

func Example() {
	ctx := mockChannelProvider("mychannel")

	ec, err := event.New(ctx, event.WithBlockEvents())
	if err != nil {
		fmt.Println(err)
	}

	if ec != nil {
		fmt.Println("event client created")
	} else {
		fmt.Println("event client is nil")
	}

	registration, notifier, err := ec.RegisterChaincodeEvent("examplecc", "event123")
	if err != nil {
		fmt.Println("failed to register chaincode event")
	}
	defer ec.Unregister(registration)

	select {
	case ccEvent := <-notifier:
		fmt.Printf("received chaincode event %v\n", ccEvent)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout while waiting for chaincode event")
	}

}

func mockChannelProvider(channelID string) context.ChannelProvider {
	channelProvider := func() (context.Channel, error) {
		return mocks.NewMockChannel(channelID)
	}
	return channelProvider
}
