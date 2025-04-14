package dto

import "github.com/poportss/jackportcs/internal/models"

type PagarmeCreateCustomerRequest struct {
	Name         string                `json:"name"`
	Email        string                `json:"email"`
	Document     string                `json:"document"`
	DocumentType string                `json:"documentType"`
	Type         string                `json:"type"`
	Address      models.PagarmeAddress `json:"address"`
	Phones       models.PagarmePhones  `json:"phones"`
}

type PagarmeCreateCustomerResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	DocumentType string `json:"documentType"`
	Type         string `json:"type"`
}

type PagarmeCreateCardRequest struct {
	Number         string                          `json:"number"`
	HolderName     string                          `json:"holderName"`
	HolderDocument string                          `json:"holderDocument"`
	ExpMonth       int                             `json:"expMonth"`
	ExpYear        int                             `json:"expYear"`
	CVV            string                          `json:"cvv"`
	Type           models.PaymentMethod            `json:"type"`
	Options        models.PagarmeCreateCardOptions `json:"options"`
}

type PagarmeCreateCardResponse struct {
	ID             string                `json:"id"`
	FirstSixDigits string                `json:"firstSixDigits"`
	LastFourDigits string                `json:"lastFourDigits"`
	Brand          string                `json:"brand"`
	HolderName     string                `json:"holderName"`
	HolderDocument string                `json:"holderDocument"`
	ExpMonth       int                   `json:"expMonth"`
	ExpYear        int                   `json:"expYear"`
	Status         string                `json:"status"`
	Type           string                `json:"type"`
	CreatedAt      string                `json:"createdAt"`
	UpdatedAt      string                `json:"updatedAt"`
	BillingAddress models.PagarmeAddress `json:"billingAddress"`
}
