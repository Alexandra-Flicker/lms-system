-- +goose Up
-- +goose StatementBegin
CREATE TABLE chapters (
      id SERIAL PRIMARY KEY,
      course_id INT REFERENCES courses(id) ON DELETE CASCADE,
      name VARCHAR(255) NOT NULL,
      description TEXT,
      order_position INT,
      created_at TIMESTAMP NOT NULL DEFAULT now(),
      updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO chapters (course_id, name, description, order_position) VALUES
     (1, 'Introduction to Go', 'Getting started with the Go programming language', 1),
     (1, 'Control Structures', 'Understanding if-else, switch, and loops in Go', 2),
     (1, 'Functions and Methods', 'Using functions, methods and receivers in Go', 3),
     (2, 'Python Basics', 'Syntax and basic structures in Python', 1),
     (2, 'Data Structures in Python', 'Lists, dictionaries, sets, and tuples', 2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chapters;
-- +goose StatementEnd
