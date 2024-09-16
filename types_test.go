package sdk_go

import (
	"encoding/json"
	"testing"

	"github.com/nsf/jsondiff"
)

func TestDetailsStructure(t *testing.T) {
	expectedJsonStr := `
		{
		  "crypto": {
			"memo": "123456"
		  },
		  "customer": {
			"first_name": "First Name",
			"last_name": "Last Name",
			"customer_id": "customer-id"
		  },
		  "billing_info": {
			"address": "Address",
			"country_code": "UA",
			"country": "Ukraine"
		  },
		  "bank": {
			"account": "Bank Account",
			"id": "Bank ID",
			"name": "Bank Name"
		  },
		  "card_data": {
			"exp_month": "08",
			"exp_year": "2030"
		  },
		  "web_data": {
			"ip": "Firefox",
			"user_agent": "127.0.0.1"
		  },
		  "ui": {
			"language": "en"
		  },
		  "escrow_payment": {
			"payment_wallet_id": "uuid"
		  }
		}
	`

	crypto := &CryptoInfo{
		Memo: strPtr("123456"),
	}

	customer := &CustomerInfo{
		FirstName:  strPtr("First Name"),
		LastName:   strPtr("Last Name"),
		CustomerID: strPtr("customer-id"),
	}

	billing := &BillingInfo{
		Address:     strPtr("Address"),
		CountryCode: strPtr("UA"),
		Country:     strPtr("Ukraine"),
	}

	bank := &BankInfo{
		Account: strPtr("Bank Account"),
		ID:      strPtr("Bank ID"),
		Name:    strPtr("Bank Name"),
	}

	cardData := &CardData{
		ExpMonth: strPtr("08"),
		ExpYear:  strPtr("2030"),
	}

	webData := &WebData{
		IP:        strPtr("Firefox"),
		UserAgent: strPtr("127.0.0.1"),
	}

	uii := &UISchema{
		Language: strPtr("en"),
	}

	escowPayment := &EscrowPayment{
		PaymentWalletID: strPtr("uuid"),
	}

	details := Details{
		Crypto:        crypto,
		Customer:      customer,
		Billing:       billing,
		Bank:          bank,
		CardData:      cardData,
		WebData:       webData,
		UI:            uii,
		EscrowPayment: escowPayment,
	}

	outJsonStr, err := json.MarshalIndent(details, "", "  ")
	if err != nil {
		t.Fatal("error marshaling package", err)
	}

	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare([]byte(expectedJsonStr), []byte(outJsonStr), &diffOpts)

	if res != jsondiff.FullMatch {
		t.Errorf("the expected result is not equal to what we have: %s", diff)
	}
}

func strPtr(s string) *string {
	return &s
}
