-- +goose Up
-- +goose StatementBegin
CREATE TABLE lessons (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     description TEXT,
     content TEXT,
     order_position INT,
     chapter_id INT REFERENCES chapters(id) ON DELETE CASCADE,
     created_at TIMESTAMP NOT NULL DEFAULT now(),
     updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO lessons (name, description, content, order_position, chapter_id) VALUES
                                                                                 (
     'If-Else Statement in Go',
     'Using conditional logic in Go',
     '<h2>If-Else in Go</h2>' ||
     '<p>In Go, the <code>if</code> statement allows you to execute code blocks conditionally. It has a simple syntax without parentheses, but curly braces are required.</p>' ||
     '<pre><code>if x &gt; 10 {\n  fmt.Println("x is greater than 10")\n} else {\n  fmt.Println("x is 10 or less")\n}</code></pre>' ||
     '<p>You can also declare variables inside the <code>if</code> statement:</p>' ||
     '<pre><code>if y := compute(); y &gt; 0 {\n  fmt.Println("Positive")\n}</code></pre>',
     1,
     2
    ),
    (
     'Switch Statement in Go',
     'Using switch-case for multi-branch logic',
     '<h2>Switch in Go</h2>' ||
     '<p>The <code>switch</code> statement simplifies multiple conditions. Go automatically breaks after a matching case.</p>' ||
     '<pre><code>switch day := time.Now().Weekday(); day {\ncase time.Saturday, time.Sunday:\n  fmt.Println("Weekend")\ndefault:\n  fmt.Println("Weekday")\n}</code></pre>' ||
     '<p>You can also use <code>switch</code> without an expression:</p>' ||
     '<pre><code>switch {\ncase x &gt; 0:\n  fmt.Println("Positive")\ncase x &lt; 0:\n  fmt.Println("Negative")\ndefault:\n  fmt.Println("Zero")\n}</code></pre>',
     2,
     2
    ),
    (
     'Defining Functions in Go',
     'Learn how to write and use functions',
     '<h2>Functions in Go</h2>' ||
     '<p>Functions in Go are defined using the <code>func</code> keyword. You must specify parameter types and return types.</p>' ||
     '<pre><code>func add(a int, b int) int {\n  return a + b\n}</code></pre>' ||
     '<p>Functions can return multiple values:</p>' ||
     '<pre><code>func divide(a, b int) (int, error) {\n  if b == 0 {\n    return 0, errors.New("division by zero")\n  }\n  return a / b, nil\n}</code></pre>' ||
     '<p>You can also define anonymous functions and closures.</p>',
     1,
     3
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lessons;
-- +goose StatementEnd
