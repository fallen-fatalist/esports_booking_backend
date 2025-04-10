CREATE TYPE booking_history AS ENUM ('finished', 'cancelled');

CREATE TABLE finished_bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(user_id),
    computer_id SERIAL NOT NULL REFERENCES computers(computer_id),
    package_id SERIAL NOT NULL REFERENCES packages(package_id),
    start_time TIMESTAMPZ NOT NULL,
    end_time TIMESTAMPZ NOT NULL,
    total_price INT NOT NULL,
    status finished_booking_status DEFAULT 'finished',
    created_at TIMESTAMP NOT NULL
);
