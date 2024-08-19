-- migrate:up
CREATE TABLE organizations (
    -- google tenant_id
    org_id TEXT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    subdomain VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- migrate:down
CREATE INDEX idx_organizations_name ON organizations(name);

CREATE INDEX idx_organizations_subdomain ON organizations(subdomain);