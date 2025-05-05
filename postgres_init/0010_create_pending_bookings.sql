CREATE SEQUENCE global_booking_id_seq;

CREATE TYPE booking_status AS ENUM ('pending', 'active', 'finished', 'cancelled');

CREATE TABLE bookings (
    booking_id BIGINT DEFAULT nextval('global_booking_id_seq') PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(user_id),
    computer_id SERIAL NOT NULL REFERENCES computers(computer_id),
    package_id SERIAL REFERENCES packages(package_id) ,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    status booking_status DEFAULT 'pending',
    created_at TIMESTAMPTZ DEFAULT NOW()
);