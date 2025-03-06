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

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    thumbnail VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE course_teacher (
    id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES course(id) ON DELETE CASCADE,
    teacher_id INTEGER REFERENCES auth(id) ON DELETE CASCADE
);

CREATE TABLE enrollment (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES auth(id) ON DELETE CASCADE,
    course_id INTEGER REFERENCES course(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT now()
);

CREATE TABLE section (
    id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES course(id) ON DELETE CASCADE,
    section_number INTEGER NOT NULL CHECK (section_number > 0),
    title VARCHAR(100) NOT NULL
);

CREATE TABLE video (
    id SERIAL PRIMARY KEY,
    section_id INTEGER REFERENCES section(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    video_url TEXT NOT NULL,
    duration INTEGER NOT NULL CHECK (duration > 0)
);
