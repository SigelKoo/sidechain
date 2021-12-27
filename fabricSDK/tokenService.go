package fabricSDK

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) Transfer(recipient string, amount string) (string, error) {
	eventID := "Transfer"
	reg, notifier := regitserEvent(t.EventClient, t.ChaincodeID, eventID)
	defer t.EventClient.Unregister(reg)
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "Transfer", Args: [][]byte{[]byte(recipient), []byte(amount)}}
	respone, err := t.ChannelClient.Execute(req)
	if err != nil {
		return "", err
	}
	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) BalanceOf(account string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "BalanceOf", Args: [][]byte{[]byte(account)}}
	respone, err := t.ChannelClient.Query(req)
	if err != nil {
		return "", err
	}
	return string(respone.Payload), nil
}

func (t *ServiceSetup) ClientAccountBalance() (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "ClientAccountBalance"}
	respone, err := t.ChannelClient.Query(req)
	if err != nil {
		return "", err
	}
	return string(respone.Payload), nil
}

func (t *ServiceSetup) ClientAccountID() (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "ClientAccountID"}
	respone, err := t.ChannelClient.Query(req)
	if err != nil {
		return "", err
	}
	return string(respone.Payload), nil
}

func (t *ServiceSetup) Mint(amount string) (string, error) {
	eventID := "eventMint"
	reg, notifier := regitserEvent(t.EventClient, t.ChaincodeID, eventID)
	defer t.EventClient.Unregister(reg)
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "Mint", Args: [][]byte{[]byte(amount)}}
	respone, err := t.ChannelClient.Execute(req)
	if err != nil {
		return "", err
	}
	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) Burn(amount string) (string, error) {
	eventID := "eventBurn"
	reg, notifier := regitserEvent(t.EventClient, t.ChaincodeID, eventID)
	defer t.EventClient.Unregister(reg)
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "Burn", Args: [][]byte{[]byte(amount)}}
	respone, err := t.ChannelClient.Execute(req)
	if err != nil {
		return "", err
	}
	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	return string(respone.TransactionID), nil
}
