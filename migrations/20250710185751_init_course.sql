-- +goose Up
-- +goose StatementBegin

-- CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE courses (

     id SERIAL PRIMARY KEY, -- id UUID PRIMARY KEY DEFAULT gen_random_uuid()
     name VARCHAR(255) NOT NULL,
     description TEXT,
     created_at TIMESTAMP NOT NULL DEFAULT now(),
     updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO courses (name, description) VALUES
    (
        'Golang Developer',
        'Master the Go programming language from the ground up. Learn syntax, data structures, concurrency, interfaces, and build real-world applications.'
    ),
    (
        'Python Developer',
        'A comprehensive Python course covering syntax, OOP, data science basics, web development with Flask, and more.'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS courses;
-- +goose StatementEnd
