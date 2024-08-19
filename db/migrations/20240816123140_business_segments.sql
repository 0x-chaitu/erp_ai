-- migrate:up
CREATE TABLE IF NOT EXISTS public.business_segments (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE INDEX idx_business_segments_id ON business_segments (id);

CREATE INDEX idx_business_segments_name ON public.business_segments (name);

-- migrate:down
DROP TABLE business_segments;

DROP INDEX idx_business_segments_name;

DROP INDEX idx_business_segments_id;