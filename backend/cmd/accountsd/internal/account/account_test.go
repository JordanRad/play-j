package account_test

import (
	"context"
	"errors"

	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account/accountfakes"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
	accountgen "github.com/JordanRad/play-j/backend/internal/accountservice/gen/account"
	payment_pb "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/grpc/payment/pb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var TestToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA"

var _ = Describe("Account Service", func() {

	var (
		fakeStore          *accountfakes.FakeStore
		fakePaymentService *accountfakes.FakePaymentService
		service            *account.Service
	)

	BeforeEach(func() {
		fakeStore = new(accountfakes.FakeStore)
		fakePaymentService = new(accountfakes.FakePaymentService)
		service = account.NewService(fakeStore, fakePaymentService)
	})

	Describe("Register", func() {
		var (
			ctx      context.Context
			response *accountgen.RegisterResponse
			err      error
		)
		Describe("Given matching passwords", func() {

			JustBeforeEach(func() {
				ctx = context.Background()
				//ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
				payload := &accountgen.RegisterPayload{
					FirstName:         "Test",
					LastName:          "Test",
					Email:             "test@mail.com",
					Password:          "123456789",
					ConfirmedPassword: "123456789",
				}
				response, err = service.Register(ctx, payload)
			})

			When("new account is created successfully", func() {
				BeforeEach(func() {
					fakeStore.CreateUserReturns(true, nil)
				})
				It("should return a success message", func() {
					Expect(response.Message).To(Equal("Successfully created."))

				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

			When("new account is NOT created successfully", func() {
				BeforeEach(func() {
					fakeStore.CreateUserReturns(false, errors.New("such user already exists"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Describe("Given NOT matching passwords", func() {

			JustBeforeEach(func() {
				ctx = context.Background()
				//ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
				payload := &accountgen.RegisterPayload{
					FirstName:         "Test",
					LastName:          "Test",
					Email:             "test@mail.com",
					Password:          "123456789",
					ConfirmedPassword: "12345679",
				}
				response, err = service.Register(ctx, payload)
			})

			When("new account is NOT created successfully", func() {
				It("should return passwords do not match", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("Login", func() {
		Describe("Given correct password", func() {
			var (
				ctx      context.Context
				response *accountgen.LoginResponse
				err      error
			)
			JustBeforeEach(func() {
				ctx = context.Background()
				//ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
				payload := &accountgen.LoginPayload{
					Email:    "test@mail.com",
					Password: "1234567",
				}
				response, err = service.Login(ctx, payload)
			})

			When("login is successful and the service return a valid JWT", func() {
				var validAccount *dbmodels.Account
				BeforeEach(func() {
					validAccount = &dbmodels.Account{
						ID:       1,
						Email:    "test@mail.com",
						Password: "$2a$10$6xReJuSpkcfWdoEOQ5ibvuErZHCjyK2DgP8CZnMHwB5mK9hGXUmce",
					}
					fakeStore.GetUserByEmailReturns(validAccount, nil)
				})

				var hasMoreThan25Chars = func(l int) bool { return l > 25 }

				It("should return a success message", func() {
					Expect(response.Email).To(Equal(validAccount.Email))
					Expect(response.Token).To(ContainSubstring("ey"))
					Expect(len(response.Token)).To(Satisfy(hasMoreThan25Chars))

				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

			When("login is NOT successful due to wrong email and the service returns invalid credentials", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("invalid credentials"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Describe("Given INCORECT password", func() {
			var (
				ctx      context.Context
				response *accountgen.LoginResponse
				err      error
			)
			JustBeforeEach(func() {
				ctx = context.Background()
				//ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
				payload := &accountgen.LoginPayload{
					Email:    "test@mail.com",
					Password: "12345678910",
				}
				response, err = service.Login(ctx, payload)
			})

			When("login is successful and the service return a valid JWT", func() {
				var validAccount *dbmodels.Account
				BeforeEach(func() {
					validAccount = &dbmodels.Account{
						ID:       1,
						Email:    "test@mail.com",
						Password: "$2a$10$6xReJuSpkcfWdoEOQ5ibvuErZHCjyK2DgP8CZnMHwB5mK9hGXUmce",
					}
					fakeStore.GetUserByEmailReturns(validAccount, nil)
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
					Expect(response).To(BeNil())
				})

			})
		})
	})

	Describe("GetPaymentsByAccountID", func() {
		Describe("Given a valid JWT", func() {
			var (
				ctx      context.Context
				response *accountgen.ProfileResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &accountgen.GetProfilePayload{
					PaymentsLimit: 3,
				}
				response, err = service.GetProfile(ctx, payload)
			})

			When("such account exists", func() {
				account := &dbmodels.Account{
					ID:        1,
					Email:     "joesmith@abv.bg",
					FirstName: "Batman",
					LastName:  "Smith",
					Username:  "underground_rat",
				}
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(account, nil)
				})
				When("and payment service responds with OK", func() {
					payment := &payment_pb.PaymentResponse{
						Id:            1,
						CreatedAt:     "02/05/2022",
						PaymentNumber: "JJSHD_121qq",
						Amount:        11.99,
					}
					list := make([]*payment_pb.PaymentResponse, 0, 1)
					list = append(list, payment)
					payments := &payment_pb.GetPaymentsByAccountIDResponse{
						Total:     1,
						Resources: list,
					}
					BeforeEach(func() {
						fakePaymentService.GetPaymentsByAccountIDReturns(payments, nil)
					})
					It("should return profile details and payments", func() {
						Expect(response.Email).To(Equal(account.Email))
						Expect(response.FirstName).To(Equal(account.FirstName))
						Expect(response.LastName).To(Equal(account.LastName))
						Expect(response.Username).To(Equal(account.Username))
						Expect(len(response.LastPayments)).To(Equal(1))
						Expect(response.LastPayments[0].ID).To(Equal(uint(payment.Id)))
						Expect(response.LastPayments[0].PaymentNumber).To(Equal(payment.PaymentNumber))
						Expect(response.LastPayments[0].CreatedAt).To(Equal(payment.CreatedAt))
					})
					It("should not return an error", func() {
						Expect(err).To(Not(HaveOccurred()))
					})
				})

				When("but payment service responds with Error", func() {
					BeforeEach(func() {
						fakePaymentService.GetPaymentsByAccountIDReturns(nil, errors.New("grpc error"))
					})
					It("should return profile details and payments", func() {
						Expect(response.Email).To(Equal(account.Email))
						Expect(response.FirstName).To(Equal(account.FirstName))
						Expect(response.LastName).To(Equal(account.LastName))
						Expect(response.Username).To(Equal(account.Username))
						Expect(response.LastPayments).To(BeNil())
					})
					It("should not return an error", func() {
						Expect(err).To(Not(HaveOccurred()))
					})
				})

			})

			When("such account does not exist", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("Such account does not exist"))
				})
				It("should return an erorr", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("such account does not exist"))
				})
			})
		})

		Describe("Given an INVALID JWT", func() {
			var (
				ctx      context.Context
				response *accountgen.ProfileResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &accountgen.GetProfilePayload{
					PaymentsLimit: 3,
				}
				response, err = service.GetProfile(ctx, payload)
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
