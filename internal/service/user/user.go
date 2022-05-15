package user

import (
	"context"
	"notary-public-online/internal/entity/model"
)

type User interface {
	Register(ctx context.Context, user model.User) (model.User, error)

	Login(ctx context.Context, email string, password string) (bool, error)
}
