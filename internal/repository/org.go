package repository

import (
	"context"

	"github.com/0x-chaitu/rag_erp/internal/domain"
)

type postgresOrgRepo struct {
	conn Conn
}

// Create implements domain.UserRepo.
func (p *postgresOrgRepo) Create(ctx context.Context, org *domain.Organization) error {
	query := `INSERT INTO organizations (org_id, name, subdomain) VALUES ($1,$2,$3) RETURNING org_id;`

	_, err := p.conn.Exec(ctx, query, org.OrgID, org.Name, org.Subdomain)
	return err

}

func NewPostgresOrg(conn Conn) domain.OrgRepo {
	return &postgresOrgRepo{conn: conn}
}
