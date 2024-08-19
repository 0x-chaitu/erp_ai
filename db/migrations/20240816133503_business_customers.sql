-- migrate:up
CREATE TABLE business_customers (
    company_id SERIAL PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL UNIQUE,
    business_segment UUID REFERENCES business_segments(id),
    address TEXT,
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    zip_code VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_business_customers_company_name ON business_customers(company_name);

CREATE INDEX idx_business_segment_id ON business_customers(business_segment);

-- migrate:down
DROP TABLE business_customers;