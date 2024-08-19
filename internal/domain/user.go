package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	OrgId        string    `json:"org_id"`
	GoogleUserID string    `json:"google_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (uuid.UUID, error)
}
