CREATE TYPE computer_status AS ENUM ('busy', 'available', 'pending', 'not working', 'under repair');

CREATE TABLE computers_statuses (
    computer_id SERIAL PRIMARY KEY REFERENCES computers(computer_id),
    status computer_status DEFAULT 'available',
    updated_at TIMESTAMPZ DEFAULT NOW()
);