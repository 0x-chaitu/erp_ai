package repository

import (
	"context"

	"github.com/0x-chaitu/rag_erp/internal/domain"
	"github.com/google/uuid"
)

type postgresUserRepo struct {
	conn Conn
}

// Create implements domain.UserRepo.
func (p *postgresUserRepo) Create(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	var id uuid.UUID
	query := `INSERT INTO users (google_user_id, org_id) VALUES ($1,$2) RETURNING id;`

	if err := p.conn.QueryRow(ctx, query, user.GoogleUserID, user.OrgId).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil

}

func NewPostgresUser(conn Conn) domain.UserRepo {
	return &postgresUserRepo{conn: conn}
}
