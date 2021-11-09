package sidechain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func registerUser(username, org, passwd string) {
	sdk, err := fabsdk.New(config.FromFile("/home/gopath/sidechain/config/crypto-config.yaml"))
	if err != nil {
		fmt.Println("failed to create new SDK")
		return
	}
	ctx := sdk.Context()
	mspClient, err := msp.New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}
	identity, err := mspClient.CreateIdentity(
		&msp.IdentityRequest{
			ID: username,
			Affiliation: org,
			Attributes: []msp.Attribute{{Name: "ethereum", Value: username}},
		})
	if err != nil {
		fmt.Printf("Create identity return error %s\n", err)
		return
	}
	fmt.Printf("identity '%s' created\n", identity.ID)
	err = mspClient.Enroll(username, msp.WithSecret(passwd))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")
	_, err = mspClient.Register(&msp.RegistrationRequest{Name: username, Secret: passwd})
	if err != nil {
		fmt.Printf("Register return error %s\n", err)
		return
	}
	fmt.Println("register user is completed")
}