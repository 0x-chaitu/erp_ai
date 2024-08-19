-- migrate:up
CREATE TABLE assets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_name VARCHAR(255) NOT NULL,
    asset_type VARCHAR(100),
    asset_model VARCHAR(100),
    serial_number VARCHAR(100),
    purchase_date DATE,
    business_vertical_id UUID REFERENCES business_verticals(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_assets_asset_name ON assets(asset_name);

CREATE INDEX idx_assets_asset_type ON assets(asset_type);

CREATE INDEX idx_assets_asset_model ON assets(asset_model);

CREATE INDEX idx_assets_serial_number ON assets(serial_number);

CREATE INDEX idx_assets_business_vertical_id ON assets(business_vertical_id);

CREATE INDEX idx_assets_asset_type_asset_model ON assets(asset_type, asset_model);

-- migrate:down
DROP TABLE assets;