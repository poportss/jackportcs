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

	customer, err := utils.CreateCustomer(requestCustomer)
	if err != nil {
		return nil, err
	}

	log.Printf("Tipo de customer.Metadata: %T\n", customer.Metadata)

	metadataJSON, err := json.Marshal(customer.Metadata)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar metadata: %v", err)
	}

	err = s.DB.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"pagarme_customer_id":       customer.ProviderCustomerID,
		"pagarme_customer_metadata": string(metadataJSON),
	}).Error
	if err != nil {
		return nil, err
	}

	customerResponse := &dto.PagarmeCreateCustomerResponse{
		Name:         customer.Name,
		Email:        customer.Email,
		Document:     customer.Document,
		DocumentType: customer.DocumentType,
		Type:         customer.Type,
	}

	return customerResponse, nil
}

func (s *srv) createCustomerCard(createCard *dto.PagarmeCreateCardRequest, userID uuid.UUID) (*dto.CardResponse, error) {

	var user models.User
	if err := s.DB.Model(&models.User{}).Preload("Address").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	card, err := utils.CreateCard(createCard, user)
	if err != nil {
		return nil, err
	}

	if err = s.DB.Model(&models.Card{}).Create(&card).Error; err != nil {
		return nil, err
	}

	cardResponse := &dto.CardResponse{
		ID:             card.ID,
		CreatedAt:      card.CreatedAt,
		LastFourDigits: card.LastFourDigits,
		PaymentMethod:  card.PaymentMethod,
		HolderName:     card.HolderName,
		HolderDocument: card.HolderDocument,
		ExpMonth:       card.ExpMonth,
		ExpYear:        card.ExpYear,
	}

	return cardResponse, nil
}
