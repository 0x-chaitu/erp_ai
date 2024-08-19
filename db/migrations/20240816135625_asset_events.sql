-- migrate:up
CREATE TABLE asset_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id UUID REFERENCES assets(id),
    event_type VARCHAR(50),
    -- Possible values: activated, deactivated, maintenance_started, maintenance_ended, calibration_started, calibration_ended, rented_out, rented_in
    event_data JSONB,
    -- Store additional event-specific data
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_asset_events_asset_id ON asset_events(asset_id);

CREATE INDEX idx_asset_events_event_type ON asset_events(event_type);

CREATE INDEX idx_asset_events_created_at ON asset_events(created_at);

CREATE INDEX idx_asset_events_created_at_asset_id ON asset_events(created_at, asset_id);

-- migrate:down
DROP TABLE asset_events;