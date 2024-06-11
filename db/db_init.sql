INSERT INTO `user` (username, email, password) VALUES ('duyhung', 'duyhung@gmail.com', 'duyhung');
INSERT INTO `page` (user_id, title) VALUES (1, 'First page');
INSERT INTO `page` (user_id, title, parent_id) VALUES (2, 'Second Sub Page', 1);