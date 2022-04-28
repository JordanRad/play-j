package payments

import (
	"context"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/internal/auth"
	"github.com/JordanRad/play-j/backend/internal/paymentservice/gen/payment"
)

type PaymentStore interface {
	GetAccountPayments(context.Context, uint, uint) ([]*dbmodels.Payment, error)
	CreatePayment(context.Context, uint, float32) (string, error)
}

type SubscriptionStore interface {
	CreateSubscription(context.Context, uint) error
}

type Service struct {
	paymentStore      PaymentStore
	subscriptionStore SubscriptionStore
}

func NewService(ps PaymentStore, ss SubscriptionStore) *Service {
	return &Service{
		paymentStore:      ps,
		subscriptionStore: ss,
	}
}

// Compile-time assertion that *Service implements the payment.Service interface
var _ payment.Service = (*Service)(nil)

// Converts a database entity into data transfer object
func toPaymentResponse(paymentEntity *dbmodels.Payment) *payment.PaymentResponse {
	return &payment.PaymentResponse{
		ID:            paymentEntity.ID,
		CreatedAt:     paymentEntity.CreatedAt.String(),
		PaymentNumber: paymentEntity.PaymentNumber,
		Amount:        paymentEntity.Amount,
	}
}

func (s *Service) GetAccountPayments(ctx context.Context, p *payment.GetAccountPaymentsPayload) (*payment.PaymentListResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	paymentEntities, err := s.paymentStore.GetAccountPayments(ctx, tokenClaims.AccountID, uint(p.Limit))

	if err != nil {
		return nil, fmt.Errorf("error getting payments: %w", err)
	}

	paymentResponseList := make([]*payment.PaymentResponse, 0, len(paymentEntities))

	for _, paymentEntity := range paymentEntities {
		payment := toPaymentResponse(paymentEntity)
		paymentResponseList = append(paymentResponseList, payment)
	}

	response := &payment.PaymentListResponse{
		Total:     uint(len(paymentEntities)),
		Resources: paymentResponseList,
	}
	return response, nil
}

func (s *Service) GetPaymentsByAccountID(ctx context.Context, p *payment.GetPaymentsByAccountIDPayload) (*payment.PaymentListResponse, error) {

	paymentEntities, err := s.paymentStore.GetAccountPayments(ctx, uint(p.AccountID), uint(p.Limit))

	if err != nil {
		return nil, fmt.Errorf("error getting payments: %w", err)
	}

	paymentResponseList := make([]*payment.PaymentResponse, 0, len(paymentEntities))

	for _, paymentEntity := range paymentEntities {
		payment := toPaymentResponse(paymentEntity)
		paymentResponseList = append(paymentResponseList, payment)
	}

	response := &payment.PaymentListResponse{
		Total:     uint(len(paymentEntities)),
		Resources: paymentResponseList,
	}
	return response, nil
}

func (s *Service) CreateAccountPayment(ctx context.Context, p *payment.CreateAccountPaymentPayload) (*payment.TransactionResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	//TODO: Turn this operation into one transaction
	// Insert the payment in the database
	paymentNumber, err := s.paymentStore.CreatePayment(ctx, tokenClaims.AccountID, 11.99)
	if err != nil {
		return nil, fmt.Errorf("error creating a payment: %w", err)
	}

	// Insert a new subscription in the database
	err = s.subscriptionStore.CreateSubscription(ctx, tokenClaims.AccountID)

	if err != nil {
		return nil, fmt.Errorf("error creating a payment: %w", err)
	}

	response := &payment.TransactionResponse{
		PaymentNumber: paymentNumber,
		Message:       "Payment completed successfully",
	}

	return response, nil
}
