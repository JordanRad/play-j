package subscription

import (
	"context"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/internal/auth"
	"github.com/JordanRad/play-j/backend/internal/paymentservice/gen/subscription"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Store
type Store interface {
	GetAccountSubscription(context.Context, uint) (*dbmodels.Subscription, error)
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
var _ subscription.Service = (*Service)(nil)

func (s *Service) GetAccountSubscriptionStatus(ctx context.Context, p *subscription.GetAccountSubscriptionStatusPayload) (*subscription.SubscriptionStatusResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	subscriptionEntity, err := s.store.GetAccountSubscription(ctx, tokenClaims.AccountID)

	if err != nil {
		return nil, fmt.Errorf("error getting account subscription status: %w", err)
	}

	response := &subscription.SubscriptionStatusResponse{
		HasPaid:    true,
		ValidUntil: subscriptionEntity.ValidUntil.String(),
		Type:       subscriptionEntity.SubscriptionType,
	}
	return response, nil
}
