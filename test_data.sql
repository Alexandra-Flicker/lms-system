-- Insert test data for development and testing

-- Insert test courses (if they don't exist)
INSERT INTO courses (id, name, description, created_at, updated_at) 
VALUES 
  (1, 'Golang Developer', 'Master the Go programming language from the ground up. Learn syntax, data structures, concurrency, interfaces, and build real-world applications.', NOW(), NOW()),
  (2, 'Python Developer', 'A comprehensive Python course covering syntax, OOP, data science basics, web development with Flask, and more.', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for courses
SELECT setval('courses_id_seq', (SELECT MAX(id) FROM courses));

-- Insert test chapters
INSERT INTO chapters (id, name, description, order_position, course_id, created_at, updated_at) 
VALUES 
  (1, 'Introduction to Go', 'Getting started with Go programming', 1, 1, NOW(), NOW()),
  (2, 'Advanced Go Concepts', 'Advanced topics in Go', 2, 1, NOW(), NOW()),
  (3, 'Python Basics', 'Introduction to Python programming', 1, 2, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for chapters
SELECT setval('chapters_id_seq', (SELECT MAX(id) FROM chapters));

-- Insert test lessons
INSERT INTO lessons (id, name, description, content, order_position, chapter_id, created_at, updated_at) 
VALUES 
  (1, 'Go Variables and Data Types', 'Learn about variables and basic data types in Go', 'Go is a statically typed language. Variables can be declared using var keyword or short declaration syntax.', 1, 1, NOW(), NOW()),
  (2, 'Go Control Flow', 'Understanding if statements, loops, and switches in Go', 'Go provides if, for, and switch statements for control flow.', 2, 1, NOW(), NOW()),
  (3, 'Go Functions', 'How to define and use functions in Go', 'Functions in Go can return multiple values and support variadic parameters.', 3, 1, NOW(), NOW()),
  (4, 'Concurrency with Goroutines', 'Introduction to concurrent programming in Go', 'Goroutines are lightweight threads managed by the Go runtime.', 1, 2, NOW(), NOW()),
  (5, 'Python Syntax Basics', 'Understanding Python syntax and indentation', 'Python uses indentation to define code blocks.', 1, 4, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for lessons
SELECT setval('lessons_id_seq', (SELECT MAX(id) FROM lessons));

-- Insert test user access (admin user has access to both courses)
-- Using the Keycloak admin user UUID and a test user ID
INSERT INTO user_access_course (user_id, course_id, unlocked, created_at)
VALUES 
  ('32bfb3d7-5b2c-4502-b08a-92ae81984f57'::uuid, 1, true, NOW()),
  ('32bfb3d7-5b2c-4502-b08a-92ae81984f57'::uuid, 2, true, NOW()),
  ('00000000-0000-0000-0000-000000000001'::uuid, 1, true, NOW()),
  ('00000000-0000-0000-0000-000000000001'::uuid, 2, true, NOW())
ON CONFLICT (user_id, course_id) DO UPDATE SET unlocked = true;

-- Insert some sample attachments for testing
INSERT INTO attachments (id, name, url, lesson_id, created_at, updated_at)
VALUES 
  (1, 'go-cheatsheet.pdf', 'lessons/1/go-cheatsheet_1234567890.pdf', 1, NOW(), NOW()),
  (2, 'variables-examples.txt', 'lessons/1/variables-examples_1234567891.txt', 1, NOW(), NOW()),
  (3, 'goroutines-guide.pdf', 'lessons/4/goroutines-guide_1234567892.pdf', 4, NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

-- Reset sequence for attachments
SELECT setval('attachments_id_seq', (SELECT MAX(id) FROM attachments));

-- Display inserted data
SELECT 'Courses:' as section;
SELECT id, name, description FROM courses ORDER BY id;

SELECT 'Chapters:' as section;
SELECT c.id, c.name, c.course_id, co.name as course_name FROM chapters c JOIN courses co ON c.course_id = co.id ORDER BY c.id;

SELECT 'Lessons:' as section;  
SELECT l.id, l.name, l.chapter_id, ch.name as chapter_name FROM lessons l JOIN chapters ch ON l.chapter_id = ch.id ORDER BY l.id;

SELECT 'User Access:' as section;
SELECT uca.user_id, uca.course_id, c.name as course_name, uca.unlocked FROM user_access_course uca JOIN courses c ON uca.course_id = c.id ORDER BY uca.user_id, uca.course_id;

SELECT 'Attachments:' as section;
SELECT a.id, a.name, a.lesson_id, l.name as lesson_name FROM attachments a JOIN lessons l ON a.lesson_id = l.id ORDER BY a.id;