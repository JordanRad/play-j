package dbmodels

import "time"

type Payment struct {
	ID            uint
	CreatedAt     time.Time
	PaymentNumber string
	PaidBy        uint
	Amount        float32
}

type Subscription struct {
	ID               uint
	ValidUntil       time.Time
	LinkedAccountIDs []int32
	SubscriptionType string
	AccountID        uint
}

type PaymentDetails struct {
	ID     uint
	Number string
}
