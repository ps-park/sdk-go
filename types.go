package sdk_go

type PSPark struct {
	Secret  string
	APIKey  string
	BaseURL string
}

type ResponseDTO[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type WalletBalance struct {
	WalletId string `json:"wallet_id"`
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
	Name     string `json:"name"`
}

type CryptoInfo struct {
	Memo *string `json:"memo,omitempty"`
}

type CustomerInfo struct {
	FirstName                    *string `json:"first_name,omitempty"`
	LastName                     *string `json:"last_name,omitempty"`
	Email                        *string `json:"email,omitempty"`
	Phone                        *string `json:"phone,omitempty"`
	CustomerID                   *string `json:"customer_id,omitempty"`
	NationalID                   *string `json:"national_id,omitempty"`
	TaxpayerIdentificationNumber *string `json:"taxpayer_identification_number,omitempty"`
	Birthdate                    *string `json:"birthdate,omitempty"`
}

type BillingInfo struct {
	Address        *string `json:"address,omitempty"`
	CountryCode    *string `json:"country_code,omitempty"`
	Country        *string `json:"country,omitempty"`
	State          *string `json:"state,omitempty"`
	City           *string `json:"city,omitempty"`
	PostCode       *string `json:"post_code,omitempty"`
	Region         *string `json:"region,omitempty"`
	PaymentPurpose *string `json:"payment_purpose,omitempty"`
	Street         *string `json:"street,omitempty"`
}

type BankInfo struct {
	Account *string `json:"account,omitempty"`
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	BicCode *string `json:"bic_code,omitempty"`
}

type CardData struct {
	Number   *string `json:"number,omitempty"`
	ExpMonth *string `json:"exp_month,omitempty"`
	ExpYear  *string `json:"exp_year,omitempty"`
	CVV      *string `json:"cvv,omitempty"`
}

type WebData struct {
	IP                       string `json:"ip,omitempty"`
	UserAgent                string `json:"user_agent,omitempty"`
	BrowserColorDepth        int    `json:"browser_color_depth,omitempty"`
	BrowserLanguage          string `json:"browser_language,omitempty"`
	BrowserScreenHeight      int    `json:"browser_screen_height,omitempty"`
	BrowserScreenWidth       int    `json:"browser_screen_width,omitempty"`
	BrowserTimezone          string `json:"browser_timezone,omitempty"`
	BrowserTimezoneOffset    int    `json:"browser_timezone_offset,omitempty"`
	BrowserJavaEnabled       string `json:"browser_java_enabled,omitempty"`
	BrowserJavaScriptEnabled string `json:"browser_java_script_enabled,omitempty"`
	BrowserAcceptHeader      string `json:"browser_accept_header,omitempty"`
}

type UISchema struct {
	Language *string `json:"language,omitempty"` // Values: "en", "ua", "ru"
}

type EscrowPayment struct {
	PaymentWalletID *string `json:"payment_wallet_id,omitempty"`
}

type Project struct {
	URL string `json:"url,omitempty"`
}

type FlowData struct {
	Action string        `json:"action"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type InvoiceRequest struct {
	Reference   string   `json:"reference"`
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	LimitMinute *int     `json:"limit_minute,omitempty"`
	CallbackURL *string  `json:"callback_url,omitempty"`
	Amount      float64  `json:"amount"`
	Currency    string   `json:"currency"`
	ReturnURL   string   `json:"return_url"`
	Details     *Details `json:"details,omitempty"`
}

type InvoiceResponse struct {
	ID            string    `json:"id"`
	Reference     string    `json:"reference"`
	WalletID      string    `json:"wallet_id"`
	Currency      string    `json:"currency"`
	Amount        float64   `json:"amount"`
	AmountInitial float64   `json:"amount_initial,omitempty"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	StatusCode    int       `json:"status_code,omitempty"`
	StatusMessage string    `json:"status_message,omitempty"`
	PaymentFee    float64   `json:"payment_fee,omitempty"`
	Address       string    `json:"address,omitempty"`
	Memo          string    `json:"memo,omitempty"`
	FlowData      *FlowData `json:"flow_data,omitempty"`
}

type AddressRequest struct {
	Reference   string  `json:"reference"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	LimitMinute *int    `json:"limit_minute,omitempty"`
	CallbackURL *string `json:"callback_url,omitempty"`
}

type AddressResponse struct {
	ID            string  `json:"id"`
	Reference     string  `json:"reference"`
	WalletID      string  `json:"wallet_id"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	AmountInitial float64 `json:"amount_initial,omitempty"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	StatusCode    int     `json:"status_code,omitempty"`
	StatusMessage string  `json:"status_message,omitempty"`
	PaymentFee    float64 `json:"payment_fee,omitempty"`
	Address       string  `json:"address,omitempty"`
	Memo          string  `json:"memo,omitempty"`
}

type Details struct {
	Crypto        *CryptoInfo    `json:"crypto,omitempty"`
	Customer      *CustomerInfo  `json:"customer,omitempty"`
	Billing       *BillingInfo   `json:"billing_info,omitempty"`
	Bank          *BankInfo      `json:"bank,omitempty"`
	CardData      *CardData      `json:"card_data,omitempty"`
	WebData       *WebData       `json:"web_data,omitempty"`
	UI            *UISchema      `json:"ui,omitempty"`
	EscrowPayment *EscrowPayment `json:"escrow_payment,omitempty"`
	Project       *Project       `json:"project,omitempty"`
}

type WithdrawalRequest struct {
	Reference   string   `json:"reference"`
	Amount      float64  `json:"amount"`
	Account     string   `json:"account"`
	CallbackURL *string  `json:"callback_url,omitempty"`
	Details     *Details `json:"details,omitempty"`
}

type WithdrawalResponse struct {
	ID            string  `json:"id"`
	Reference     string  `json:"reference"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	StatusCode    int     `json:"status_code,omitempty"`
	StatusMessage string  `json:"status_message,omitempty"`
	PaymentFee    float64 `json:"payment_fee,omitempty"`
	AmountSpent   float64 `json:"amount_spent,omitempty"`
}
