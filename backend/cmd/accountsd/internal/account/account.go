package account

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/internal/accountservice/gen/account"
	auth "github.com/JordanRad/play-j/backend/internal/auth"
	payment_pb "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/grpc/payment/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// You only need **one** of these per package!
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Store
type Store interface {
	CreateUser(context.Context, *User) (bool, error)
	GetUserByEmail(context.Context, string) (*dbmodels.Account, error)
}

//counterfeiter:generate . PaymentService
type PaymentService interface {
	GetPaymentsByAccountID(context.Context, *payment_pb.GetPaymentsByAccountIDRequest, ...grpc.CallOption) (*payment_pb.GetPaymentsByAccountIDResponse, error)
}

type Service struct {
	store          Store
	paymentService PaymentService
}

type User struct {
	ID              string
	Email           string
	Username        string
	FirstName       string
	LastName        string
	Password        string
	ConfirmPassword string
}

func NewService(s Store, ps PaymentService) *Service {
	return &Service{
		store:          s,
		paymentService: ps,
	}
}

// Compile-time assertion that *Service implements the user.Service interface.
var _ account.Service = (*Service)(nil)

func convertToPaymentResponse(payments *payment_pb.GetPaymentsByAccountIDResponse) []*account.PaymentResponse {
	arr := make([]*account.PaymentResponse, 0, payments.Total)

	for _, payment := range payments.Resources {
		dto := &account.PaymentResponse{
			ID:            uint(payment.Id),
			CreatedAt:     payment.CreatedAt,
			PaymentNumber: payment.PaymentNumber,
			Amount:        payment.Amount,
		}
		arr = append(arr, dto)
	}
	return arr
}

func (s *Service) Register(ctx context.Context, p *account.RegisterPayload) (*account.RegisterResponse, error) {
	// Check passwords
	if p.Password != p.ConfirmedPassword {
		return nil, errors.New("passwords do not match")
	}

	// Encrypt password
	encryptedPassword, encryptErr := auth.EncryptPassword(p.Password)
	if encryptErr != nil {
		return nil, fmt.Errorf("error encrypting the password: %w", encryptErr)
	}
	log.Println("Password has been encrypted successfully.")

	newUser := &User{
		Email:     p.Email,
		FirstName: p.FirstName,
		Username:  uuid.New().String(),
		LastName:  p.LastName,
		Password:  encryptedPassword,
	}

	// Save in the database
	_, err := s.store.CreateUser(ctx, newUser)

	if err != nil {
		return nil, fmt.Errorf("error creating new user: %w", err)
	}

	response := &account.RegisterResponse{
		Message: "Successfully created.",
	}

	return response, nil
}

func (s *Service) Login(ctx context.Context, p *account.LoginPayload) (*account.LoginResponse, error) {
	// Get the user by email
	foundAccount, err := s.store.GetUserByEmail(ctx, p.Email)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify the password hash
	isPasswordCorrect := auth.CheckPassword(foundAccount.Password, p.Password)
	if !isPasswordCorrect {
		return nil, errors.New("invalid credentials")
	}
	log.Println("Password has been verified successfully.")

	// Generate JWT
	token, err := auth.GenerateJWT(foundAccount.ID, foundAccount.Email)
	if err != nil {
		return nil, fmt.Errorf("error extracting jwt: %w", err)
	}

	response := &account.LoginResponse{
		Email:        p.Email,
		Token:        token,
		RefreshToken: "dfsdvsdfjv sjf",
		Role:         "ROLE_REG_USER",
	}
	return response, nil
}

func (s *Service) GetProfile(ctx context.Context, p *account.GetProfilePayload) (*account.ProfileResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	// Get the user by email
	foundAccount, err := s.store.GetUserByEmail(ctx, tokenClaims.Email)

	if err != nil {
		return nil, errors.New("such account does not exist")
	}

	// Get the payments from payment service
	grpcPayload := &payment_pb.GetPaymentsByAccountIDRequest{
		AccountId: int32(foundAccount.ID),
		Limit:     int32(p.PaymentsLimit),
	}

	grpcResponse, err := s.paymentService.GetPaymentsByAccountID(ctx, grpcPayload)

	response := &account.ProfileResponse{
		Email:        foundAccount.Email,
		FirstName:    foundAccount.FirstName,
		LastName:     foundAccount.LastName,
		Username:     foundAccount.Username,
		LastPayments: nil,
	}
	if err != nil {
		return response, nil
	}
	paymentDtos := convertToPaymentResponse(grpcResponse)
	if err != nil {
		fmt.Printf("error gRPC response: %s", err)
	}
	response.LastPayments = paymentDtos
	return response, nil
}
