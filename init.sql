CREATE DATABASE kk;

CREATE TABLE auth (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_activated BOOLEAN DEFAULT false NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    role VARCHAR(10) DEFAULT 'student' CHECK (role IN ('student', 'teacher', 'admin'))
);
