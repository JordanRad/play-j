package dbsubscriptions

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	subscription "github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/subscription"
	"github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

var _ subscription.Store = (*Store)(nil)

func (s *Store) GetAccountSubscription(ctx context.Context, accountID uint) (*dbmodels.Subscription, error) {
	row := s.DB.QueryRow(`SELECT * FROM subscriptions WHERE accountID = $1 ORDER BY validUntil ASC LIMIT 1;`, accountID)

	var result dbmodels.Subscription
	err := row.Scan(
		&result.ID,
		&result.ValidUntil,
		(*pq.Int32Array)(&result.LinkedAccountIDs),
		&result.SubscriptionType,
		&result.AccountID)

	if err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}

	return &result, nil
}

func (s *Store) CreateSubscription(ctx context.Context, accountID uint) error {
	result, err := s.DB.Exec("INSERT INTO subscriptions (validuntil, linkedaccountids, subscriptiontype,accountid) VALUES ((select date_trunc('day', NOW() + interval '1 month')),$1,$2,$3);", "{10,20,30}", "family", accountID)

	if err != nil {
		return fmt.Errorf("error creating a new subscription: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 1 {
		return nil
	}

	return fmt.Errorf("error inserting a subscription")
}
