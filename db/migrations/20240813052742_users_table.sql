-- migrate:up
CREATE TABLE IF NOT EXISTS public.users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    google_user_id VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_google_user_id ON users (google_user_id);

CREATE INDEX idx_users_id ON users (id);

-- migrate:down
DROP TABLE users;

DROP INDEX idx_users_google_user_id;

DROP INDEX idx_users_id;