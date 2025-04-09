package payment

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"github.com/poportss/jackportcs/internal/utils"
	"log"
)

func (s *srv) createPaymentOrder(createOrderRequest *dto.CreateOrderRequest, userID uuid.UUID) (*dto.CreateOrderResponse, error) {

	if createOrderRequest.CardID != nil {
		var card models.Card
		s.DB.Model(&models.Card{}).Where("id = ? ", createOrderRequest.CardID).Find(&card)

		createOrderRequest.Card = &card
	}

	var user models.User
	s.DB.Model(&models.User{}).Where("id = ? ", userID).Find(&user)

	paymentOrder, err := utils.CreatePaymentOrder(user, createOrderRequest)
	if err != nil {
		return nil, err
	}

	err = s.DB.Model(&models.PaymentOrder{}).Create(&paymentOrder).Error
	if err != nil {
		return nil, err
	}

	paymentOrderResponse := &dto.CreateOrderResponse{
		ID:            paymentOrder.ID,
		CreatedAt:     paymentOrder.CreatedAt,
		Amount:        paymentOrder.Amount,
		Status:        string(paymentOrder.Status),
		PaymentMethod: string(paymentOrder.PaymentMethod),
		PIX:           paymentOrder.PIX,
	}

	return paymentOrderResponse, nil
}

func (s *srv) createPaymentCustomer(requestCustomer *dto.PagarmeCreateCustomerRequest, userID uuid.UUID) (*dto.PagarmeCreateCustomerResponse, error) {

	// Create customer using the external utility
	customer, err := utils.CreateCustomer(requestCustomer)
	if err != nil {
		return nil, err
	}

	log.Printf("Tipo de customer.Metadata: %T\n", customer.Metadata)

	// Verifique o tipo e converta para um tipo adequado
	metadataJSON, err := json.Marshal(customer.Metadata)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar metadata: %v", err)
	}

	// Atualize os dados no banco de dados
	err = s.DB.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"pagarme_customer_id":       customer.ProviderCustomerID,
		"pagarme_customer_metadata": string(metadataJSON), // Converted para string JSON
	}).Error
	if err != nil {
		return nil, err
	}

	// Prepare response structure
	customerResponse := &dto.PagarmeCreateCustomerResponse{
		Name:         customer.Name,
		Email:        customer.Email,
		Document:     customer.Document,
		DocumentType: customer.DocumentType,
		Type:         customer.Type,
	}

	return customerResponse, nil
}
