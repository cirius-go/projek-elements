package api

import "github.com/projek-elements/internal/types/tservice/tauth"

// AuthHTTP is the HTTP handler for auth api.
type AuthHTTP struct {
	svc tauth.Service
}

// NewAuthHTTP creates new AuthHTTP instance.
func NewAuthHTTP(svc tauth.Service) *AuthHTTP {
	return &AuthHTTP{svc}
}
