package repository

import (
	"context"

	"github.com/0x-chaitu/rag_erp/internal/domain"
)

type postgresAssetRepo struct {
	conn Conn
}

func (p *postgresAssetRepo) GetAssets(ctx context.Context, offset, limit int) ([]*domain.Asset, error) {
	query := `SELECT asset_name, asset_model, serial_number FROM assets OFFSET $1 LIMIT $2;`

	rows, err := p.conn.Query(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var assets []*domain.Asset
	for rows.Next() {
		var asset domain.Asset
		err := rows.Scan(&asset.AssetName, &asset.AssetModel, &asset.SerialNumber)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return assets, nil

}

func (p *postgresAssetRepo) GetAssetCount(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM assets;`
	var count int
	rows, err := p.conn.Query(ctx, query)
	if err != nil {
		return count, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, nil

}

func NewPostgresAsset(conn Conn) domain.AssetRepo {
	return &postgresAssetRepo{conn: conn}
}
