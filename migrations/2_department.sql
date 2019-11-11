-- +goose Up
CREATE TABLE department (
  id INT(10) UNSIGNED NOT NULL,
  name VARCHAR(255) NOT NULL,
  campus VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE department;
