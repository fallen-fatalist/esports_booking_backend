package entities

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

// user_id SERIAL PRIMARY KEY,
// login VARCHAR(255) UNIQUE NOT NULL,
// email VARCHAR(255) UNIQUE NOT NULL,
// hashed_password CHAR(60) NOT NULL,
// status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'blocked')),
// role VARCHAR(20) DEFAULT 'client',
// balance INT DEFAULT 0,
// created_at TIMESTAMP DEFAULT NOW(),
