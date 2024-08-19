package domain

import (
	"context"
	"time"
)

type Organization struct {
	OrgID     string    `json:"org_id"`
	Name      string    `json:"name"`
	Subdomain string    `json:"subdomain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrgRepo interface {
	Create(ctx context.Context, org *Organization) error
}
