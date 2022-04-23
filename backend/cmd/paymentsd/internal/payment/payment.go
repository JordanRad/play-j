package payments

import (
	"context"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/internal/auth"
	"github.com/JordanRad/play-j/backend/internal/paymentservice/gen/payment"
)

type Store interface {
	GetAccountPayments(context.Context, uint) ([]*dbmodels.Payment, error)
	CreatePayment(context.Context, uint, float32) (bool, error)
}
type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

// Compile-time assertion that *Service implements the payment.Service interface
var _ payment.Service = (*Service)(nil)

// Converts a database entity into data transfer object
func toPaymentResponse(paymentEntity *dbmodels.Payment) *payment.PaymentResponse {
	return &payment.PaymentResponse{
		ID:            paymentEntity.ID,
		CreatedAt:     paymentEntity.CreatedAt.Format("2006-01-02 15:04"),
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

	paymentEntities, err := s.store.GetAccountPayments(ctx, tokenClaims.AccountID)

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

	ok, err := s.store.CreatePayment(ctx, tokenClaims.AccountID, 25)
	return nil, nil
}
