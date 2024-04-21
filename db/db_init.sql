INSERT INTO users (username, email, password) VALUES ('duyhung', 'duyhung@gmail.com', 'duyhung');
INSERT INTO pages (user_id, title) VALUES (1, 'First page');
INSERT INTO pages (user_id, title, parent_id) VALUES (2, 'Second Sub Page', 1);