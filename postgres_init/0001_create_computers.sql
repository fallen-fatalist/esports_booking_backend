CREATE TYPE computer_tier AS ENUM ('standard', 'vip');
CREATE TYPE computer_status AS ENUM ('busy', 'available', 'pending', 'not working', 'under repair');

CREATE TABLE computers (
    computer_id SERIAL PRIMARY KEY,
    status computer_status DEFAULT 'available',
    tier computer_tier DEFAULT 'standard',
    cpu TEXT NOT NULL,
    gpu TEXT NOT NULL,
    ram TEXT NOT NULL,
    ssd TEXT NOT NULL,
    hdd TEXT NOT NULL,
    monitor TEXT NOT NULL,
    keyboard TEXT NOT NULL,
    headset TEXT NOT NULL,
    mouse TEXT NOT NULL,
    mousepad TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
    updated_at TIMESTAMP DEFAULT NOW()
);



