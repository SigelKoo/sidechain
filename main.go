package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	mspID        = "Org2MSP"
	cryptoPath   = "../../fabric-samples/test-network/organizations/peerOrganizations/org2.example.com"
	certPath     = cryptoPath + "/users/0xb1ABb29CC3CD7b6c8D028866c370f92A2D1c870c@org1.example.com/msp/signcerts/cert.pem"
	keyPath      = cryptoPath + "/users/0xb1ABb29CC3CD7b6c8D028866c370f92A2D1c870c@org1.example.com/msp/keystore/152b9af6ac6b95478496c0b7a23665a956d864b9303dc6e1770c8fee3d760d13_sk"
	tlsCertPath  = cryptoPath + "/peers/peer0.org2.example.com/tls/ca.crt"
	peerEndpoint = "localhost:9051"
)

func main() {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gateway.Close()

	fmt.Println("exampleSubmit:")
	exampleSubmit(gateway)
	fmt.Println()
}

func exampleSubmit(gateway *client.Gateway) {
	network := gateway.GetNetwork("mychannel")
	contract := network.GetContract("token_erc20")

	timestamp := time.Now().String()
	fmt.Printf("Submitting \"Transfer\" transaction with arguments: time, %s\n", timestamp)

	// Submit transaction, blocking until the transaction has been committed on the ledger
	submitResult, err := contract.SubmitTransaction("Transfer", "eDUwOTo6Q049MHg0MTZiMWU1MzI5QmQ5N0JCNzA0ODY2YkQ0ODk3NDdiMjY4NDhmQTQyLE9VPWNsaWVudCxPPUh5cGVybGVkZ2VyLFNUPU5vcnRoIENhcm9saW5hLEM9VVM6OkNOPWNhLm9yZzIuZXhhbXBsZS5jb20sTz1vcmcyLmV4YW1wbGUuY29tLEw9SHVyc2xleSxTVD1IYW1wc2hpcmUsQz1VSw==", "1")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("Submit result: %s\n", string(submitResult))
	fmt.Println("Evaluating \"ClientAccountBalance\"")

	evaluateResult, err := contract.EvaluateTransaction("ClientAccountBalance")
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}

	fmt.Printf("Query result = %s\n", string(evaluateResult))
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, "peer0.org2.example.com")

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	privateKeyPEM, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}

	return identity.CertificateFromPEM(certificatePEM)
}