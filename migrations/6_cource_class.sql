-- +goose Up
CREATE TABLE cource_class (
	cource_id INT(10) UNSIGNED NOT NULL,
	class_id INT(10) UNSIGNED NOT NULL,
	FOREIGN KEY (cource_id) REFERENCES cource(id),
	FOREIGN KEY (class_id) REFERENCES class(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE cource_class;
