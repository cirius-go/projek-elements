package auth

import (
	"context"

	"github.com/projek-elements/internal/types/tservice/tauth"
)

var (
	// implement.
	_ tauth.Service = &Service{}
)

// New creates new Service instance.
func New() *Service {
	return &Service{}
}

// Service auth.
type Service struct{}

// Login implements tauth.Service.
func (*Service) Login(c context.Context, req *tauth.LoginRequest) (res *tauth.LoginResponse, err error) {
	panic("unimplemented")
}

// Register implements tauth.Service.
func (*Service) Register(c context.Context, req *tauth.RegisterRequest) (res *tauth.RegisterResponse, err error) {
	panic("unimplemented")
}
