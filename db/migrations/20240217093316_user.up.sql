CREATE TABLE `user` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `username` varchar(30) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL,
  `deleted_at` TIMESTAMP DEFAULT NULL
);