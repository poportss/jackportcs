CREATE TABLE IF NOT EXISTS schema_versions (
                                               id UUID PRIMARY KEY,
                                               created_at TIMESTAMP WITH TIME ZONE,
                                               updated_at TIMESTAMP WITH TIME ZONE,
                                               deleted_at TIMESTAMP WITH TIME ZONE,
                                               service VARCHAR(255) NOT NULL,
                                               version INT NOT NULL
)