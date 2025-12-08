CREATE TABLE sensors (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    device_id UUID NOT NULL REFERENCES devices(id),
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);