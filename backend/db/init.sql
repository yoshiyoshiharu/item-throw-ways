DROP TABLE IF EXISTS items;
CREATE TABLE items (
  id INT NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  price INT NOT NULL DEFAULT 0,
  remarks VARCHAR(255) NOT NULL DEFAULT '',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

DROP TABLE IF EXISTS kinds;
CREATE TABLE kinds (
  id INT NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

DROP TABLE IF EXISTS item_kinds;
CREATE TABLE item_kinds (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  item_id INT NOT NULL,
  kind_id INT NOT NULL,
  price INT NOT NULL DEFAULT 0,
  remarks VARCHAR(255) NOT NULL DEFAULT '',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

DROP TABLE IF EXISTS areas;
CREATE TABLE areas (
  id INT NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

DROP TABLE IF EXISTS area_collect_weekdays;
CREATE TABLE area_collect_weekdays (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  area_id INT NOT NULL,
  kind_id INT NOT NULL,
  weekday INT NOT NULL,
  lap INT NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

INSERT INTO kinds (id, name) VALUES
  (1, '可燃ごみ'),
  (2, '不燃ごみ'),
  (3, '不燃ごみ（水銀含有物）'),
  (4, '資源'),
  (5, '粗大ごみ'),
  (6, '不可')
;
