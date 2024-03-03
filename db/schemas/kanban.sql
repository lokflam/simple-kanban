-- create "status" table
CREATE TABLE
    status (
        id uuid PRIMARY KEY,
        name text NOT NULL,
        created_at timestamptz NOT NULL DEFAULT current_timestamp,
        updated_at timestamptz NOT NULL DEFAULT current_timestamp
    );

-- create "board_status" table
CREATE TABLE
    board_status (
        id uuid PRIMARY KEY,
        status_id uuid NOT NULL,
        position int NOT NULL,
        created_at timestamptz NOT NULL DEFAULT current_timestamp,
        updated_at timestamptz NOT NULL DEFAULT current_timestamp,
        FOREIGN KEY (status_id) REFERENCES status (id)
    );

-- create "card" table
CREATE TABLE
    card (
        id uuid PRIMARY KEY,
        title text NOT NULL,
        content text NOT NULL,
        status_id uuid NOT NULL,
        created_at timestamptz NOT NULL DEFAULT current_timestamp,
        updated_at timestamptz NOT NULL DEFAULT current_timestamp,
        FOREIGN KEY (status_id) REFERENCES status (id)
    );