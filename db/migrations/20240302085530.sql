-- insert status
INSERT INTO
    status (id, name)
VALUES
    ('018dfe5d-7447-79a4-bb18-e133cad6501f', 'To-Do'),
    ('018dfe5d-7447-72ef-8cc6-5125b887a4d6', 'In Progress'),
    ('018dfe5d-7447-77cb-914b-b515a0eb1eef', 'Done');

-- insert board_status
INSERT INTO
    board_status (id, status_id, position)
VALUES
    ('018dfe62-5e54-7d66-a45c-7c685e3f553e', '018dfe5d-7447-79a4-bb18-e133cad6501f', 0),
    ('018dfe62-5e54-77b2-8ca6-b31b69d19147', '018dfe5d-7447-72ef-8cc6-5125b887a4d6', 1),
    ('018dfe62-5e54-77c6-bcf5-4664139c2f55', '018dfe5d-7447-77cb-914b-b515a0eb1eef', 2);