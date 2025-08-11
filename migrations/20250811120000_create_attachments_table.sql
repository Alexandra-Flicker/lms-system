-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS attachments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    lesson_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_lesson
        FOREIGN KEY (lesson_id) 
        REFERENCES lessons(id) 
        ON DELETE CASCADE
);

CREATE INDEX idx_attachments_lesson_id ON attachments(lesson_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attachments;
-- +goose StatementEnd