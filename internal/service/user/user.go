package user

import (
	"context"
	"notary-public-online/internal/dto"
)

type User interface {
	Register(ctx context.Context, inp dto.RegisterCredential) error

	Login(ctx context.Context, inp dto.LoginCredential) (bool, error)
}
