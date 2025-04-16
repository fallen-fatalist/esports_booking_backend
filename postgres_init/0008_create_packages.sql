CREATE TABLE packages (
    package_id SERIAL PRIMARY KEY,
    package_name TEXT,
    tier computer_tier DEFAULT 'common',
    price INT NOT NULL,
    startTime TIME,
    endTime TIME,
    created_at TIMESTAMP DEFAULT NOW()
);