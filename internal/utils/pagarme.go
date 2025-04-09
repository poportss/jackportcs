package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
)

func CreatePaymentOrder(user models.User, order *dto.CreateOrderRequest) (*models.PaymentOrder, error) {
	if order.Amount < 5000 {
		return nil, errors.New("valor da compra não ultrapassa o valor mínimo de compra")
	}

	req := models.PagarmeOrder{
		Closed:     true,
		Currency:   "BRL",
		CustomerID: user.PagarmeCustomerID,
		Items: []models.PagarmeOrderItem{
			{
				Code:        "A1",
				Quantity:    1,
				Description: "Pagamento referente ao JackportCS",
				Amount:      order.Amount,
			},
		},
		Payments: []models.PagarmeOrderPayment{
			{
				PaymentMethod: string(order.PaymentMethod),
			},
		},
	}

	switch order.PaymentMethod {
	case models.PaymentMethodCreditCard:
		req.Payments[0].CreditCard = &models.PagarmeOrderCreditCard{
			Recurrence:          false,
			Installments:        1,
			StatementDescriptor: "Compra via cartão no JackportCS",
			CardID:              order.Card.ProviderCardID,
		}
	case models.PaymentMethodPIX:
		req.Payments[0].PIX = &models.PagarmeOrderPIX{
			ExpiresIn: "600",
		}
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar pedido para o Pagar.me: %w", err)
	}

	fmt.Printf(string(reqBytes))

	httpReq, err := http.NewRequest("POST", "https://api.pagar.me/core/v5/orders", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("erro ao montar requisição: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	apiKey := os.Getenv("PAGARME_SECRET_KEY")

	httpReq.SetBasicAuth(apiKey, "") // ✅ Chave secreta de teste

	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar pedido para o Pagar.me: %w", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta do Pagar.me: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("erro na resposta do Pagar.me: %s", resBody)
	}

	fmt.Printf("Resposta do Pagar.me: %s\n", resBody)

	var pagarmeOrder models.PagarmeOrder
	if err := json.Unmarshal(resBody, &pagarmeOrder); err != nil {
		return nil, fmt.Errorf("erro ao parsear resposta do Pagar.me: %w", err)
	}

	paymentOrder := &models.PaymentOrder{
		UserID:            user.ID,
		CustomerID:        user.PagarmeCustomerID,
		ReferenceID:       pagarmeOrder.ID,
		Amount:            order.Amount,
		Status:            models.OrderStatus(pagarmeOrder.Status),
		PaymentMethod:     order.PaymentMethod,
		ProviderResponses: resBody,
	}

	if paymentOrder.PaymentMethod == models.PaymentMethodPIX && pagarmeOrder.Charges != nil {
		paymentOrder.PIX = &models.PagarmeOrderPIX{
			QRCode:    pagarmeOrder.Charges[0].LastTransaction.QRCode,
			QRCodeURL: pagarmeOrder.Charges[0].LastTransaction.QRCodeURL,
			ExpiresAt: pagarmeOrder.Charges[0].LastTransaction.ExpiresAt,
		}
	}

	return paymentOrder, nil
}

func CreateCustomer(requestCustomer *dto.PagarmeCreateCustomerRequest) (*models.PagarmeCreateCustomerResponse, error) {
	var pagarmeCustomerResponse *models.PagarmeCreateCustomerResponse

	if len(requestCustomer.Name) > 63 {
		requestCustomer.Name = requestCustomer.Name[:63]
	}

	reqBytes, err := json.Marshal(requestCustomer)
	if err != nil {
		return nil, fmt.Errorf("failed to create the customer on pagarme provider: %w", err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.pagar.me/core/v5/customers", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("erro ao montar requisição: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	apiKey := os.Getenv("PAGARME_SECRET_KEY")

	httpReq.SetBasicAuth(apiKey, "")

	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar pedido para o Pagar.me: %w", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create the customer on pagarme provider: %w", err)
	}

	if err := json.Unmarshal(resBody, &pagarmeCustomerResponse); err != nil {
		return nil, fmt.Errorf("failed to create the customer on pagarme provider: %w", err)
	}

	pagarmeCustomer := pagarmeCustomerResponse
	pagarmeCustomer.Metadata = resBody
	pagarmeCustomer.ProviderCustomerID = pagarmeCustomerResponse.ID

	return pagarmeCustomer, nil
}
