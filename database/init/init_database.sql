-- CREATE DATABASE IF NOT EXISTS persona;
-- 1. usersテーブルの作成
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(255)
);

-- 2. personaテーブルの作成
CREATE TABLE IF NOT EXISTS persona (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  sex VARCHAR(10) CHECK (sex IN ('male', 'female', 'other')),
  age INTEGER,
  profession VARCHAR(100),
  problems TEXT,
  behavior TEXT
);

-- 3. conversationテーブルの作成
CREATE TABLE IF NOT EXISTS conversation (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  persona_id INTEGER REFERENCES persona(id) ON DELETE CASCADE
);

-- 4. commentテーブルの作成
CREATE TABLE IF NOT EXISTS comment (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  persona_id INTEGER REFERENCES persona(id) ON DELETE CASCADE,
  comment TEXT NOT NULL,
  is_user_comment BOOLEAN,
  good BOOLEAN DEFAULT FALSE
);


-- usersテーブルに5件のデータを挿入
INSERT INTO users (id, name, email, password)
VALUES 
  (1, 'test1', 'test1@example.com', 'password1'),
  (2, 'test2', 'test2@example.com', 'password2'),
  (3, 'test3', 'test3@example.com', 'password3'),
  (4, 'test4', 'test4@example.com', 'password4'),
  (5, 'test5', 'test5@example.com', 'password5')
ON CONFLICT DO NOTHING;

-- personaテーブルに5件のデータを挿入
INSERT INTO persona (id, name, user_id, sex, age, profession, problems, behavior)
VALUES 
  (1, 'Persona1', 1, 'male', 25, 'Engineer', 'Problem1', 'Behavior1'),
  (2, 'Persona2', 2, 'female', 30, 'Designer', 'Problem2', 'Behavior2'),
  (3, 'Persona3', 3, 'male', 22, 'Student', 'Problem3', 'Behavior3'),
  (4, 'Persona4', 4, 'female', 27, 'Doctor', 'Problem4', 'Behavior4'),
  (5, 'Persona5', 5, 'other', 35, 'Writer', 'Problem5', 'Behavior5')
ON CONFLICT DO NOTHING;

-- conversationテーブルに5件のデータを挿入
INSERT INTO conversation (id, user_id, persona_id)
VALUES 
  (1, 1, 1),
  (2, 2, 2),
  (3, 3, 3),
  (4, 4, 4),
  (5, 5, 5)
ON CONFLICT DO NOTHING;

-- commentテーブルに5件のデータを挿入
INSERT INTO comment (id, user_id, persona_id, comment, is_user_comment, good)
VALUES 
  (1, 1, 1, 'Comment1', TRUE, TRUE),
  (2, 2, 2, 'Comment2', TRUE, FALSE),
  (3, 3, 3, 'Comment3', FALSE, TRUE),
  (4, 4, 4, 'Comment4', FALSE, FALSE),
  (5, 5, 5, 'Comment5', TRUE, TRUE)
ON CONFLICT DO NOTHING;
