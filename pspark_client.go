package sdk_go

import (
	"fmt"
)

func PSParkClient(secret string, apiKey string) *PSPark {
	return &PSPark{
		Secret:  secret,
		APIKey:  apiKey,
		BaseURL: BaseURL,
	}
}

func (client *PSPark) GetWalletBalance(walletId string) (WalletBalance, error) {

	url := fmt.Sprintf("wallet/%s/balance", walletId)

	return makeAuthenticatedRequest[WalletBalance](url, map[string]interface{}{}, client)
}

func (client *PSPark) GetBalances() ([]WalletBalance, error) {

	url := "balances"

	res, err := makeAuthenticatedRequest[map[string]WalletBalance](url, map[string]interface{}{}, client)

	if err != nil {
		return nil, err
	}

	var balances []WalletBalance

	for _, balance := range res {
		balances = append(balances, balance)
	}

	return balances, nil
}

func (client *PSPark) CreateInvoice(walletId string, data InvoiceRequest) (InvoiceResponse, error) {

	url := fmt.Sprintf("wallet/%s/invoice/create", walletId)

	return makeAuthenticatedRequest[InvoiceResponse](url, data, client)

}

func (client *PSPark) CreateAddress(walletId string, data AddressRequest) (AddressResponse, error) {

	url := fmt.Sprintf("wallet/%s/address/create", walletId)

	return makeAuthenticatedRequest[AddressResponse](url, data, client)

}

func (client *PSPark) CreateWithdrawal(walletId string, data WithdrawalRequest) (WithdrawalResponse, error) {

	url := fmt.Sprintf("wallet/%s/withdrawal/create", walletId)

	return makeAuthenticatedRequest[WithdrawalResponse](url, data, client)

}
