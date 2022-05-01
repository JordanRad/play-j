package account_test

import (
	"context"
	"errors"

	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account/accountfakes"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
	accountgen "github.com/JordanRad/play-j/backend/internal/accountservice/gen/account"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account Service", func() {

	var (
		fakeStore          *accountfakes.FakeStore
		fakePaymentService *accountfakes.FakePaymentService
		service            *account.Service
	)

	BeforeEach(func() {
		fakeStore = new(accountfakes.FakeStore)
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

})
