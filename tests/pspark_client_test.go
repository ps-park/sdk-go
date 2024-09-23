package tests

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	sdk_go "github.com/ps-park/sdk-go"
)

var secret = generatePrivateKey()

func TestGetWalletBalance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{ "code": 0, "message": "ok", "data": {"balance":"100", "currency":"USD", "name":"Test Wallet", "wallet_id":"test-id"}}`))
	}))
	defer server.Close()

	client := sdk_go.PSPark{
		Secret:  secret,
		APIKey:  "test-api",
		BaseURL: server.URL,
	}

	balance, err := client.GetWalletBalance("test-id")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the returned balance
	if balance.Balance != "100" {
		t.Errorf("Expected balance 100, got %v", balance.Balance)
	}
}

func TestGetBalances(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{ "code": 0, "message": "ok", "data": { "wallet1": {"balance":"100", "currency":"USD", "name":"Test Wallet", "wallet_id":"test-id"}}}`))
	}))
	defer server.Close()

	client := sdk_go.PSPark{
		Secret:  secret,
		APIKey:  "test-api",
		BaseURL: server.URL,
	}

	balances, err := client.GetBalances()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the returned balance
	if balances[0].Balance != "100" {
		t.Errorf("Expected balance 100, got %v", balances[0].Balance)
	}
}

func TestCreateInvoice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{ "code": 0, "message": "ok", "data": {"id":"test-id"}}`))
	}))
	defer server.Close()

	client := sdk_go.PSPark{
		Secret:  secret,
		APIKey:  "test-api",
		BaseURL: server.URL,
	}

	invoice, err := client.CreateInvoice("test-id", sdk_go.InvoiceRequest{
		Reference: "test-ref",
		Amount:    100,
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the returned invoice id
	if invoice.ID != "test-id" {
		t.Errorf("Expected invoice id test-id, got %v", invoice.ID)
	}
}

func TestCreateAddress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{ "code": 0, "message": "ok", "data": {"id":"test-address"}}`))
	}))
	defer server.Close()

	client := sdk_go.PSPark{
		Secret:  secret,
		APIKey:  "test-api",
		BaseURL: server.URL,
	}

	address, err := client.CreateAddress("test-id", sdk_go.AddressRequest{
		Reference: "test-ref",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the returned address
	if address.ID != "test-address" {
		t.Errorf("Expected address test-address, got %v", address.Address)
	}
}

func TestCreateWithdrawal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{ "code": 0, "message": "ok", "data": {"id":"test-withdrawal"}}`))
	}))
	defer server.Close()

	client := sdk_go.PSPark{
		Secret:  secret,
		APIKey:  "test-api",
		BaseURL: server.URL,
	}

	withdrawal, err := client.CreateWithdrawal("test-id", sdk_go.WithdrawalRequest{
		Reference: "test-ref",
		Amount:    100,
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the returned withdrawal id
	if withdrawal.ID != "test-withdrawal" {
		t.Errorf("Expected withdrawal id test-withdrawal, got %v", withdrawal.ID)
	}

}

func generatePrivateKey() string {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate private key: %s", err)
	}

	// Encode the private key into PEM format.
	privateKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Convert the PEM block into a string.
	privateKeyPEMBytes := pem.EncodeToMemory(privateKeyPEM)
	privateKeyPEMString := string(privateKeyPEMBytes)

	return privateKeyPEMString
}
