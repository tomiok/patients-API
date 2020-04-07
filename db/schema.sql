CREATE TABLE IF NOT EXISTS patient (
    id serial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    address text NOT NULL,
    phone text NOT NULL,
    email text NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    CONSTRAINT unique_patient_name UNIQUE(first_name, last_name)
);

CREATE TABLE IF NOT EXISTS physician (
    id serial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    CONSTRAINT unique_physician_name UNIQUE(first_name, last_name)
);

CREATE TABLE IF NOT EXISTS visit (
    id serial PRIMARY KEY,
    patient_id integer NOT NULL REFERENCES patient(id),
    physician_id integer NOT NULL REFERENCES physician(id),
    visited_at timestamptz DEFAULT NOW(),
    location text NOT NULL,
    reason text NOT NULL
);

