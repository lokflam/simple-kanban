-- create "statuses" table
CREATE TABLE
    statuses (
        id uuid PRIMARY KEY,
        name text NOT NULL,
        created_at timestamptz NOT NULL DEFAULT current_timestamp,
        updated_at timestamptz NOT NULL DEFAULT current_timestamp
    );

-- create "cards" table
CREATE TABLE
    cards (
        id uuid PRIMARY KEY,
        title text NOT NULL,
        content text NOT NULL,
        status_id uuid NOT NULL,
        created_at timestamptz NOT NULL DEFAULT current_timestamp,
        updated_at timestamptz NOT NULL DEFAULT current_timestamp,
        FOREIGN KEY (status_id) REFERENCES statuses (id)
    );