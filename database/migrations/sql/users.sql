CREATE TABLE IF NOT EXISTS eplanner.users (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `email` longtext,
  `password` longtext,
  `address` longtext,
  `occupation` longtext,
  `phone` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;