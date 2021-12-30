package fabricSDK

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"encoding/base64"

	"sidechain/ethSDK"

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
		thisEvent := &anaEvent{}
		json.Unmarshal(ccEvent.Payload, thisEvent)
		fmt.Println(thisEvent.From, thisEvent.To, thisEvent.Value)
		fromDec, err := base64.StdEncoding.DecodeString(thisEvent.From)
		if err != nil {
			return fmt.Errorf("ccEvent From解析失败：%s", err)
		}

		toDec, err := base64.StdEncoding.DecodeString(thisEvent.To)
		if err != nil {
			return fmt.Errorf("ccEvent To解析失败：%s", err)
		}
		fmt.Println(GetX509UserName(string(fromDec)))
		fmt.Println(GetX509UserName(string(toDec)))
		fmt.Println(thisEvent.Value)
		if GetX509UserName(string(toDec)) == "minter" {
			fmt.Println(ethSDK.Transfer("HTTP://127.0.0.1:8501", "0xD78d66C33933a05c57c503d61667918f95cee351", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", GetX509UserName(string(fromDec)), strconv.Itoa(thisEvent.Value)))
		}
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

func GetX509UserName(str string) string {
	s1 := strings.Index(str, ",")
	s2 := strings.Index(str, "=")
	return str[s2+1 : s1]
}
