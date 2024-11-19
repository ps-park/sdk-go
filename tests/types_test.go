package tests

import (
	"encoding/json"
	"testing"

	"github.com/nsf/jsondiff"
	sdk_go "github.com/ps-park/sdk-go"
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
			"customer_id": "customer-id",
			"taxpayer_identification_number": "12345-6688"
		  },
		  "billing_info": {
			"address": "Address",
			"country_code": "UA",
			"country": "Ukraine",
			"city": "Kyiv",
			"post_code": "01001",
			"region": "Kyiv Oblast",
			"payment_purpose": "Bank transfer",
			"street": "Baker Street"
		  },
		  "bank": {
			"account": "Bank Account",
			"id": "Bank ID",
			"name": "Bank Name",
			"bic_code": "SABADE5S"
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

	crypto := &sdk_go.CryptoInfo{
		Memo: strPtr("123456"),
	}

	customer := &sdk_go.CustomerInfo{
		FirstName:                    strPtr("First Name"),
		LastName:                     strPtr("Last Name"),
		CustomerID:                   strPtr("customer-id"),
		TaxpayerIdentificationNumber: strPtr("12345-6688"),
	}

	billing := &sdk_go.BillingInfo{
		Address:        strPtr("Address"),
		CountryCode:    strPtr("UA"),
		Country:        strPtr("Ukraine"),
		City:           strPtr("Kyiv"),
		PostCode:       strPtr("01001"),
		Region:         strPtr("Kyiv Oblast"),
		PaymentPurpose: strPtr("Bank transfer"),
		Street:         strPtr("Baker Street"),
	}

	bank := &sdk_go.BankInfo{
		Account: strPtr("Bank Account"),
		ID:      strPtr("Bank ID"),
		Name:    strPtr("Bank Name"),
		BicCode: strPtr("SABADE5S"),
	}

	cardData := &sdk_go.CardData{
		ExpMonth: strPtr("08"),
		ExpYear:  strPtr("2030"),
	}

	webData := &sdk_go.WebData{
		IP:        strPtr("Firefox"),
		UserAgent: strPtr("127.0.0.1"),
	}

	uii := &sdk_go.UISchema{
		Language: strPtr("en"),
	}

	escowPayment := &sdk_go.EscrowPayment{
		PaymentWalletID: strPtr("uuid"),
	}

	details := sdk_go.Details{
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
