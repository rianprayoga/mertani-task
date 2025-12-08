CREATE TABLE devices (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    latitude double precision 	NOT NULL,
    longitude double precision NOT NULL,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);