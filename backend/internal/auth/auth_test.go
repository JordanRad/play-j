package auth_test

import (
	"github.com/JordanRad/play-j/backend/internal/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {

	Describe("EncryptPassword", func() {
		var (
			hash string
			err  error
		)

		When("encrypts password and returns hashed string", func() {
			JustBeforeEach(func() {
				hash, err = auth.EncryptPassword("12345678")
			})
			It("should return a valid hash", func() {
				Expect(hash).To(ContainSubstring("$2a$10$"))
				Expect(err).To(Not(HaveOccurred()))
			})
		})
	})

	Describe("CheckPassword", func() {
		var funcResult bool

		When("verifies a correct password and returns true", func() {
			JustBeforeEach(func() {
				funcResult = auth.CheckPassword("$2a$10$6xReJuSpkcfWdoEOQ5ibvuErZHCjyK2DgP8CZnMHwB5mK9hGXUmce", "1234567")
			})
			It("should return a valid hash", func() {
				Expect(funcResult).To(BeTrue())

			})
		})

		When("verifies a wrong password and returns false", func() {
			JustBeforeEach(func() {
				funcResult = auth.CheckPassword("$2a$10$6xReJuSpkcfWdoEOQ5ibvuErZHCjyK2DgP8CZnMHwB5mK9hGXUmce", "defwrongpass")
			})
			It("should return a valid hash", func() {
				Expect(funcResult).To(BeFalse())
			})
		})
	})

	Describe("GenerateJWT", func() {
		var (
			token string
			err   error
		)

		When("generate a token successfully", func() {
			JustBeforeEach(func() {
				token, err = auth.GenerateJWT(1, "test@test.com")
			})
			It("should return a valid hash", func() {
				Expect(token).To(ContainSubstring("ey"))
				Expect(err).To(Not(HaveOccurred()))
			})
		})

	})

	Describe("ValidateJWT", func() {
		var (
			token   string
			isValid bool
			err     error
		)

		When("the given token is valid", func() {
			JustBeforeEach(func() {
				token, err = auth.GenerateJWT(1, "test@test.com")
				isValid, err = auth.ValidateJWT(token)
			})
			It("should return true", func() {
				Expect(token).To(ContainSubstring("ey"))
				Expect(isValid).To(BeTrue())
				Expect(err).To(Not(HaveOccurred()))
			})
		})

		When("the given token is invalid", func() {
			JustBeforeEach(func() {
				isValid, err = auth.ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6InRlc3RAYWJ2LmJnIn0sImV4cCI6MTY1MDIwMDI0NywiaXNzIjoicGxheWotYWNjb3VudC1zZXJ2aWNlIn0.qZGIlfE5HHw1O2fatYHqn8HPo8gLEngEC8okQUcEBUY")
			})
			It("should return false", func() {
				Expect(isValid).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
