package dbpayments

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	payment "github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/payment"
	"github.com/google/uuid"
)

type Store struct {
	DB *sql.DB
}

var _ payment.PaymentStore = (*Store)(nil)

func generatePaymentNumber(accountID uint) string {
	uniqueString := uuid.New()
	return fmt.Sprintf("%v-PJ%v-%v", time.Now().Unix(), accountID, uniqueString)
}

func (s *Store) GetAccountPayments(ctx context.Context, accountID uint, limit uint) ([]*dbmodels.Payment, error) {
	rows, err := s.DB.Query("SELECT * FROM payments WHERE paidBy = $1 ORDER BY createdAt DESC LIMIT $2;", accountID, limit)

	if err != nil {
		return nil, fmt.Errorf("error extracting accounts's payments: %w", err)
	}
	defer rows.Close()

	payments := make([]*dbmodels.Payment, 0)
	for rows.Next() {
		payment := &dbmodels.Payment{}

		err := rows.Scan(
			&payment.ID,
			&payment.CreatedAt,
			&payment.PaymentNumber,
			&payment.Amount,
			&payment.PaidBy,
		)

		if err != nil {
			return nil, fmt.Errorf("error mapping a payment row: %w", err)
		}

		payments = append(payments, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return payments, nil
}

func (s *Store) CreatePayment(ctx context.Context, accountID uint, amount float32) (string, error) {
	paymentNumber := generatePaymentNumber(accountID)

	result, err := s.DB.Exec("INSERT INTO payments (paymentnumber,paidby,amount) VALUES ($1,$2,$3);", paymentNumber, accountID, amount)

	if err != nil {
		return "", fmt.Errorf("error creating a new payment: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return "", fmt.Errorf("rows affected error: %w", err)
	}

	if rowsAffected == 1 {
		return paymentNumber, nil
	}

	return "", fmt.Errorf("error inserting a payment")
}
