CREATE TYPE user_role AS ENUM ('client', 'admin', 'analyst', 'manager');
CREATE TYPE user_status AS ENUM ('active', 'blocked');
CREATE TYPE user_promotion AS ENUM ('flyer', 'birthday', '')

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    status user_status DEFAULT 'active',
    role user_role DEFAULT 'client',
    promotion promotion DEFAULT '',
    created_at TIMESTAMPZ DEFAULT NOW(),
    CONSTRAINT email_format_check CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE TABLE balances (
    user_id SERIAL PRIMARY KEY REFERENCES users(user_id),
    balance INT DEFAULT 0
)