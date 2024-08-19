-- migrate:up
CREATE TABLE IF NOT EXISTS public.business_verticals (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE INDEX idx_business_verticals_id ON business_verticals (id);

CREATE INDEX idx_business_verticals_name ON public.business_verticals (name);

-- migrate:down
DROP TABLE business_verticals;

DROP INDEX idx_business_verticals_name;

DROP INDEX idx_business_verticals_id;