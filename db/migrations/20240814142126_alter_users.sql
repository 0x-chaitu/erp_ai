-- migrate:up
ALTER TABLE
    users
ADD
    COLUMN org_id TEXT NOT NULL,
ADD
    CONSTRAINT fk_org_id FOREIGN KEY (org_id) REFERENCES organizations(org_id);

-- migrate:down
ALTER TABLE
    users DROP CONSTRAINT fk_org_id;

ALTER TABLE
    users DROP COLUMN org_id;