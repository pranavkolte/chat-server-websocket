CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password TEXT NOT NULL
);
