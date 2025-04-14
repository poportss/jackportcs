package models

import (
	"time"
)

type PagarmeCreateCardRequest struct {
	Number         string `json:"number"`
	HolderName     string `json:"holder_name"`
	HolderDocument string `json:"holder_document"`
	ExpMonth       int    `json:"exp_month"`
	ExpYear        int    `json:"exp_year"`
	CVV            string `json:"cvv"`

	BillingAddress PagarmeAddress           `json:"billing_address"`
	Options        PagarmeCreateCardOptions `json:"options"`
}

type PagarmeCreateCardResponse struct {
	ID             string                  `json:"id"`
	FirstSixDigits string                  `json:"first_six_digits"`
	LastFourDigits string                  `json:"last_four_digits"`
	Brand          string                  `json:"brand"`
	HolderName     string                  `json:"holder_name"`
	HolderDocument string                  `json:"holder_document"`
	ExpMonth       int                     `json:"exp_month"`
	ExpYear        int                     `json:"exp_year"`
	Status         string                  `json:"status"`
	Type           string                  `json:"type"`
	CreatedAt      string                  `json:"created_at"`
	UpdatedAt      string                  `json:"updated_at"`
	BillingAddress PagarmeAddress          `json:"billing_address"`
	Customer       PagarmeCustomerResponse `json:"customer"`
}

type PagarmeCreateCustomerRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	Type         string `json:"type"`

	Address PagarmeAddress `json:"address"`
	Phones  PagarmePhones  `json:"phones"`
}

type PagarmeCreateCustomerResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	Type         string `json:"type"`
	Delinquent   bool   `json:"delinquent"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`

	Address            PagarmeAddress `json:"address"`
	Phones             PagarmePhones  `json:"phones"`
	Metadata           []byte         `json:"metadata,omitempty"`
	ProviderCustomerID string         `json:"provider_customer_id,omitempty"`
}

type PagarmeAddress struct {
	ID        string `json:"id,omitempty"`
	Line1     string `json:"line_1"`
	Line2     string `json:"line_2"`
	ZipCode   string `json:"zip_code"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type PagarmePhones struct {
	HomePhone PagarmeHomePhone `json:"home_phone"`
}

type PagarmeHomePhone struct {
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
}

type PagarmeOrder struct {
	ID         string `json:"id,omitempty"`
	Code       string `json:"code,omitempty"`
	Closed     bool   `json:"closed,omitempty"`
	Currency   string `json:"currency,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Status     string `json:"status,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	ClosedAt   string `json:"closed_at,omitempty"`

	Charges []struct {
		LastTransaction struct {
			Line      string `json:"line"`
			Barcode   string `json:"barcode"`
			QRCode    string `json:"qr_code"`
			QRCodeURL string `json:"qr_code_url"`
			URL       string `json:"url"`
			PDF       string `json:"pdf"`
			DueAt     string `json:"due_at"`
			ExpiresAt string `json:"expires_at"`
		} `json:"last_transaction,omitempty"`
	} `json:"charges,omitempty"`

	Items    []PagarmeOrderItem    `json:"items,omitempty"`
	Payments []PagarmeOrderPayment `json:"payments,omitempty"`
}

type PagarmeOrderItem struct {
	ID          string `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type PagarmeOrderPayment struct {
	PaymentMethod string                  `json:"payment_method,omitempty"`
	CreditCard    *PagarmeOrderCreditCard `json:"credit_card,omitempty"`
	Boleto        *PagarmeOrderBoleto     `json:"boleto,omitempty"`
	PIX           *PagarmeOrderPIX        `json:"pix,omitempty"`
	Split         []PagarmeOrderSplit     `json:"split,omitempty"`
}

type PagarmeOrderBoleto struct {
	Bank         string `json:"bank,omitempty"`
	Instructions string `json:"instructions,omitempty"`
	DueAt        string `json:"due_at,omitempty"`
}

type PagarmeOrderPIX struct {
	ExpiresIn             string `json:"expires_in"`
	AdditionalInformation []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"additional_information,omitempty"`
	QRCode    string `json:"qrCode,omitempty"`
	QRCodeURL string `json:"qrCodeURL,omitempty"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

type PagarmeOrderCreditCard struct {
	Recurrence          bool   `json:"recurrence,omitempty"`
	Installments        int    `json:"installments,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	CardID              string `json:"card_id,omitempty"`
}

type PagarmeOrderSplit struct {
	RecipientID string                   `json:"recipient_id,omitempty"`
	Amount      uint                     `json:"amount,omitempty"`
	Type        string                   `json:"type,omitempty"`
	Options     PagarmeOrderSplitOptions `json:"options,omitempty"`
}

type PagarmeOrderSplitOptions struct {
	ChargeProcessingFee bool `json:"charge_processing_fee,omitempty"`
	ChargeRemainderFee  bool `json:"charge_remainder_fee,omitempty"`
	Liable              bool `json:"liable,omitempty"`
}

type PagarmeOrderHook struct {
	ID        string       `json:"id"`
	Event     string       `json:"event"`
	Status    string       `json:"string"`
	CreatedAt time.Time    `json:"created_at"`
	Data      PagarmeOrder `json:"data"`
}

type PagarmeCreateCard struct {
	Number         string `json:"number"`
	HolderName     string `json:"holder_name"`
	HolderDocument string `json:"holder_document"`
	ExpMonth       int    `json:"exp_month"`
	ExpYear        int    `json:"exp_year"`
	CVV            string `json:"cvv"`

	BillingAddress PagarmeAddress           `json:"billing_address"`
	Options        PagarmeCreateCardOptions `json:"options"`
}

type PagarmeCreateCardOptions struct {
	VerifyCard bool `json:"verify_card"`
}

type PagarmeCustomerResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	Type         string `json:"type"`
	Delinquent   bool   `json:"delinquent"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`

	Customer struct {
		Phones struct {
			CountryCode string `json:"country_code"`
			Number      string `json:"number"`
			AreaCode    string `json:"area_code"`
		} `json:"home_phone"`
	} `json:"phones"`
}
