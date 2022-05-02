package subscription_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/subscription"
	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/subscription/subscriptionfakes"
	subgen "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/subscription"
)

var TestToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA"
var _ = Describe("Subscription", func() {
	var (
		fakeStore *subscriptionfakes.FakeStore
		service   *subscription.Service
	)
	BeforeEach(func() {
		fakeStore = new(subscriptionfakes.FakeStore)
		service = subscription.NewService(fakeStore)
	})

	Describe("GetAccounSubscriptionStatus", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *subgen.SubscriptionStatusResponse
				err      error
			)
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)

				payload := &subgen.GetAccountSubscriptionStatusPayload{
					Auth: "Bearer token",
				}

				response, err = service.GetAccountSubscriptionStatus(ctx, payload)
			})
			When("the service returns the account subscription's status successfully", func() {
				subscription := &dbmodels.Subscription{
					ID:               1,
					ValidUntil:       time.Now(),
					LinkedAccountIDs: []int32{32, 33},
					SubscriptionType: "family",
					AccountID:        1,
				}
				BeforeEach(func() {
					fakeStore.GetAccountSubscriptionReturns(subscription, nil)
				})
				It("should return subscrion status successfully", func() {
					Expect(response.Type).To(Equal(subscription.SubscriptionType))
					Expect(response.ValidUntil).To(Equal(subscription.ValidUntil.String()))
				})
				It("should NOT return an error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})
		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *subgen.SubscriptionStatusResponse
				err      error
			)
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")

				payload := &subgen.GetAccountSubscriptionStatusPayload{
					Auth: "Bearer token",
				}

				response, err = service.GetAccountSubscriptionStatus(ctx, payload)
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
