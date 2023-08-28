package tauth

import "context"

// Service represents the auth service.
type Service interface {
	Login(c context.Context, req *LoginRequest) (res *LoginResponse, err error)
	Register(c context.Context, req *RegisterRequest) (res *RegisterResponse, err error)
}

type (
	// RegisterRequest is the input data for the business logic of Register.
	RegisterRequest struct{}

	// RegisterResponse is the output data for the business logic of Register.
	RegisterResponse struct{}
)

type (
	// LoginRequest is the input data for the business logic of Login.
	LoginRequest struct{}

	// LoginResponse is the output data for the business logic of Login.
	LoginResponse struct{}
)
