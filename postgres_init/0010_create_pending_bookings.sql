CREATE TYPE booking_status AS ENUM ('pending', 'active');

CREATE TABLE pending_bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(user_id),
    computer_id SERIAL NOT NULL REFERENCES computers(computer_id),
    package_id SERIAL REFERENCES packages(package_id) ,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    total_price INT NOT NULL,
    status booking_status DEFAULT 'pending',
    created_at TIMESTAMPTZ DEFAULT NOW()
);