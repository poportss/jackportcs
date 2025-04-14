package models

type PaymentMethod string

const (
	PaymentMethodPIX        PaymentMethod = "pix"
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodBoleto     PaymentMethod = "boleto"
)
