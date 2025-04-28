CREATE TYPE finished_booking_status AS ENUM ('finished', 'cancelled');

CREATE TABLE finished_bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(user_id),
    computer_id SERIAL NOT NULL REFERENCES computers(computer_id),
    package_id SERIAL NOT NULL REFERENCES packages(package_id),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    total_price INT NOT NULL,
    status finished_booking_status DEFAULT 'finished',
    created_at TIMESTAMPTZ NOT NULL
);
