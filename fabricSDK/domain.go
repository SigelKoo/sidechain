package fabricSDK

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type ServiceSetup struct {
	ChaincodeID   string
	ChannelClient *channel.Client
	EventClient   *event.Client
}

type anaEvent struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

func regitserEvent(client *event.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {

	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
		thisEvent := &anaEvent{}
		json.Unmarshal(ccEvent.Payload, thisEvent)
		fmt.Println(string(ccEvent.Payload[:]))
		fmt.Printf("%v\n", thisEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
