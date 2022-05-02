package payments_test

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	payments "github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/payment"
	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/payment/paymentfakes"
	"github.com/JordanRad/play-j/backend/internal/paymentservice/gen/payment"
	paymentgen "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/payment"
)

var TestToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA"
var _ = Describe("Payment Service", func() {
	var (
		fakePaymentStore      *paymentfakes.FakePaymentStore
		fakeSubscriptionStore *paymentfakes.FakeSubscriptionStore
		service               *payments.Service
	)

	BeforeEach(func() {
		fakePaymentStore = new(paymentfakes.FakePaymentStore)
		fakeSubscriptionStore = new(paymentfakes.FakeSubscriptionStore)
		service = payments.NewService(fakePaymentStore, fakeSubscriptionStore)
	})

	Describe("GetAccountsByAccountID", func() {
		var (
			ctx      context.Context
			response *payment.PaymentListResponse
			err      error
		)
		JustBeforeEach(func() {
			ctx = context.Background()

			payload := &paymentgen.GetPaymentsByAccountIDPayload{
				AccountID: 1,
				Limit:     1,
			}
			response, err = service.GetPaymentsByAccountID(ctx, payload)
		})
		When("the service returns the payments successfully", func() {
			payment := &dbmodels.Payment{
				ID:            1,
				CreatedAt:     time.Now(),
				PaymentNumber: "1234",
				PaidBy:        1,
				Amount:        11.99,
			}
			BeforeEach(func() {
				payments := make([]*dbmodels.Payment, 0)
				payments = append(payments, payment)
				fakePaymentStore.GetAccountPaymentsReturns(payments, nil)
			})
			It("should return a list of payments", func() {
				Expect(response.Total).To(Equal(uint(1)))
				Expect(response.Resources[0].Amount).To(Equal(payment.Amount))
				Expect(response.Resources[0].ID).To(Equal(payment.ID))
				Expect(response.Resources[0].PaymentNumber).To(Equal(payment.PaymentNumber))
			})
			It("should NOT return an error", func() {
				Expect(err).To(Not(HaveOccurred()))
			})
		})
		When("the service DOES NOT return the payments successfully", func() {
			BeforeEach(func() {
				fakePaymentStore.GetAccountPaymentsReturns(nil, errors.New("error"))
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("CreateAccountPayment", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *payment.TransactionResponse
				err      error
			)
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)

				payload := &paymentgen.CreateAccountPaymentPayload{
					Auth: "Bearer Token",
				}
				response, err = service.CreateAccountPayment(ctx, payload)
			})
			When("the service creates new payment successfully", func() {
				When("the service returns the payments successfully", func() {
					BeforeEach(func() {
						fakePaymentStore.CreatePaymentReturns("123456", nil)
					})
					When("and subscription service creates new subscription successfully", func() {

						BeforeEach(func() {
							fakeSubscriptionStore.CreateSubscriptionReturns(nil)
						})

						It("should return a list of payments", func() {
							Expect(response.PaymentNumber).To(Equal("123456"))
							Expect(response.Message).To(Equal("Payment completed successfully"))
						})
						It("should NOT return an error", func() {
							Expect(err).To(Not(HaveOccurred()))
						})
					})
				})
			})
			When("the service DOES NOT create new payments successfully", func() {

				BeforeEach(func() {
					fakePaymentStore.CreatePaymentReturns("", errors.New("error"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *payment.TransactionResponse
				err      error
			)
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")

				payload := &paymentgen.CreateAccountPaymentPayload{
					Auth: "Bearer Token",
				}
				response, err = service.CreateAccountPayment(ctx, payload)
			})
			When("operation is unathorized", func() {
				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})

	})
})
