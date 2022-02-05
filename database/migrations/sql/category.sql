CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO categories(name) VALUES("Arts");
INSERT INTO categories(name) VALUES("Technology");
INSERT INTO categories(name) VALUES("Sports");
INSERT INTO categories(name) VALUES("Music");
INSERT INTO categories(name) VALUES("Education");