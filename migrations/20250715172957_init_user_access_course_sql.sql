-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_access_course (
      user_id       UUID    NOT NULL,
      course_id     INTEGER    NOT NULL,
      unlocked      BOOLEAN NOT NULL DEFAULT FALSE,
      created_at    TIMESTAMP NOT NULL DEFAULT now(),

      PRIMARY KEY (user_id, course_id),
      FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_access_course;
-- +goose StatementEnd
