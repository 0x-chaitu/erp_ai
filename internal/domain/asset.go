package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID           uuid.UUID `json:"id"`
	SerialNumber string    `json:"serial_number"`
	AssetName    string    `json:"asset_name"`
	AssetModel   string    `json:"asset_model"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AssetRepo interface {
	GetAssets(ctx context.Context, offset, limit int) ([]*Asset, error)
	GetAssetCount(ctx context.Context) (int, error)
}
