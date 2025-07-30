-- ✅ Безопасная очистка только данных из realm 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'

-- Удаление зависимостей, связанных с пользователями
DELETE FROM user_role_mapping WHERE user_id IN (
    SELECT id FROM user_entity WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'
);
DELETE FROM credential WHERE user_id IN (
    SELECT id FROM user_entity WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'
);
DELETE FROM user_entity WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e';

-- Удаление зависимостей, связанных с ролями
DELETE FROM composite_role
WHERE child_role IN (
    SELECT id FROM keycloak_role WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'
);
DELETE FROM group_role_mapping
WHERE role_id IN (SELECT id FROM keycloak_role WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e');
DELETE FROM client_scope_role_mapping
WHERE role_id IN (SELECT id FROM keycloak_role WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e');
DELETE FROM client_session_role
WHERE role_id IN (SELECT id FROM keycloak_role WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e');
DELETE FROM fed_user_role_mapping
WHERE role_id IN (SELECT id FROM keycloak_role WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e');

-- Удаление самих ролей
DELETE FROM keycloak_role WHERE realm_id = 'f3956579-8e47-46ae-8fb7-c7f92168ad1e';

-- Вставка ролей
INSERT INTO keycloak_role (id, name, realm_id) VALUES
                                                   ('role-admin-id', 'ROLE_ADMIN', 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'),
                                                   ('role-teacher-id', 'ROLE_TEACHER', 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'),
                                                   ('role-user-id', 'ROLE_USER', 'f3956579-8e47-46ae-8fb7-c7f92168ad1e');

-- Вставка пользователей
INSERT INTO user_entity (id, username, email, email_verified, enabled, realm_id) VALUES
                                                                                     ('user-admin-id', 'admin', 'admin@example.com', true, true, 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'),
                                                                                     ('user-teacher-id', 'teacher', 'teacher@example.com', true, true, 'f3956579-8e47-46ae-8fb7-c7f92168ad1e'),
                                                                                     ('user-user-id', 'user', 'user@example.com', true, true, 'f3956579-8e47-46ae-8fb7-c7f92168ad1e');

-- Назначение ролей пользователям
INSERT INTO user_role_mapping (user_id, role_id) VALUES
                                                     ('user-admin-id', 'role-admin-id'),
                                                     ('user-teacher-id', 'role-teacher-id'),
                                                     ('user-user-id', 'role-user-id');

-- Вставка паролей (bcrypt: admin123 / teacher123 / user123)
INSERT INTO credential (id, type, created_date, user_label, secret_data, credential_data, user_id, priority) VALUES
                                                                                                                 ('cred-admin-id', 'password', extract(epoch from now()) * 1000, NULL,
                                                                                                                  '{"value":"$2a$12$eXz62oTOyNntZ5VxxkCVTOjWyDP4idWeuZ.TOSt4fRDLgTZ/tT6j2"}',
                                                                                                                  '{"hashIterations":27500,"algorithm":"bcrypt"}',
                                                                                                                  'user-admin-id', 0),

                                                                                                                 ('cred-teacher-id', 'password', extract(epoch from now()) * 1000, NULL,
                                                                                                                  '{"value":"$2a$12$zBJeGMw7AXF7agapOrPq3eZRAeMa5xjFgPvTF5EZimAR2ZGxwQAAW"}',
                                                                                                                  '{"hashIterations":27500,"algorithm":"bcrypt"}',
                                                                                                                  'user-teacher-id', 0),

                                                                                                                 ('cred-user-id', 'password', extract(epoch from now()) * 1000, NULL,
                                                                                                                  '{"value":"$2a$12$6qLY5TyRPpbK5AHmJgOq0O6nPb9cUdYrZQVGnPE52VNKboMntcF3K"}',
                                                                                                                  '{"hashIterations":27500,"algorithm":"bcrypt"}',
                                                                                                                  'user-user-id', 0);
