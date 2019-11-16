-- +goose Up
CREATE TABLE review (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  class_id INT(10) UNSIGNED NOT NULL,
  difficulty_credit INT(10) UNSIGNED NOT NULL,
  -- attendence INT(10) UNSIGNED,
  evaluation INT(10) UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (class_id) REFERENCES classes(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE review;
