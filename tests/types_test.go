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
			"taxpayer_identification_number": "12345-6688",
			"birthdate": "2004-09-21",
			"document_type": "CNPJ"
		  },
		  "billing_info": {
			"address": "Address",
			"country_code": "UA",
			"country": "Ukraine",
			"state": "UA",
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
			"number": "4111111111111111",
			"exp_month": "08",
			"exp_year": "2030",
			"cvv": "123"
		  },
		  "web_data": {
			"ip": "Firefox",
			"user_agent": "127.0.0.1",
			"browser_color_depth": 30,
			"browser_language": "en-GB,en-US;q=0.9,en;q=0.8",
			"browser_screen_height": 1080,
			"browser_screen_width": 1920,
			"browser_timezone": "Europe/Kiev",
			"browser_timezone_offset": -120,
			"browser_java_enabled": "false",
			"browser_java_script_enabled": "true",
			"browser_accept_header": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,/;q=0.8"
		  },
		  "ui": {
			"language": "en"
		  },
		  "escrow_payment": {
			"payment_wallet_id": "uuid"
		  },
     	  "project": {
			"url": "https://project-url.com"
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
		Birthdate:                    strPtr("2004-09-21"),
		DocumentType:                 strPtr("CNPJ"),
	}

	billing := &sdk_go.BillingInfo{
		Address:        strPtr("Address"),
		CountryCode:    strPtr("UA"),
		Country:        strPtr("Ukraine"),
		State:          strPtr("UA"),
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
		Number:   strPtr("4111111111111111"),
		ExpMonth: strPtr("08"),
		ExpYear:  strPtr("2030"),
		CVV:      strPtr("123"),
	}

	webData := &sdk_go.WebData{
		IP:                       "Firefox",
		UserAgent:                "127.0.0.1",
		BrowserColorDepth:        30,
		BrowserLanguage:          "en-GB,en-US;q=0.9,en;q=0.8",
		BrowserScreenHeight:      1080,
		BrowserScreenWidth:       1920,
		BrowserTimezone:          "Europe/Kiev",
		BrowserTimezoneOffset:    -120,
		BrowserJavaEnabled:       "false",
		BrowserJavaScriptEnabled: "true",
		BrowserAcceptHeader:      "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,/;q=0.8",
	}

	uii := &sdk_go.UISchema{
		Language: strPtr("en"),
	}

	escowPayment := &sdk_go.EscrowPayment{
		PaymentWalletID: strPtr("uuid"),
	}

	project := &sdk_go.Project{
		URL: "https://project-url.com",
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
		Project:       project,
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
