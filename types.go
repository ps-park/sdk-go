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
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

type BillingInfo struct {
	Address     *string `json:"address,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Country     *string `json:"country,omitempty"`
}

type BankInfo struct {
	Account *string `json:"account,omitempty"`
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type CardData struct {
	ExpMonth *string `json:"exp_month,omitempty"`
	ExpYear  *string `json:"exp_year,omitempty"`
}

type WebData struct {
	IP        *string `json:"ip,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
}

type UISchema struct {
	Language *string `json:"language,omitempty"` // Values: "en", "ua", "ru"
}

type FlowData struct {
	Action string        `json:"action"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type InvoiceRequest struct {
	Reference   string        `json:"reference"`
	Title       *string       `json:"title,omitempty"`
	Description *string       `json:"description,omitempty"`
	LimitMinute *int          `json:"limit_minute,omitempty"`
	CallbackURL *string       `json:"callback_url,omitempty"`
	Amount      float64       `json:"amount"`
	Currency    string        `json:"currency"`
	ReturnURL   string        `json:"return_url"`
	Customer    *CustomerInfo `json:"customer,omitempty"`
	BillingInfo *BillingInfo  `json:"billing_info,omitempty"`
	Bank        *BankInfo     `json:"bank,omitempty"`
	CardData    *CardData     `json:"card_data,omitempty"`
	WebData     *WebData      `json:"web_data,omitempty"`
	UI          *UISchema     `json:"ui,omitempty"`
}

type InvoiceResponse struct {
	ID            string    `json:"id"`
	Reference     string    `json:"reference"`
	WalletID      string    `json:"wallet_id"`
	Currency      string    `json:"currency"`
	Amount        float64   `json:"amount"`
	AmountInitial float64   `json:"amount_initial"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	StatusCode    int       `json:"status_code,omitempty"`
	StatusMessage string    `json:"status_message,omitempty"`
	PaymentFee    float64   `json:"payment_fee"`
	Address       string    `json:"address,omitempty"`
	Memo          string    `json:"memo,omitempty"`
	FlowData      *FlowData `json:"flowData,omitempty"`
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
	AmountInitial float64 `json:"amount_initial"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	StatusCode    int     `json:"status_code,omitempty"`
	StatusMessage string  `json:"status_message,omitempty"`
	PaymentFee    float64 `json:"payment_fee"`
	Address       string  `json:"address,omitempty"`
	Memo          string  `json:"memo,omitempty"`
}

type WithdrawalDetails struct {
	Crypto   *CryptoInfo   `json:"crypto,omitempty"`
	Customer *CustomerInfo `json:"customer,omitempty"`
	Billing  *BillingInfo  `json:"billing_info,omitempty"`
	Bank     *BankInfo     `json:"bank,omitempty"`
	CardData *CardData     `json:"card_data,omitempty"`
	WebData  *WebData      `json:"web_data,omitempty"`
}

type WithdrawalRequest struct {
	Reference string             `json:"reference"`
	Amount    float64            `json:"amount"`
	Account   string             `json:"account"`
	Details   *WithdrawalDetails `json:"details,omitempty"`
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
	PaymentFee    float64 `json:"payment_fee"`
	AmountSpent   float64 `json:"amount_spent"`
}
